package main

import (
	"log"
	"net/http"
	"os"
	"service_boilter_plate/handlers"
	"service_boilter_plate/server"
)

var (
	Port = os.Getenv("SERVICE_PORT")
)

func main() {
	logFlags := log.LstdFlags | log.Lshortfile
	logger := log.New(os.Stdout, "API - ", logFlags)

	mux := http.NewServeMux()
	//conn := db.OpenConnection("postgres")
	handler := handlers.NewHandler(logger, nil)
	handler.SetupRoutes(mux)

	if Port == "" {
		Port = ":8080"
	}

	svr := server.New(Port, mux)
	log.Printf("starting the service on port: %s", Port)
	if err := svr.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
