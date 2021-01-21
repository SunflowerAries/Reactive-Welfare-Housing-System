package storage

import (
	"fmt"
	"log"
	"strings"

	"github.com/thoas/go-funk"
)

func (db *DB) DeleteHouse(HouseID []int32) error {
	args := make([]interface{}, len(HouseID))
	for i, id := range HouseID {
		args[i] = id
	}
	stmt := fmt.Sprintf("UPDATE house SET house.deleted = true WHERE id IN (?%s)", strings.Repeat(",?", len(args)-1))
	_, err := db.Exec(stmt, args...)
	return err
}

func (db *DB) DeleteReside(reside Reside) {
	db.Exec(`DELETE FROM reside WHERE house_id = ? AND family_id = ?`, reside.HouseID, reside.FamilyID)
}

func (db *DB) DeleteHouseAndReside(HouseID []int32) error {
	args := make([]interface{}, len(HouseID))
	for i, id := range HouseID {
		args[i] = id
	}
	stmt := fmt.Sprintf("UPDATE house LEFT JOIN reside "+
		"ON house.id = reside.house_id SET house.deleted = true, reside.checkout = true "+
		"WHERE house.id IN (?%s)", strings.Repeat(",?", len(args)-1))
	_, err := db.Exec(stmt, args...)
	return err
}

func (db *DB) QueryReside(HouseID []int32) []Reside {
	args := make([]interface{}, len(HouseID))
	for i, id := range HouseID {
		args[i] = id
	}
	stmt := fmt.Sprintf("SELECT H.id, H.level, R.family_id "+
		"FROM (SELECT * FROM house WHERE deleted = false AND id IN (?%s)) AS H "+
		"LEFT JOIN (SELECT * FROM reside WHERE checkout = false) AS R ON H.id = R.house_id", strings.Repeat(", ?", len(args)-1))
	rows, err := db.Query(stmt, args...)
	if err != nil {
		fmt.Println("QueryHouse error: ", err)
	}
	var examined []Reside

	for rows.Next() {
		var house NullableReside
		if err := rows.Scan(&house.HouseID, &house.Level, &house.FamilyID); err != nil {
			log.Panic(err)
		}
		if house.FamilyID.Valid {
			examined = append(examined, Reside{HouseID: house.HouseID.Int32, FamilyID: house.FamilyID.Int32, Level: house.Level})
		} else {
			examined = append(examined, Reside{HouseID: house.HouseID.Int32, FamilyID: 0, Level: house.Level})
		}
	}
	return examined
}

func (db *DB) BatchInsertMatches(resides Resides) Resides {
	args := make([]interface{}, len(resides))
	deletedHouses := Resides{}
	restHouses := Resides{}
	var id int32
	for i, reside := range resides {
		args[i] = reside.HouseID
	}
	stmt := fmt.Sprintf("SELECT id FROM house WHERE id IN (?%s) AND deleted = true", strings.Repeat(", ?", len(resides)-1))
	rows, err := db.Query(stmt, args...)
	if err != nil {
		fmt.Println("BatchInsertMatches[manager]: query deleted houses fail, ", err)
	}
	idx := 0
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			fmt.Println("BatchInsertMatches[manager]: scan the query result of deleted houses fail, ", err)
		}
		for resides[idx].HouseID < id {
			restHouses = append(restHouses, resides[idx])
			idx++
		}
		if resides[idx].HouseID == id {
			deletedHouses = append(deletedHouses, resides[idx])
			idx++
		}
	}
	restHouses = append(restHouses, resides[idx:]...)

	// mysql do not support (full) outer join so refer https://stackoverflow.com/questions/4796872/how-to-do-a-full-outer-join-in-mysql
	//rows, _ = tx.Query(`SELECT H.id, R.family_id FROM (SELECT * FROM house where id = ? AND deleted = false) as H
	//					LEFT JOIN (SELECT * FROM reside WHERE checkout = false) AS R ON H.id = R.house_id
	//					UNION
	//					SELECT H.id, R.family_id FROM (SELECT * FROM house where deleted = false) as H
	//					RIGHT JOIN (SELECT * FROM reside WHERE family_id = ? AND checkout = false) AS R ON H.id = R.house_id`, reside.HouseID, reside.FamilyID)
	//
	//var res MatchResponse
	//for rows.Next() {
	//	var rid NullableReside
	//	if err := rows.Scan(&rid.HouseID, &rid.FamilyID); err != nil {
	//		log.Print(err)
	//	}
	//	if rid.FamilyID.Valid {
	//		if rid.HouseID.Int32 != reside.HouseID {
	//			res.FamilyOwnHouse = true
	//		} else if rid.FamilyID.Int32 != reside.FamilyID {
	//			res.HouseMatched = true
	//		} else {
	//			res.Success = true
	//		}
	//	}
	//}
	db.batchInsertMatches(restHouses)
	return deletedHouses
}

