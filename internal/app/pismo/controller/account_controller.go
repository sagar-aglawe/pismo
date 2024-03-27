package controller

import (
	"net/http"
	"self-projects/pismo/internal/app/pismo/dto"
	"self-projects/pismo/internal/app/pismo/response"
	"self-projects/pismo/internal/app/pismo/services"
	"strconv"

	"self-projects/pismo/pkg/request_context"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountService services.IAccountService
}

type IAccountController interface {
	CreateAccount(ctx *gin.Context)
	GetAccount(ctx *gin.Context)
}

func NewAccountController(accountService services.IAccountService) IAccountController {
	return &AccountController{accountService: accountService}
}

func (ac *AccountController) CreateAccount(ctx *gin.Context) {
	rCtx := request_context.GetRCtx(ctx)

	reqBody := dto.AccountCreateRequest{}

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		rCtx.Log.WithError(err).Error("error while validating create-account request")
		response.FormatResponse(ctx, false, http.StatusBadRequest, "", err.Error())
		return
	}

	resp, err := ac.accountService.CreateAccount(&rCtx, &reqBody)
	if err != nil {
		rCtx.Log.WithError(err).Error("error from create-account service")
		response.FormatResponse(ctx, false, http.StatusBadRequest, "", err.Error())
		return
	}

	response.FormatResponse(ctx, true, http.StatusOK, resp, "")

}

func (ac *AccountController) GetAccount(ctx *gin.Context) {
	rCtx := request_context.GetRCtx(ctx)
	sAccountId := ctx.Param("account_id")
	accountId, err := strconv.Atoi(sAccountId)
	if err != nil {
		rCtx.Log.WithError(err).Error("error while converting account id to int")
		response.FormatResponse(ctx, false, http.StatusBadRequest, "", err.Error())
		return
	}

	resp, err := ac.accountService.GetAccount(&rCtx, accountId)
	if err != nil {
		rCtx.Log.WithError(err).Error("error while fetching account information")
		response.FormatResponse(ctx, false, http.StatusBadRequest, "", err.Error())
		return
	}

	response.FormatResponse(ctx, true, http.StatusOK, resp, "")

}
