package server

import (
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"

	"wallet-api/api"
)

func SetupHTTPServer(routers ...api.Router) *http.Server {
	r := chi.NewRouter()
	for _, handler := range routers {
		handler.Routes(r)
	}

	a := os.Getenv("SERVER_ADDRESS")
	zap.S().Debugf("Starting server on %s", a)

	return &http.Server{
		Addr:         a,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
}
