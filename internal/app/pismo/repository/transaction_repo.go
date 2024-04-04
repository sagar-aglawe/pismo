package repository

import "self-projects/pismo/internal/app/pismo/model"

type TransctionRepo struct {
	*BaseRepo
}

type ITransactionRepo interface {
	Create(model interface{}, tableMame string) error
	FetchAllTransactionsWithAmountLessThan(accountId int, amount float64) ([]model.Transaction, error)
	Update(model interface{}) error
}

func NewTransactionRepo(baseRepo *BaseRepo) ITransactionRepo {
	return &TransctionRepo{baseRepo}
}

func (transactionRepo *TransctionRepo) FetchAllTransactionsWithAmountLessThan(accountId int, amount float64) ([]model.Transaction, error) {
	var resp []model.Transaction

	// find will not return no record found error
	tx := transactionRepo.DB.Table(model.TableTransaction).
		Where("account_id = ? AND amount < ?", accountId, amount).
		Find(&resp)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return resp, nil
}
