package service

import (
	"encoding/json"
	"log"
	"wb_l0/db"
	"wb_l0/structs"
)

type CacheMap struct {
	cacheFolder map[string][]byte
	db          *db.DataBase
}

func CreateCacheMap(base *db.DataBase) *CacheMap {
	return &CacheMap{cacheFolder: make(map[string][]byte)}
}

func CacheFill(database *db.DataBase, c *CacheMap) {
	rows := database.DbGetAll()
	for rows.Next() {

		var jsonData string
		err := rows.Scan(&jsonData)
		if err != nil {
			log.Fatal(err)
		}

		var order structs.Order
		err = json.Unmarshal([]byte(jsonData), &order)
		if err != nil {
			log.Fatal(err)
		}

		c.cacheFolder[order.OrderUid] = []byte(jsonData)

	}

}
