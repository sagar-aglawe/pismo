package model

type Account struct {
	Id             int    `json:"id" gorm:"primaryKey"`
	DocumentNumber string `json:"document_number"`
}

const TableAccount = "accounts"
