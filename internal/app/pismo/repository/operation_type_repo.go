package repository

type OperationTypeRepo struct {
	*BaseRepo
}

type IOperationTypeRepo interface {
	Create(model interface{}, tableMame string) error
	First(searchModel interface{}, destinationModel interface{}) error
}

func NewOperationTypeRepo(baseRepo *BaseRepo) IOperationTypeRepo {
	return &OperationTypeRepo{baseRepo}
}
