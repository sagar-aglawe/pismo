package controller

import (
	"net/http"
	"self-projects/pismo/internal/app/pismo/dto"
	"self-projects/pismo/internal/app/pismo/response"
	"self-projects/pismo/internal/app/pismo/services"
	"self-projects/pismo/pkg/request_context"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionService services.ITransactionService
}

type ITransactionController interface {
	CreateTransaction(ctx *gin.Context)
}

func NewTransactionController(transactionService services.ITransactionService) ITransactionController {
	return &TransactionController{transactionService: transactionService}
}

// CreateTransaction 	godoc
// @Summary 			Create Transactions
// @Description 		Create transactions for the given input
// @param 				request body dto.TransactionCreateRequest true "Create Transactions"
// @Produce 			application/json
// @Tags 				Transactions
// @Success 			200 {object} map[string]interface{}
// @Failure 			400 {object} map[string]interface{}
// @Router 				/transactions [post]
func (tc *TransactionController) CreateTransaction(ctx *gin.Context) {
	rCtx := request_context.GetRCtx(ctx)

	reqBody := dto.TransactionCreateRequest{}

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		rCtx.Log.WithError(err).Error("error while validating create-transaction request")
		response.FormatResponse(ctx, false, http.StatusBadRequest, "", err.Error())
		return
	}

	resp, err := tc.transactionService.CreateTransaction(&rCtx, &reqBody)
	if err != nil {
		rCtx.Log.WithError(err).Error("error from create-transaction service")
		response.FormatResponse(ctx, false, http.StatusBadRequest, "", err.Error())
		return
	}

	response.FormatResponse(ctx, true, http.StatusOK, resp, "")
}
