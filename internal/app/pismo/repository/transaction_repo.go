package repository

type TransctionRepo struct {
	*BaseRepo
}

type ITransactionRepo interface {
	Create(model interface{}, tableMame string) error
}

func NewTransactionRepo(baseRepo *BaseRepo) ITransactionRepo {
	return &TransctionRepo{baseRepo}
}
