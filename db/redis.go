package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func NewRedisConnection(hostname string) *redis.Client {
	addr := fmt.Sprintf("%s:6379", hostname)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(pong)

	return client
}
