package services_test

import (
	"net/http/httptest"
	"self-projects/pismo/internal/app/pismo/dto"
	"self-projects/pismo/internal/app/pismo/model"
	repo_mock "self-projects/pismo/internal/app/pismo/repository/mocks"
	"self-projects/pismo/internal/app/pismo/services"
	"self-projects/pismo/pkg/logger"
	"self-projects/pismo/pkg/request_context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

/*
// func TestServiceTransaction(t *testing.T) {
// 	log := logger.GetLogger()

// 	type testCase struct {
// 		name              string
// 		mock              func(transactionRepo *repo_mock.ITransactionRepo, accountRepo *repo_mock.IAccountRepo, operationTypeRepo *repo_mock.IOperationTypeRepo)
// 		input             *dto.TransactionCreateRequest
// 		actualResponse    *dto.TransactionCreateResponse
// 		actualErrorString string
// 	}

// 	tests := []testCase{
// 		{
// 			name: "success",
// 			mock: func(transactionRepo *repo_mock.ITransactionRepo, accountRepo *repo_mock.IAccountRepo, operationTypeRepo *repo_mock.IOperationTypeRepo) {

// 				accountRepo.On("First",
// 					mock.Anything,
// 					mock.Anything).
// 					Return(func(searchModel interface{}, destinationModel interface{}) error {
// 						return nil
// 					}).Once()

// 				operationTypeRepo.On("First",
// 					mock.Anything,
// 					mock.Anything).
// 					Return(func(searchModel interface{}, destinationModel interface{}) error {
// 						return nil
// 					}).Once()

// 				transactionRepo.On("Create",
// 					mock.Anything,
// 					mock.Anything).
// 					Return(func(transactionModel interface{}, tableName string) error {
// 						ac := transactionModel.(*model.Transaction)
// 						ac.Id = 2
// 						transactionModel = ac
// 						return nil
// 					}).Once()
// 			},
// 			input: &dto.TransactionCreateRequest{
// 				AccountId:       123,
// 				OperationTypeId: 234,
// 				Amount:          110.1,
// 			},
// 			actualResponse: &dto.TransactionCreateResponse{
// 				TransactionId: 2,
// 			},
// 			actualErrorString: "",
// 		},
// 		{
// 			name: "failure from transaction repo",
// 			mock: func(transactionRepo *repo_mock.ITransactionRepo, accountRepo *repo_mock.IAccountRepo, operationTypeRepo *repo_mock.IOperationTypeRepo) {

// 				accountRepo.On("First",
// 					mock.Anything,
// 					mock.Anything).
// 					Return(func(searchModel interface{}, destinationModel interface{}) error {
// 						return nil
// 					}).Once()

// 				operationTypeRepo.On("First",
// 					mock.Anything,
// 					mock.Anything).
// 					Return(func(searchModel interface{}, destinationModel interface{}) error {
// 						return nil
// 					}).Once()

// 				transactionRepo.On("Create",
// 					mock.Anything,
// 					mock.Anything).
// 					Return(func(transactionModel interface{}, tableName string) error {
// 						return errors.New("non processing entity")
// 					}).Once()
// 			},
// 			input: &dto.TransactionCreateRequest{
// 				AccountId:       123,
// 				OperationTypeId: 234,
// 				Amount:          110.1,
// 			},
// 			actualResponse:    nil,
// 			actualErrorString: "non processing entity",
// 		},
// 		{
// 			name: "failure invalid operation type",
// 			mock: func(transactionRepo *repo_mock.ITransactionRepo, accountRepo *repo_mock.IAccountRepo, operationTypeRepo *repo_mock.IOperationTypeRepo) {

// 				accountRepo.On("First",
// 					mock.Anything,
// 					mock.Anything).
// 					Return(func(searchModel interface{}, destinationModel interface{}) error {
// 						return nil
// 					}).Once()

// 				operationTypeRepo.On("First",
// 					mock.Anything,
// 					mock.Anything).
// 					Return(func(searchModel interface{}, destinationModel interface{}) error {
// 						return errors.New("invalid operation type")
// 					}).Once()

// 			},
// 			input: &dto.TransactionCreateRequest{
// 				AccountId:       123,
// 				OperationTypeId: 234,
// 				Amount:          110.1,
// 			},
// 			actualResponse:    nil,
// 			actualErrorString: "no operations associated with given operation_type_id",
// 		},
// 		{
// 			name: "failure invalid account",
// 			mock: func(transactionRepo *repo_mock.ITransactionRepo, accountRepo *repo_mock.IAccountRepo, operationTypeRepo *repo_mock.IOperationTypeRepo) {

// 				accountRepo.On("First",
// 					mock.Anything,
// 					mock.Anything).
// 					Return(func(searchModel interface{}, destinationModel interface{}) error {
// 						return errors.New("invalid account")
// 					}).Once()
// 			},
// 			input: &dto.TransactionCreateRequest{
// 				AccountId:       123,
// 				OperationTypeId: 234,
// 				Amount:          110.1,
// 			},
// 			actualResponse:    nil,
// 			actualErrorString: "no account associated with given account_id",
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			transactionRepoMock := repo_mock.NewITransactionRepo(t)
// 			accountRepoMock := repo_mock.NewIAccountRepo(t)
// 			operationTypeRepoMock := repo_mock.NewIOperationTypeRepo(t)

// 			test.mock(transactionRepoMock, accountRepoMock, operationTypeRepoMock)

// 			transactionService := services.NewTransactionService(transactionRepoMock, accountRepoMock, operationTypeRepoMock)

// 			w := httptest.NewRecorder()
// 			ginCtx, _ := gin.CreateTestContext(w)
// 			ginCtx.Set("logger", log)

// 			rCtx := request_context.GetRCtx(ginCtx)
// 			resp, err := transactionService.CreateTransaction(&rCtx, test.input)
// 			if err != nil {
// 				assert.Equal(t, test.actualErrorString, err.Error())
// 			}
// 			assert.Equal(t, test.actualResponse, resp)

// 		})
// 	}

// }

*/

func TestServiceTransaction(t *testing.T) {
	log := logger.GetLogger()

	type testCase struct {
		name              string
		mock              func(transactionRepo *repo_mock.ITransactionRepo, accountRepo *repo_mock.IAccountRepo, operationTypeRepo *repo_mock.IOperationTypeRepo)
		input             *dto.TransactionCreateRequest
		actualResponse    *dto.TransactionCreateResponse
		actualErrorString string
	}

	tests := []testCase{
		{
			name: "success",
			mock: func(transactionRepo *repo_mock.ITransactionRepo, accountRepo *repo_mock.IAccountRepo, operationTypeRepo *repo_mock.IOperationTypeRepo) {

				accountRepo.On("First",
					mock.Anything,
					mock.Anything).
					Return(func(searchModel interface{}, destinationModel interface{}) error {
						return nil
					}).Once()

				operationTypeRepo.On("First",
					mock.Anything,
					mock.Anything).
					Return(func(searchModel interface{}, destinationModel interface{}) error {
						return nil
					}).Once()

				transactionRepo.On("FetchAllTransactionsWithAmountLessThan",
					mock.Anything,
					mock.Anything).
					Return(func(accountId int, amount float64) ([]model.Transaction, error) {
						return []model.Transaction{
							{AccountId: accountId, Amount: -10.2, Balance: -10.2},
							{AccountId: accountId, Amount: -30.2, Balance: -30.2},
							{AccountId: accountId, Amount: -100.2, Balance: -100.2},
						}, nil
					}).Once()

				transactionRepo.On("Update",
					mock.Anything).
					Return(func(model interface{}) error {
						return nil
					}).Once()

				transactionRepo.On("Update",
					mock.Anything).
					Return(func(model interface{}) error {
						return nil
					}).Once()

				transactionRepo.On("Create",
					mock.Anything,
					mock.Anything).
					Return(func(transactionModel interface{}, tableName string) error {
						ac := transactionModel.(*model.Transaction)

						ac.Id = 2
						transactionModel = ac
						return nil
					}).Once()
			},
			input: &dto.TransactionCreateRequest{
				AccountId:       123,
				OperationTypeId: 234,
				Amount:          110.1,
			},
			actualResponse: &dto.TransactionCreateResponse{
				TransactionId: 2,
			},
			actualErrorString: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			transactionRepoMock := repo_mock.NewITransactionRepo(t)
			accountRepoMock := repo_mock.NewIAccountRepo(t)
			operationTypeRepoMock := repo_mock.NewIOperationTypeRepo(t)

			test.mock(transactionRepoMock, accountRepoMock, operationTypeRepoMock)

			transactionService := services.NewTransactionService(transactionRepoMock, accountRepoMock, operationTypeRepoMock)

			w := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(w)
			ginCtx.Set("logger", log)

			rCtx := request_context.GetRCtx(ginCtx)
			resp, err := transactionService.CreateTransaction(&rCtx, test.input)
			if err != nil {
				assert.Equal(t, test.actualErrorString, err.Error())
			}
			assert.Equal(t, test.actualResponse, resp)

		})
	}

}
