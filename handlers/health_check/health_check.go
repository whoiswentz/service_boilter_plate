package health_check

import (
	"log"
	"net/http"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{logger: logger}
}

func (h Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	const msg = "System's up\n"

	h.logger.Println("Checking systems health")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}
