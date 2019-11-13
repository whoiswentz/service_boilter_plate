package main

import (
	"log"
	"net/http"
	"os"
	"service_boilter_plate/handlers/health_check"
	"service_boilter_plate/server"
)

//var	 (
//	Port = os.Getenv("SERVICE_PORT")
//)

const Port = ":8080"

func main() {
	logFlags := log.LstdFlags | log.Lshortfile
	logger := log.New(os.Stdout, "API", logFlags)

	mux := http.NewServeMux()

	hcHandler := health_check.NewHandler(logger)
	hcHandler.SetupRoutes(mux)

	svr := server.New(Port, mux)

	log.Printf("starting the service on port: %s", Port)
	if err := svr.ListenAndServeTLS("", ""); err != nil {
		log.Fatalln(err)
	}
}
