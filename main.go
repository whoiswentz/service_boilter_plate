package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"task/db"
	"task/handlers"
	"task/server"
)

var (
	Port = os.Getenv("SERVICE_PORT")
)

func main() {
	logFlags := log.LstdFlags | log.Lshortfile
	logger := log.New(os.Stdout, "API - ", logFlags)

	mongo := db.NewMongoConnection("task")
	redis := db.NewRedisConnection("redis")
	handler := handlers.NewHandler(mongo, redis)

	mux := http.NewServeMux()
	handler.SetupRoutes(mux)

	svr := server.New(fmt.Sprintf(":%v", Port), mux)
	logger.Printf("starting the service on port: %s", Port)
	if err := svr.ListenAndServe(); err != nil {
		logger.Fatalln(err)
	}
}
