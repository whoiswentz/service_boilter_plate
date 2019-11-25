package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"os"
	"time"
)

type HealthState struct {
	State         int
	ErrorMessages []string
}

func (h Handler) healthCheck(w http.ResponseWriter, r *http.Request) {
	logFlags := log.LstdFlags | log.Lshortfile
	logger := log.New(os.Stdout, "HEALTH CHECK - ", logFlags)

	healthState := &HealthState{}

	_, err := h.cache.Ping().Result()
	if err != nil {
		msg := fmt.Sprintf("Redis: %s", err.Error())
		healthState.ErrorMessages = append(healthState.ErrorMessages, msg)
		logger.Fatalln(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := h.db.Ping(ctx, readpref.Primary()); err != nil {
		msg := fmt.Sprintf("Redis: %s", err.Error())
		healthState.ErrorMessages = append(healthState.ErrorMessages, msg)
		logger.Fatalln(err)
	}

	if len(healthState.ErrorMessages) > 0 {
		healthState.State = http.StatusServiceUnavailable
	} else {
		healthState.State = http.StatusOK
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(healthState.State)

	if err := json.NewEncoder(w).Encode(healthState); err != nil {
		logger.Fatalln(err)
	}
}