func (db *DB) batchInsertMatches(resides Resides) {
	size := MAXARGS / ResideArgLens
	tx, _ := db.Begin()
	chunkList := funk.Chunk(resides, size)
	for _, chunk := range chunkList.([][]Reside) {
		valueStrings := make([]string, 0, len(chunk))
		valueArgs := make([]interface{}, 0, len(chunk)*ResideArgLens)

		for _, record := range chunk {
			valueStrings = append(valueStrings, "(?, ?)")
			valueArgs = append(valueArgs, record.HouseID)
			valueArgs = append(valueArgs, record.FamilyID)
		}
		stmt := fmt.Sprintf("INSERT INTO reside (house_id, family_id) VALUES %s", strings.Join(valueStrings, ","))
		_, err := tx.Exec(stmt, valueArgs...)
		if err != nil {
			fmt.Println("batchInsertMatches[manager]: ", err)
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				fmt.Println("batchInsertMatches[manager]: unable to rollback, ", rollbackErr)
			}
			break
		}
	}
	if err := tx.Commit(); err != nil {
		fmt.Println("batchInsertMatches[manager]: commit fail, ", err)
	}
}

func (db *DB) CheckOutHouse(reside Reside) error {
	_, err := db.Exec("UPDATE reside SET checkout = true WHERE house_id = ? AND family_id = ?", reside.HouseID, reside.FamilyID)
	return err
}

func (db *DB) BatchCheckOutHouses(checkouts []Reside) {
	size := MAXARGS
	tx, _ := db.Begin()
	chunkList := funk.Chunk(checkouts, size)
	for _, chunk := range chunkList.([][]Reside) {
		valueArgs := make([]interface{}, 0, len(chunk))
		for _, record := range chunk {
			valueArgs = append(valueArgs, record.HouseID)
		}
		stmt := fmt.Sprintf("UPDATE reside SET reside.checkout = true "+
			"WHERE reside.house_id IN (?%s)", strings.Repeat(",?", len(valueArgs)-1))
		_, err := tx.Exec(stmt, valueArgs...)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				fmt.Println("BatchCheckOutHouses[manager]: unable to rollback, ", rollbackErr)
			}
			fmt.Println(err)
			break
		}
	}
	err := tx.Commit()
	if err != nil {
		fmt.Println(err)
	}
}

// https://medium.com/better-programming/how-to-bulk-create-and-update-the-right-way-in-golang-part-i-e15a8e5585d1

func (db *DB) BatchInsertHouses(records []House) BatchOPRes {
	size := MAXARGS / HouseArgLens
	tx, _ := db.Begin()
	chunkList := funk.Chunk(records, size)
	var HouseID BatchOPRes
	for _, chunk := range chunkList.([][]House) {
		valueStrings := make([]string, 0, len(chunk))
		valueArgs := make([]interface{}, 0, len(chunk)*HouseArgLens)

		for _, record := range chunk {
			valueStrings = append(valueStrings, "(?, ?, ?)")
			valueArgs = append(valueArgs, record.Age)
			valueArgs = append(valueArgs, record.Area)
			valueArgs = append(valueArgs, record.Level)
		}
		stmt := fmt.Sprintf("INSERT INTO house (age, area, level) VALUES %s", strings.Join(valueStrings, ","))
		res, err := tx.Exec(stmt, valueArgs...)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				fmt.Println("BatchInsertHouses[manager]: unable to rollback, ", rollbackErr)
			}
			fmt.Println(err)
			break
		} else {
			id, _ := res.LastInsertId()
			if HouseID.Idx == 0 {
				HouseID.Idx = int(id)
			}
			HouseID.Length = append(HouseID.Length, len(chunk))
		}
	}
	err := tx.Commit()
	if err != nil {
		fmt.Println(err)
	}
	return HouseID
}
