package rest

import "github.com/gin-gonic/gin"

func RegisterRoutes(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	contaier := NewContainer()

	v1 := routerGroup.Group("v1")

	healthCheckRoutes(v1, &contaier)
	AccountRoutes(v1, &contaier)
	TransactionRoutes(v1, &contaier)
	OperationTypeRoutes(v1, &contaier)

	return v1

}

func healthCheckRoutes(routerGroup *gin.RouterGroup, container *Container) {
	routerGroup.GET("/health", container.healthController.HealthStatus)
}

func TransactionRoutes(routerGroup *gin.RouterGroup, container *Container) {
	routerGroup.POST("/transactions", container.transactionController.CreateTransaction)
}

func AccountRoutes(routerGroup *gin.RouterGroup, container *Container) {
	routerGroup.POST("/accounts", container.accountController.CreateAccount)
	routerGroup.GET("/accounts/:account_id", container.accountController.GetAccount)
}

func OperationTypeRoutes(routerGroup *gin.RouterGroup, container *Container) {
	routerGroup.POST("/operations", container.operationTypeController.CreateOperationType)
}
