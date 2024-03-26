package repository

type TransctionRepo struct {
	*BaseRepo
}

type ITransactionRepo interface {
	Create(model interface{}, tableMame string) error
	First(searchModel interface{}, destinationModel interface{}) error
}

func NewTransactionRepo(baseRepo *BaseRepo) ITransactionRepo {
	return &TransctionRepo{baseRepo}
}
