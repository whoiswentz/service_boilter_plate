package handlers

import (
	"net/http"
	"task/middleware"
)

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	go mux.HandleFunc("/health-check",
		middleware.HandlerFunc(h.healthCheck,
			middleware.HttpMethod("GET")))

	go mux.HandleFunc("/task",
		middleware.HandlerFunc(h.postTodo,
			middleware.HttpMethod("POST")))
}
