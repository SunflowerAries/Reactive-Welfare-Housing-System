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
	stmt := `SELECT id, age, area, level FROM house WHERE id IN (?` + strings.Repeat(",?", len(args)-1) + `) AND deleted = false`
	rows, err := db.Query(stmt, args...)
	if err != nil {
		fmt.Println("QueryHouse error: ", err)
	}
	var examined []House

	for rows.Next() {
		var house House
		if err := rows.Scan(&house.ID, &house.Age, &house.Area, &house.Level); err != nil {
			log.Panic(err)
		}
		examined = append(examined, House{ID: house.ID, Area: house.Area, Age: house.Age, Level: house.Level})
	}
	return examined
}
