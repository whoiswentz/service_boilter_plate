package handlers

import (
	"net/http"
	"todo-api/middleware"
)

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health-check",
		middleware.HandlerFunc(h.healthCheck,
			middleware.HttpMethod("GET")))
}
