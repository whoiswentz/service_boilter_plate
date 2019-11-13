package health_check

import (
	"log"
	"net/http"
)

type Handler struct {
	logger *log.Logger
}

// Kind of dependency injection
func NewHandler(logger *log.Logger) *Handler {
	return &Handler{logger: logger}
}

func (h Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	const msg = "System's up\n"
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health-check", h.Logger(h.HealthCheck))
}

// Middleware
func (h *Handler) Logger (next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer next(w, r)
		h.logger.Println("Checking systems health")
	}
}