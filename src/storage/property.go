package storage

import (
	"fmt"
	"log"
	"strings"
)

func (db *DB) QueryHouse(HouseID []int32) []House {
	args := make([]interface{}, len(HouseID))
	for i, id := range HouseID {
		args[i] = id
	}
	stmt := fmt.Sprintf("SELECT id, age, area, level, deleted"+
		"FROM house WHERE id IN (?%s)", strings.Repeat(", ?", len(args)-1))
	rows, err := db.Query(stmt, args...)
	if err != nil {
		fmt.Println("QueryHouse[property]: query house fail, ", err)
	}
	var examined []House

	for rows.Next() {
		var house House
		if err := rows.Scan(&house.ID, &house.Age, &house.Area, &house.Level, &house.Deleted); err != nil {
			log.Panic(err)
		}
		examined = append(examined, House{ID: house.ID, Area: house.Area, Age: house.Age, Level: house.Level, Deleted: house.Deleted})
	}
	return examined
}
