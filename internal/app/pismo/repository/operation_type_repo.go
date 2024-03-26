package repository

type OperationTypeRepo struct {
	*BaseRepo
}

func NewOperationTypeRepo(baseRepo *BaseRepo) *OperationTypeRepo {
	return &OperationTypeRepo{baseRepo}
}
