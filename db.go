package main

import (
	"./config"
	"./types"
)

const DBFILE = "./db/db.json"

func ServeDb(m Messaging, db map[string]types.Entity) {
	for {
		select {
		case txm := <-m.Txnnel:
			if db[txm.From].Balance < txm.Amount {
				txm.Returnel <- false
			} else {
				eto := db[txm.To]
				efrom := db[txm.From]
				eto.Balance += txm.Amount
				efrom.Balance -= txm.Amount
				db[txm.To] = eto
				db[txm.From] = efrom
				txm.Returnel <- true
			}
		case im := <-m.Inforeqnel:
			if im.string == "" {
				for k, v := range db {
					im.Infonnel <- types.Information{k, v.Balance}
				}
			} else {
				im.Infonnel <- types.Information{im.string, db[im.string].Balance}
			}
			close(im.Infonnel)
		case am := <-m.Authennel:
			am.Returnel <- IsValid(am.Authentication, db[am.From].Password, config.MAX_DELAY)
		}
	}
}
