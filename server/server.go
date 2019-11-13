package server

import (
	"net/http"
	"time"
	"config"
)

func New(port string, mux *http.ServeMux) *http.Server {
	tlsConfig := config.NewTlsConfig()

	return &http.Server{
		Addr:         port,
		Handler:      mux,
		TLSConfig:    tlsConfig,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}
