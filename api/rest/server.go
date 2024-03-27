package rest

import (
	"net/http"
	"self-projects/pismo/internal/app/pismo/middlewares"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func BuildServer() *http.Server {
	handler := gin.New()

	// adding swagger
	handler.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	handler.Use(middlewares.CustomLogger())

	RegisterRoutes(handler.Group("pismo/api"))

	return &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler,
	}
}
