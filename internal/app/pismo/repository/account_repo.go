package repository

type AccountRepo struct {
	*BaseRepo
}

type IAccountRepo interface {
	Create(model interface{}, tableMame string) error
	First(searchModel interface{}, destinationModel interface{}) error
}

func NewAccountRepo(baseRepo *BaseRepo) IAccountRepo {
	return &AccountRepo{baseRepo}
}
