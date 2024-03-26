package repository

type AccountRepo struct {
	*BaseRepo
}

func NewAccountRepo(baseRepo *BaseRepo) *AccountRepo {
	return &AccountRepo{baseRepo}
}
