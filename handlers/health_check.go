package handlers

import (
	"net/http"
)

func (h Handler) healthCheck(w http.ResponseWriter, r *http.Request) {
	defer h.logger.Printf("%s - %s", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("System's up\n"))
}
