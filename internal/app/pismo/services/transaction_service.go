package services

import (
	"errors"
	"self-projects/pismo/internal/app/pismo/dto"
	"self-projects/pismo/internal/app/pismo/model"
	"self-projects/pismo/internal/app/pismo/repository"
	"self-projects/pismo/pkg/request_context"
)

type TransactionService struct {
	transactionRepo   repository.ITransactionRepo
	accountRepo       repository.IAccountRepo
	operationTypeRepo repository.IOperationTypeRepo
}

type ITransactionService interface {
	CreateTransaction(rCtx *request_context.ReqCtx, reqBody *dto.TransactionCreateRequest) (*dto.TransactionCreateResponse, error)
}

func NewTransactionService(
	transactionRepo repository.ITransactionRepo,
	accountRepo repository.IAccountRepo,
	operationTypeRepo repository.IOperationTypeRepo) ITransactionService {
	return &TransactionService{
		transactionRepo:   transactionRepo,
		accountRepo:       accountRepo,
		operationTypeRepo: operationTypeRepo,
	}
}

func (as *TransactionService) CreateTransaction(
	rCtx *request_context.ReqCtx, reqBody *dto.TransactionCreateRequest) (*dto.TransactionCreateResponse, error) {

	accountModel := model.Account{}
	accountSearchModel := model.Account{Id: reqBody.AccountId}
	err := as.accountRepo.First(accountSearchModel, &accountModel)
	if err != nil {
		rCtx.Log.WithError(err).WithField("account_id", reqBody.AccountId).Error("no account present for given account_id")
		return nil, errors.New("no account associated with given account_id")
	}

	operationTypeModel := model.OperationType{}
	operationTypeSearchModel := model.OperationType{Id: reqBody.OperationTypeId}
	err = as.operationTypeRepo.First(operationTypeSearchModel, &operationTypeModel)
	if err != nil {
		rCtx.Log.WithError(err).WithField("operation_type_id", reqBody.OperationTypeId).Error("no operations present for given operation type id")
		return nil, errors.New("no operations associated with given operation_type_id")
	}

	transactionModel := model.Transaction{
		AccountId:       reqBody.AccountId,
		OperationTypeId: reqBody.OperationTypeId,
		Amount:          reqBody.Amount,
	}

	err = as.transactionRepo.Create(&transactionModel, model.TableTransaction)
	if err != nil {
		return nil, err
	}

	return &dto.TransactionCreateResponse{
		TransactionId: transactionModel.Id,
	}, nil
}
