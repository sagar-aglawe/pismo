package dto

type AccountCreateRequest struct {
	DocumentNumber string `json:"document_number" binding:"required"`
}

type AccountCreateResponse struct {
	AccountId int `json:"account_id"`
}

type AccountGetResponse struct {
	AccountNumber  int    `json:"account_number"`
	DocumentNumber string `json:"document_number"`
}

type TransactionCreateRequest struct {
	AccountId       int     `json:"account_id" binding:"required"`
	OperationTypeId int     `json:"operation_type_id" binding:"required"`
	Amount          float64 `json:"amount" binding:"required"`
}

type TransactionCreateResponse struct {
	TransactionId int `json:"transaction_id"`
}
