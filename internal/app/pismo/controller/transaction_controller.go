package controller

import (
	"self-projects/pismo/internal/app/pismo/services"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionService *services.TransactionService
}

type ITransactionController interface {
	CreateTransaction(ctx *gin.Context)
}

func NewTransactionController(transactionService *services.TransactionService) ITransactionController {
	return &TransactionController{transactionService: transactionService}
}

func (tc *TransactionController) CreateTransaction(ctx *gin.Context) {

}
