package services

import (
	"self-projects/pismo/internal/app/pismo/dto"
	"self-projects/pismo/internal/app/pismo/model"
	"self-projects/pismo/internal/app/pismo/repository"
	"self-projects/pismo/pkg/request_context"
)

type OperationTypeService struct {
	operationTypeRepo repository.IOperationTypeRepo
}

type IOperationTypeService interface {
	CreateOperationType(rCtx *request_context.ReqCtx, reqBody *dto.OperationTypeCreateRequest) (*dto.OperationTypeCreateResponse, error)
}

func NewOperationTypeService(operationTypeRepo repository.IOperationTypeRepo) IOperationTypeService {
	return &OperationTypeService{operationTypeRepo: operationTypeRepo}
}

func (os *OperationTypeService) CreateOperationType(
	rCtx *request_context.ReqCtx, reqBody *dto.OperationTypeCreateRequest) (*dto.OperationTypeCreateResponse, error) {

	operationTypeModel := model.OperationType{
		Description: reqBody.Description,
	}
	err := os.operationTypeRepo.Create(&operationTypeModel, model.TableOperationType)
	if err != nil {
		return nil, err
	}

	return &dto.OperationTypeCreateResponse{
		OperationTypeId: operationTypeModel.Id,
	}, nil
}
