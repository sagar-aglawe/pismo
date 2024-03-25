package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuildServer() *http.Server {
	handler := gin.New()

	RegisterRoutes(handler.Group("pismo/api"))

	return &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler,
	}
}
