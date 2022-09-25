package server

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"wallet-api/api"
	"wallet-api/api/middlewares"
)

func SetupHTTPServer(routers ...api.Router) *http.Server {
	r := gin.Default()
	r.Use(middlewares.ErrorHandler)

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
