package storage

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/thoas/go-funk"
)

// dependency injection with interface https://www.alexedwards.net/blog/organising-database-access

type House struct {
	ID    int32
	Level int32
	Age   int32
	Area  int32
}

type Family struct {
	ID     int32
	Level  int32
	Income int32
}

type FamilyCheckOut struct {
	FamilyID int32
	Level    int32
}

type Reside struct {
	HouseID  int32
	FamilyID int32
	Level    int32
	Age      int32
	Area     int32
	CheckOut bool
}

type NullableReside struct {
	HouseID  sql.NullInt32
	Level    int32
	Age      int32
	Area     int32
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

const MAXARGS = 65535

const HouseArgs = "house.id, house.age, house.area, house.level"

type MatchResponse struct {
	FamilyOwnHouse bool
	HouseMatched   bool
	Success        bool
}

type HouseSystem interface {
	InitMatchCache() ([HouseLevel + 1][]Reside, [HouseLevel + 1][]Reside)
	BatchInsertHouse(records []House) []int
	InsertMatch(reside Reside) (MatchResponse, error)
	CheckOutHouse(reside Reside) error
	QueryReside(HouseID []int32) []Reside
	QueryHouse(HouseID []int32) []House
	SignedDeleteHouse(HouseID []int32) error
	DeleteHouse(HouseID []int32) error
	DeleteReside(reside Reside)
	ClearHouse() []Reside
}

type DB struct {
	*sql.DB
}

func RemoveReside(s []Reside, i int) ([]Reside, Reside) {
	var val Reside = s[i]
	s[i] = s[len(s)-1]
	return s[:len(s)-1], val
}

func BatchInsertFamily(db *sql.DB, records []Family) {
	size := MAXARGS
	tx, _ := db.Begin()
	chunkList := funk.Chunk(records, size)
	for _, chunk := range chunkList.([][]Family) {
		valueStrings := make([]string, 0, len(chunk))
		valueArgs := make([]interface{}, 0, len(chunk))
		for _, record := range chunk {
			valueStrings = append(valueStrings, "(?)")
			valueArgs = append(valueArgs, record.Income)
		}
		stmt := fmt.Sprintf("INSERT INTO family (income) VALUES %s", strings.Join(valueStrings, ","))
		_, err := tx.Exec(stmt, valueArgs...)
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
		}
	}
	err := tx.Commit()
	if err != nil {
		fmt.Println(err)
	}
}

func BatchInsertReside(db *sql.DB, records []Reside) {
	size := MAXARGS / 2
	tx, _ := db.Begin()
	chunkList := funk.Chunk(records, size)
	for _, chunk := range chunkList.([][]Reside) {
		valueStrings := make([]string, 0, len(chunk))
		valueArgs := make([]interface{}, 0, len(chunk)*2)
		for _, record := range chunk {
			valueStrings = append(valueStrings, "(?, ?)")
			valueArgs = append(valueArgs, record.HouseID)
			valueArgs = append(valueArgs, record.FamilyID)
		}
		stmt := fmt.Sprintf("INSERT INTO reside (house_id, family_id) VALUES %s", strings.Join(valueStrings, ","))
		_, err := tx.Exec(stmt, valueArgs...)
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
		}
	}
	err := tx.Commit()
	if err != nil {
		fmt.Println(err)
	}
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
