package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	connection := "host=192.168.39.254 user=postgres password=gaia@2020 dbname=web_store sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func PingDb() {
	db := ConnectDatabase()
	defer db.Close()
}
