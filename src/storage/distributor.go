package storage

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func (db *DB) InitMatchCache() ([HouseLevel + 1][]Reside, [HouseLevel + 1][]Reside) {
	rows, err := db.Query(`SELECT H.id, H.level, R.family_id FROM (SELECT * FROM house WHERE deleted
		= false) AS H LEFT JOIN (SELECT * FROM reside WHERE checkout = false) AS R ON H.id = R.house_id`)
	if err != nil {
		fmt.Println(err)
		log.Panic(err)
	}
	var rented [HouseLevel + 1][]Reside
	var vacant [HouseLevel + 1][]Reside
	for rows.Next() {
		var reside NullableReside
		if err := rows.Scan(&reside.HouseID, &reside.Level, &reside.FamilyID); err != nil {
			log.Panic(err)
		}
		if reside.FamilyID.Valid {
			rented[reside.Level] = append(rented[reside.Level], Reside{HouseID: reside.HouseID.Int32, FamilyID: reside.FamilyID.Int32, Level: reside.Level})
		} else {
			vacant[reside.Level] = append(vacant[reside.Level], Reside{HouseID: reside.HouseID.Int32, Level: reside.Level})
		}
	}
	return rented, vacant
}
