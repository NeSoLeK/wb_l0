package main

import (
	"wb_l0/db"
	"wb_l0/service"
)

func main() {
	sql := db.DbStruct()
	cache := service.CreateCacheMap(sql)
	service.CacheFill(sql, cache)
	nats := service.NatsCreateConnection(sql, cache)
	nats.Subscribe()
	sql.DbCreate()

	for {

	}
}
