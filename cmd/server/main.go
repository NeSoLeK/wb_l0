package main

import (
	"wb_l0/db"
	"wb_l0/handlers"
	"wb_l0/service"
)

func main() {
	sql := db.DbStruct()
	sql.DbCreate()
	cache := service.CreateCacheMap(sql)
	service.CacheFill(sql, cache)
	nats := service.NatsCreateConnection(sql, cache)
	nats.Subscribe()
	sql.DbCreate()
	router := handlers.CreateRouter(sql, cache)
	router.StartHandlers()
}
