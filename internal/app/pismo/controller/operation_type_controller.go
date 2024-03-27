package controller

import (
	"net/http"
	"self-projects/pismo/internal/app/pismo/dto"
	"self-projects/pismo/internal/app/pismo/response"
	"self-projects/pismo/internal/app/pismo/services"
	"self-projects/pismo/pkg/request_context"

	"github.com/gin-gonic/gin"
)

type OperationTypeController struct {
	opertaionTypeService services.IOperationTypeService
}

type IOperationTypeController interface {
	CreateOperationType(ctx *gin.Context)
}

func NewOperationTypeController(operationTypeService services.IOperationTypeService) IOperationTypeController {
	return &OperationTypeController{opertaionTypeService: operationTypeService}
}

func (oc *OperationTypeController) CreateOperationType(ctx *gin.Context) {
	rCtx := request_context.GetRCtx(ctx)

	reqBody := dto.OperationTypeCreateRequest{}

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		rCtx.Log.WithError(err).Error("error while validating operation-type request")
		response.FormatResponse(ctx, false, http.StatusBadRequest, "", err.Error())
		return
	}

	resp, err := oc.opertaionTypeService.CreateOperationType(&rCtx, &reqBody)
	if err != nil {
		rCtx.Log.WithError(err).Error("error from operation-type service")
		response.FormatResponse(ctx, false, http.StatusBadRequest, "", err.Error())
		return
	}

	response.FormatResponse(ctx, true, http.StatusOK, resp, "")
}
