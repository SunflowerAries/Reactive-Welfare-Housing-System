package storage

import (
	"Reactive-Welfare-Housing-System/src/config"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func (db *DB) InitMatchCache() (map[int32]Reside, [config.HouseLevel + 1][]Reside) {
	rows, err := db.Query(`SELECT H.id, H.level, R.family_id FROM (SELECT * FROM house WHERE deleted
		= false) AS H LEFT JOIN (SELECT * FROM reside WHERE checkout = false) AS R ON H.id = R.house_id`)
	if err != nil {
		fmt.Println(err)
		log.Panic(err)
	}
	occupied := make(map[int32]Reside)
	vacant := [config.HouseLevel + 1][]Reside{}
	for rows.Next() {
		var reside NullableReside
		if err := rows.Scan(&reside.HouseID, &reside.Level, &reside.FamilyID); err != nil {
			log.Panic(err)
		}
		if reside.FamilyID.Valid {
			if value, ok := occupied[reside.FamilyID.Int32]; !ok {
				occupied[reside.FamilyID.Int32] = Reside{HouseID: reside.HouseID.Int32, FamilyID: reside.FamilyID.Int32, Level: reside.Level}
			} else {
				fmt.Printf("Family[%d] already has house[%d]", reside.FamilyID.Int32, value.HouseID)
			}
		} else {
			vacant[reside.Level] = append(vacant[reside.Level], Reside{HouseID: reside.HouseID.Int32, Level: reside.Level})
		}
	}
	return occupied, vacant
}
