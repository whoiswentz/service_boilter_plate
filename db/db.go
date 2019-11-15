package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var (
	Host     = os.Getenv("DB_HOST")
	Port     = os.Getenv("DB_PORT")
	User     = os.Getenv("DB_USER")
	Password = os.Getenv("DB_PASSWORD")
	Name     = os.Getenv("DB_DATABASE")
)

func OpenConnection(driverName string) *sql.DB {
	connString := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	dataSourceName := fmt.Sprintf(connString, Host, Port, User, Password, Name)
	log.Println(dataSourceName)
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	log.Println("Successfully connected!")
	return db
}
