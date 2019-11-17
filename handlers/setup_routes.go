package handlers

import (
	"net/http"
	"service_boilter_plate/middleware"
)

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health-check",
		middleware.HandlerFunc(h.healthCheck,
			middleware.HttpMethod("GET")))

	mux.HandleFunc("/task",
		middleware.HandlerFunc(h.postTodo,
			middleware.HttpMethod("POST")))
}
