package services

import (
	"self-projects/pismo/internal/app/pismo/dto"
	"self-projects/pismo/internal/app/pismo/model"
	"self-projects/pismo/internal/app/pismo/repository"
	"self-projects/pismo/pkg/request_context"
)

type AccountService struct {
	accountRepo repository.IAccountRepo
}

type IAccountService interface {
	CreateAccount(rCtx *request_context.ReqCtx, reqBody *dto.AccountCreateRequest) (*dto.AccountCreateResponse, error)
	GetAccount(rCtx *request_context.ReqCtx, accountId int) (*dto.AccountGetResponse, error)
}

func NewAccountService(accountRepo repository.IAccountRepo) IAccountService {
	return &AccountService{
		accountRepo: accountRepo,
	}
}

func (as *AccountService) CreateAccount(
	rCtx *request_context.ReqCtx, reqBody *dto.AccountCreateRequest) (*dto.AccountCreateResponse, error) {

	accountModel := model.Account{
		DocumentNumber: reqBody.DocumentNumber,
	}

	err := as.accountRepo.Create(&accountModel, model.TableAccount)
	if err != nil {
		return nil, err
	}

	return &dto.AccountCreateResponse{
		AccountId: accountModel.Id,
	}, nil
}

func (as *AccountService) GetAccount(rCtx *request_context.ReqCtx, accountId int) (*dto.AccountGetResponse, error) {
	searchModel := model.Account{
		Id: accountId,
	}
	var destinationModel model.Account

	err := as.accountRepo.First(searchModel, &destinationModel)
	if err != nil {
		return nil, err
	}
	return &dto.AccountGetResponse{
		AccountNumber:  destinationModel.Id,
		DocumentNumber: destinationModel.DocumentNumber,
	}, nil
}
