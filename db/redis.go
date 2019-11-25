package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"os"
)

var (
	RedisPort = os.Getenv("REDIS_PORT")
)

func NewRedisConnection(hostname string) *redis.Client {
	logFlags := log.LstdFlags | log.Lshortfile
	logger := log.New(os.Stdout, "REDIS - ", logFlags)

	addr := fmt.Sprintf("%s:%v", hostname, RedisPort)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		logger.Fatalln(err)
	}
	logger.Println("Up")

	return client
}
