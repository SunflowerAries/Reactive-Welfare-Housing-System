package storage

import (
	"fmt"
	"log"
	"strings"

	"github.com/thoas/go-funk"
)

func (db *DB) ClearHouse() []Reside {
	rows, err := db.Query(`SELECT house_id, family_id FROM reside WHERE checkout = true`)
	db.Exec(`DELETE FROM reside WHERE checkout = true`)
	var checkouted []Reside
	for rows.Next() {
		var reside Reside
		if err := rows.Scan(&reside.HouseID, &reside.FamilyID); err != nil {
			log.Print(err)
		}
		checkouted = append(checkouted, reside)
	}
	return checkouted
}

func (db *DB) DeleteHouse(HouseID []int32) error {
	args := make([]interface{}, len(HouseID))
	for i, id := range HouseID {
		args[i] = id
	}
	stmt := `DELETE FROM house WHERE id IN (?` + strings.Repeat(",?", len(args)-1) + `)`
	_, err := db.Exec(stmt, args...)
	return err
}

func (db *DB) DeleteReside(reside Reside) {
	db.Exec(`DELETE FROM reside WHERE house_id = ? AND family_id = ?`, reside.HouseID, reside.FamilyID)
}

func (db *DB) SignedDeleteHouse(HouseID []int32) error {
	args := make([]interface{}, len(HouseID))
	for i, id := range HouseID {
		args[i] = id
	}
	stmt := `UPDATE house LEFT JOIN reside ON house.id = reside.house_id SET house.deleted = true, reside.checkout = true WHERE house.id IN (?` + strings.Repeat(",?", len(args)-1) + `)`
	_, err := db.Exec(stmt, args...)
	return err
}

func (db *DB) QueryReside(HouseID []int32) []Reside {
	args := make([]interface{}, len(HouseID))
	for i, id := range HouseID {
		args[i] = id
	}
	stmt := `SELECT H.id, H.age, H.area, H.level, R.family_id FROM (SELECT * FROM house WHERE id IN (?` + strings.Repeat(",?", len(args)-1) + `)) AS H
	 LEFT JOIN (SELECT * FROM reside WHERE checkout = false) AS R ON H.id = R.house_id`
	rows, err := db.Query(stmt, args...)
	if err != nil {
		fmt.Println("QueryHouse error: ", err)
	}
	var examined []Reside

	for rows.Next() {
		var house NullableReside
		if err := rows.Scan(&house.HouseID, &house.Age, &house.Area, &house.Level, &house.FamilyID); err != nil {
			log.Panic(err)
		}
		if house.FamilyID.Valid {
			examined = append(examined, Reside{HouseID: house.HouseID.Int32, Area: house.Area, Age: house.Age, FamilyID: house.FamilyID.Int32, Level: house.Level})
		} else {
			examined = append(examined, Reside{HouseID: house.HouseID.Int32, Area: house.Area, Age: house.Age, FamilyID: 0, Level: house.Level})
		}
	}
	return examined
}

func (db *DB) InsertMatch(reside Reside) (MatchResponse, error) {
	tx, _ := db.Begin()
	var id int32
	if err := tx.QueryRow(`SELECT id FROM house WHERE id = ? AND deleted = false`, reside.HouseID).Scan(&id); err != nil {
		return MatchResponse{}, err
	}
	// mysql do not support (full) outer join so refer https://stackoverflow.com/questions/4796872/how-to-do-a-full-outer-join-in-mysql
	rows, _ := tx.Query(`SELECT H.id, R.family_id FROM (SELECT * FROM house where id = ? AND deleted = false) as H 
						LEFT JOIN (SELECT * FROM reside WHERE checkout = false) AS R ON H.id = R.house_id
						UNION
						SELECT H.id, R.family_id FROM (SELECT * FROM house where deleted = false) as H
						RIGHT JOIN (SELECT * FROM reside WHERE family_id = ? AND checkout = false) AS R ON H.id = R.house_id`, reside.HouseID, reside.FamilyID)

	var res MatchResponse
	for rows.Next() {
		var rid NullableReside
		if err := rows.Scan(&rid.HouseID, &rid.FamilyID); err != nil {
			log.Print(err)
		}
		if rid.FamilyID.Valid {
			if rid.HouseID.Int32 != reside.HouseID {
				res.FamilyOwnHouse = true
			} else if rid.FamilyID.Int32 != reside.FamilyID {
				res.HouseMatched = true
			} else {
				res.Success = true
			}
		}
	}
	if res.FamilyOwnHouse || res.HouseMatched || res.Success {
		return res, nil
	}
	_, err := db.Exec("INSERT INTO reside (house_id, family_id) VALUES (?, ?)", reside.HouseID, reside.FamilyID)
	return MatchResponse{Success: true}, err
}

func (db *DB) CheckOutHouse(reside Reside) error {
	_, err := db.Exec("UPDATE reside SET checkout = true WHERE house_id = ? AND family_id = ?", reside.HouseID, reside.FamilyID)
	return err
}

// https://medium.com/better-programming/how-to-bulk-create-and-update-the-right-way-in-golang-part-i-e15a8e5585d1

func (db *DB) BatchInsertHouse(records []House) []BatchOPRes {
	size := MAXARGS / HouseArgLens
	tx, _ := db.Begin()
	chunkList := funk.Chunk(records, size)
	HouseID := []BatchOPRes{}
	for _, chunk := range chunkList.([][]House) {
		valueStrings := make([]string, 0, len(chunk))
		valueArgs := make([]interface{}, 0, len(chunk)*3)

		for _, record := range chunk {
			valueStrings = append(valueStrings, "(?, ?, ?)")
			valueArgs = append(valueArgs, record.Age)
			valueArgs = append(valueArgs, record.Area)
			valueArgs = append(valueArgs, record.Level)
		}
		stmt := fmt.Sprintf("INSERT INTO house (age, area, level) VALUES %s", strings.Join(valueStrings, ","))
		res, err := tx.Exec(stmt, valueArgs...)
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
			break
		} else {
			id, _ := res.LastInsertId()
			HouseID = append(HouseID, BatchOPRes{Idx: int(id), Length: len(chunk)})
		}
	}
	err := tx.Commit()
	if err != nil {
		fmt.Println(err)
	}
	return HouseID
}
