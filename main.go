package main

import (
	"log"
	"net/http"
	"os"
	"todo-api/db"
	"todo-api/handlers"
	"todo-api/server"
)

//var	 (
//	Port = os.Getenv("SERVICE_PORT")
//)

const Port = ":8080"

func main() {
	logFlags := log.LstdFlags | log.Lshortfile
	logger := log.New(os.Stdout, "API - ", logFlags)

	mux := http.NewServeMux()

	db := db.OpenConnection("", "")

	handler := handlers.NewHandler(logger, db)
	handler.SetupRoutes(mux)
	svr := server.New(Port, mux)

	log.Printf("starting the service on port: %s", Port)
	if err := svr.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
