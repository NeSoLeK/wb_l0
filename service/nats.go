package service

import (
	"encoding/json"
	stan "github.com/nats-io/stan.go"
	"log"
	"wb_l0/db"
	"wb_l0/structs"
)

type NatsStream struct {
	conn  stan.Conn
	db    *db.DataBase
	cache *CacheMap
}

func NatsCreateConnection(database *db.DataBase, cachemap *CacheMap) *NatsStream {
	sc, err := stan.Connect("test-cluster", "hochyvwildberries", stan.NatsURL("nats://nats:4222"))
	if err != nil {
		log.Fatal(err)
	}
	return &NatsStream{conn: sc, db: database, cache: cachemap}
}

func (c *NatsStream) Subscribe() {
	_, err := c.conn.Subscribe("order", func(m *stan.Msg) {

		var data structs.Order
		err := json.Unmarshal(m.Data, &data)
		if err != nil {
			log.Fatal(err)
		}
		c.cache.cacheFolder[data.OrderUid] = m.Data
		c.db.DbInsertData(m.Data)
	})
	if err != nil {
		log.Fatal(err)
	}
}
