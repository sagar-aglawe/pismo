package rest

import (
	"net/http"
	"self-projects/pismo/internal/app/pismo/middlewares"

	"github.com/gin-gonic/gin"
)

func BuildServer() *http.Server {
	handler := gin.New()
	handler.Use(middlewares.CustomLogger())

	RegisterRoutes(handler.Group("pismo/api"))

	return &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler,
	}
}
