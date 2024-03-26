package model

import "time"

type Transaction struct {
	Id              int       `json:"id" gorm:"primaryKey"`
	AccountId       int       `json:"account_id"`
	OperationTypeId int       `json:"operation_type_id"`
	Amount          float64   `json:"amount"`
	EventDate       time.Time `json:"event_date" gorm:"default:CURRENT_TIMESTAMP()"`
}

const TableTransaction = "transactions"
