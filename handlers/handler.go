package handlers

import (
	"database/sql"
	"log"
)

type Handler struct {
	logger *log.Logger
	db     *sql.DB
}

// Kind of dependency injection
func NewHandler(logger *log.Logger, db *sql.DB) *Handler {
	return &Handler{logger: logger, db: db}
}
