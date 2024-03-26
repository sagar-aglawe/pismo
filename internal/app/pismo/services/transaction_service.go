package services

import "self-projects/pismo/internal/app/pismo/repository"

type TransactionService struct {
	transactionRepo repository.ITransactionRepo
}

func NewTransactionService(transactionRepo repository.ITransactionRepo) *TransactionService {
	return &TransactionService{transactionRepo: transactionRepo}
}
