package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type DataBase struct {
	db *sql.DB
}

var connStr = "user=postgres password=gHsdasx12Adx dbname=wb sslmode=disable"

func DbStruct() *DataBase {

	database, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	return &DataBase{db: database}
}

func (s *DataBase) DbCreate() {
	_, err := s.db.Exec("CREATE TABLE IF NOT EXISTS Orders(data JSONB)")
	if err != nil {
		log.Fatal(err)
	}

}

func (s *DataBase) DbInsertData(order []byte) {

	_, err := s.db.Exec("INSERT INTO Orders VALUES($1)", order)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *DataBase) DbGetAll() *sql.Rows {
	rows, err := s.db.Query("SELECT * FROM Orders")
	if err != nil {
		log.Fatal(err)
	}

	return rows

}
