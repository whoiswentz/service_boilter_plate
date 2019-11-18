package main

import (
	"github.com/go-redis/redis/v7"
	"log"
	"net/http"
	"os"
	"task_golang_api/handlers"
	"task_golang_api/server"
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

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	svr := server.New(Port, mux)
	log.Printf("starting the service on port: %s", Port)
	if err := svr.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
