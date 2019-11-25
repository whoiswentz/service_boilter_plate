package handlers

import (
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	cache  *redis.Client
	db     *mongo.Client
}

// Kind of dependency injection
func NewHandler(db *mongo.Client, cache *redis.Client) *Handler {
	return &Handler{db: db, cache: cache}
}
