package repository

import "gorm.io/gorm"

type BaseRepo struct {
	DB *gorm.DB
}

func NewBaseRepo(db *gorm.DB) *BaseRepo {
	return &BaseRepo{
		DB: db,
	}
}

func (baseRepo *BaseRepo) Create(model interface{}, tableMame string) error {
	if tx := baseRepo.DB.Table(tableMame).Create(model); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (baseRepo *BaseRepo) First(searchModel interface{}, destinationModel interface{}) error {
	if tx := baseRepo.DB.First(destinationModel, searchModel); tx.Error != nil {
		return tx.Error
	}

	return nil
}
