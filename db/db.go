package db

import (
	"database/sql"
	"log"
)

func OpenConnection(driverName, dataSourceName string) *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}
	CheckConnection(db)
	return db
}

func CheckConnection(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}
