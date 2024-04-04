package services

import (
	"errors"
	"fmt"
	"self-projects/pismo/internal/app/pismo/dto"
	"self-projects/pismo/internal/app/pismo/model"
	"self-projects/pismo/internal/app/pismo/repository"
	"self-projects/pismo/pkg/request_context"
	"sort"
)

type TransactionService struct {
	transactionRepo   repository.ITransactionRepo
	accountRepo       repository.IAccountRepo
	operationTypeRepo repository.IOperationTypeRepo
}

type ITransactionService interface {
	CreateTransaction(rCtx *request_context.ReqCtx, reqBody *dto.TransactionCreateRequest) (*dto.TransactionCreateResponse, error)
}

func NewTransactionService(
	transactionRepo repository.ITransactionRepo,
	accountRepo repository.IAccountRepo,
	operationTypeRepo repository.IOperationTypeRepo) ITransactionService {
	return &TransactionService{
		transactionRepo:   transactionRepo,
		accountRepo:       accountRepo,
		operationTypeRepo: operationTypeRepo,
	}
}

func (as *TransactionService) CreateTransaction(
	rCtx *request_context.ReqCtx, reqBody *dto.TransactionCreateRequest) (*dto.TransactionCreateResponse, error) {


	// accountId 
	accountModel := model.Account{}
	accountSearchModel := model.Account{Id: reqBody.AccountId}
	err := as.accountRepo.First(accountSearchModel, &accountModel)
	if err != nil {
		rCtx.Log.WithError(err).WithField("account_id", reqBody.AccountId).Error("no account present for given account_id")
		return nil, errors.New("no account associated with given account_id")
	}

	operationTypeModel := model.OperationType{}
	operationTypeSearchModel := model.OperationType{Id: reqBody.OperationTypeId}
	err = as.operationTypeRepo.First(operationTypeSearchModel, &operationTypeModel)
	if err != nil {
		rCtx.Log.WithError(err).WithField("operation_type_id", reqBody.OperationTypeId).Error("no operations present for given operation type id")
		return nil, errors.New("no operations associated with given operation_type_id")
	}

	// fetch all the transactions for the user which are in negative balance
	negativeBalanceTransactions, err := as.transactionRepo.FetchAllTransactionsWithAmountLessThan(reqBody.AccountId, 0)
	if err != nil {
		rCtx.Log.WithError(err).Error("failed to fetch the negative transactions")
		return nil, err
	}
	// group the transactions with the primary key so that update will be easier  not required will override balance
	// sort the transactions with the amount in asc so that we will settle the earlier transaction
	sort.Slice(negativeBalanceTransactions, func(i int, j int) bool {
		return negativeBalanceTransactions[i].Balance > negativeBalanceTransactions[j].Balance
	})

	var transctionsRequiredUpdate []model.Transaction

	pendingBalance := reqBody.Amount

	/*
		100
								{AccountId: accountId, Amount: -10.2},
								{AccountId: accountId, Amount: -30.2},
								{AccountId: accountId, Amount: -100.2},
	*/

	for _, negativeBalanceTransaction := range negativeBalanceTransactions {
		if pendingBalance+negativeBalanceTransaction.Balance > 0 {
			if pendingBalance > -1*(negativeBalanceTransaction.Balance) {
				pendingBalance = pendingBalance + negativeBalanceTransaction.Balance
				negativeBalanceTransaction.Balance = 0
			} else {
				negativeBalanceTransaction.Balance = negativeBalanceTransaction.Balance + pendingBalance
				pendingBalance = 0
			}
			transctionsRequiredUpdate = append(transctionsRequiredUpdate, negativeBalanceTransaction)

		}
		fmt.Println("************")
		fmt.Println(negativeBalanceTransaction.Amount, negativeBalanceTransaction.Balance)
		fmt.Println("************")
	}

	fmt.Println("balance is updated")
	fmt.Println(pendingBalance)
	fmt.Println("balance is updated")
	// accordingly update the transactions with the balance for only transctionsRequiredUpdate
	// now need to implement all these operations in the db Transaction

	for _, updateRequiredTransaction := range transctionsRequiredUpdate {
		_ = as.transactionRepo.Update(updateRequiredTransaction)
	}

	// make the entry of the new transaction in the system with pending balance
	transactionModel := model.Transaction{
		AccountId:       reqBody.AccountId,
		OperationTypeId: reqBody.OperationTypeId,
		Amount:          reqBody.Amount,
		Balance:         pendingBalance,
	}

	err = as.transactionRepo.Create(&transactionModel, model.TableTransaction)
	if err != nil {
		return nil, err
	}

	return &dto.TransactionCreateResponse{
		TransactionId: transactionModel.Id,
	}, nil
}
