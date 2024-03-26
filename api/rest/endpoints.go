package rest

import "github.com/gin-gonic/gin"

func RegisterRoutes(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	contaier := NewContainer()

	v1 := routerGroup.Group("v1")

	healthCheckRoutes(v1, &contaier)

	return v1

}

func healthCheckRoutes(routerGroup *gin.RouterGroup, container *Container) {
	routerGroup.GET("/health", container.healthController.HealthStatus)
}

func TransactionRoutes(routerGroup *gin.RouterGroup, container *Container) {
}

func AccountRoutes(routerGroup *gin.RouterGroup, container *Container) {
	
}
