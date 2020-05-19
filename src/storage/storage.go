package storage

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// dependency injection with interface https://www.alexedwards.net/blog/organising-database-access

type House struct {
	ID    uint
	Level int32
	Age   int32
	Area  int32
}

type Family struct {
	ID     uint
	Income int32
}

type Reside struct {
	HouseID  uint
	Level    uint
	FamilyID uint
}

type NullableReside struct {
	HouseID  uint
	Level    uint
	FamilyID sql.NullInt32
}

const (
	USERNAME = "housing"
	PASSWORD = "housing@2020"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "housing"
)

const HouseLevel = 3

type HouseSystem interface {
	InitMatchCache() ([]Reside, [HouseLevel + 1][]uint)
	BatchInsertHouse(records []House)
}

type DB struct {
	*sql.DB
}

func (db *DB) InitMatchCache() ([]Reside, [HouseLevel + 1][]uint) {
	rows, err := db.Query(`SELECT house.id, house.level, reside.family_id FROM house LEFT JOIN reside ON house.id = reside.house_id`)
	if err != nil {
		log.Panic(err)
	}
	var occupied []Reside
	var vacant [HouseLevel + 1][]uint
	for rows.Next() {
		var reside NullableReside
		if err := rows.Scan(&reside.HouseID, &reside.Level, &reside.FamilyID); err != nil {
			log.Panic(err)
		}
		if reside.FamilyID.Valid {
			occupied = append(occupied, Reside{HouseID: reside.HouseID, Level: reside.Level, FamilyID: uint(reside.FamilyID.Int32)})
		} else {
			vacant[reside.Level] = append(vacant[reside.Level], reside.HouseID)
		}
	}
	return occupied, vacant
}

func (db *DB) BatchInsertHouse(records []House) {
	valueStrings := make([]string, 0, len(records))
	valueArgs := make([]interface{}, 0, len(records)*3)
	for _, record := range records {
		valueStrings = append(valueStrings, "(?, ?, ?)")
		valueArgs = append(valueArgs, record.Age)
		valueArgs = append(valueArgs, record.Area)
		valueArgs = append(valueArgs, record.Level)
	}
	stmt := fmt.Sprintf("INSERT INTO house (age, area, level) VALUES %s", strings.Join(valueStrings, ","))
	sql, _ := db.Prepare(stmt)
	sql.Exec(valueArgs...)
}

func BatchInsertFamily(db *sql.DB, records []Family) {
	valueStrings := make([]string, 0, len(records))
	valueArgs := make([]interface{}, 0, len(records))
	for _, record := range records {
		valueStrings = append(valueStrings, "(?)")
		valueArgs = append(valueArgs, record.Income)
	}
	stmt := fmt.Sprintf("INSERT INTO family (income) VALUES %s", strings.Join(valueStrings, ","))
	db.Exec(stmt, valueArgs...)
}

func BatchInsertReside(db *sql.DB, records []Reside) {
	valueStrings := make([]string, 0, len(records))
	valueArgs := make([]interface{}, 0, len(records)*2)
	for _, record := range records {
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, record.HouseID)
		valueArgs = append(valueArgs, record.FamilyID)
	}
	stmt := fmt.Sprintf("INSERT INTO reside (house_id, family_id) VALUES %s", strings.Join(valueStrings, ","))
	db.Exec(stmt, valueArgs...)
}

func InitDB() (*DB, error) {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}
	// Houses := []House{
	// 	{Age: 1, Area: 100, Level: 3},
	// 	{Age: 2, Area: 80, Level: 2},
	// 	{Age: 3, Area: 60, Level: 1},
	// }
	// fmt.Printf("%+v\n", Houses)
	// myDB := &DB{db}
	// myDB.BatchInsertHouse(Houses)
	// Familys := []Family{
	// 	{Income: 1000},
	// 	{Income: 800},
	// }
	// BatchInsertFamily(db, Familys)
	// Resides := []Reside{
	// 	{HouseID: 1, FamilyID: 2},
	// 	{HouseID: 2, FamilyID: 1},
	// }
	// BatchInsertReside(db, Resides)
	return &DB{db}, nil
}
