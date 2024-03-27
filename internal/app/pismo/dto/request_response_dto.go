package dto

type AccountCreateRequest struct {
	DocumentNumber string `json:"document_number" binding:"required"`
}

type AccountCreateResponse struct {
	AccountId int
}

type AccountGetResponse struct {
	AccountNumber  int
	DocumentNumber string
}
