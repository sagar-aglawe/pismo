package services

import (
	"self-projects/pismo/internal/app/pismo/dto"
	"self-projects/pismo/internal/app/pismo/model"
	"self-projects/pismo/internal/app/pismo/repository"
	"self-projects/pismo/pkg/request_context"
)

type TransactionService struct {
	transactionRepo repository.ITransactionRepo
}

type ITransactionService interface {
	CreateTransaction(rCtx *request_context.ReqCtx, reqBody *dto.TransactionCreateRequest) (*dto.TransactionCreateResponse, error)
}

func NewTransactionService(transactionRepo repository.ITransactionRepo) ITransactionService {
	return &TransactionService{transactionRepo: transactionRepo}
}

func (as *TransactionService) CreateTransaction(
	rCtx *request_context.ReqCtx, reqBody *dto.TransactionCreateRequest) (*dto.TransactionCreateResponse, error) {

	// todo: add logi for validation of operation type
	transactionModel := model.Transaction{
		AccountId:       reqBody.AccountId,
		OperationTypeId: reqBody.OperationTypeId,
		Amount:          reqBody.Amount,
	}

	err := as.transactionRepo.Create(&transactionModel, model.TableTransaction)
	if err != nil {
		return nil, err
	}

	return &dto.TransactionCreateResponse{
		TransactionId: transactionModel.Id,
	}, nil
}
