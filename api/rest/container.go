package rest

import (
	"self-projects/pismo/internal/app/pismo/controller"
	"self-projects/pismo/internal/app/pismo/repository"
	"self-projects/pismo/internal/app/pismo/services"
	"self-projects/pismo/internal/providers"
)

type Container struct {
	healthController      controller.IHealthController
	accountController     controller.IAccountController
	transactionController controller.ITransactionController
}

func NewContainer() Container {
	reg := providers.New()
	reg.Resolve(providers.PostgresClient)

	db := providers.GetPostgresClient()

	baseRepo := repository.NewBaseRepo(db)

	transactionRepo := repository.NewTransactionRepo(baseRepo)
	accountsRepo := repository.NewAccountRepo(baseRepo)

	transactionService := services.NewTransactionService(transactionRepo)
	accountService := services.NewAccountService(accountsRepo)

	healthController := controller.NewHealthController()
	transactionController := controller.NewTransactionController(transactionService)
	accountController := controller.NewAccountController(accountService)

	return Container{
		healthController:      healthController,
		transactionController: transactionController,
		accountController:     accountController,
	}
}
