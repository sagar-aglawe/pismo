package model

type OperationType struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Description string `json:"description"`
}

const TableOperationType = "operation_types"
