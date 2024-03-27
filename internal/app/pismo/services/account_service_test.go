package services_test

import (
	"errors"
	"net/http/httptest"
	"self-projects/pismo/internal/app/pismo/dto"
	"self-projects/pismo/internal/app/pismo/model"
	"self-projects/pismo/internal/app/pismo/services"
	"self-projects/pismo/pkg/logger"
	"self-projects/pismo/pkg/request_context"
	"testing"

	repo_mock "self-projects/pismo/internal/app/pismo/repository/mocks"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func TestServiceCreateAccount(t *testing.T) {
	log := logger.GetLogger()

	type testCase struct {
		name              string
		mock              func(mockService *repo_mock.IAccountRepo)
		input             *dto.AccountCreateRequest
		actualResponse    *dto.AccountCreateResponse
		actualErrorString string
	}

	tests := []testCase{
		{
			name: "success",
			mock: func(mockService *repo_mock.IAccountRepo) {
				mockService.On("Create",
					mock.Anything,
					mock.Anything).
					Return(func(accountModel interface{}, tableName string) error {
						ac := accountModel.(*model.Account)
						ac.Id = 2
						accountModel = ac
						return nil
					}).Once()
			},
			input: &dto.AccountCreateRequest{
				DocumentNumber: "1234",
			},
			actualResponse:    &dto.AccountCreateResponse{AccountId: 2},
			actualErrorString: "",
		},
		{
			name: "failure from repository",
			mock: func(mockService *repo_mock.IAccountRepo) {
				mockService.On("Create",
					mock.Anything,
					mock.Anything).
					Return(func(accountModel interface{}, tableName string) error {
						return errors.New("non processing request")
					}).Once()
			},
			input: &dto.AccountCreateRequest{
				DocumentNumber: "1234",
			},
			actualResponse:    nil,
			actualErrorString: "non processing request",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			accountRepoMock := repo_mock.NewIAccountRepo(t)
			test.mock(accountRepoMock)

			accountService := services.NewAccountService(accountRepoMock)

			w := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(w)
			ginCtx.Set("logger", log)

			rCtx := request_context.GetRCtx(ginCtx)
			resp, err := accountService.CreateAccount(&rCtx, test.input)
			if err != nil {
				assert.Equal(t, test.actualErrorString, err.Error())
			}
			assert.Equal(t, test.actualResponse, resp)

		})
	}

}

func TestServiceGetAccount(t *testing.T) {
	log := logger.GetLogger()

	type testCase struct {
		name              string
		mock              func(mockService *repo_mock.IAccountRepo)
		input             int
		actualResponse    *dto.AccountGetResponse
		actualErrorString string
	}

	tests := []testCase{
		{
			name: "success",
			mock: func(mockService *repo_mock.IAccountRepo) {
				mockService.On("First",
					mock.Anything,
					mock.Anything).
					Return(func(searchModel interface{}, destinationModel interface{}) error {
						ac := destinationModel.(*model.Account)
						ac.Id = 2
						ac.DocumentNumber = "123456"
						destinationModel = ac
						return nil
					}).Once()
			},
			input:             2,
			actualResponse:    &dto.AccountGetResponse{AccountNumber: 2, DocumentNumber: "123456"},
			actualErrorString: "",
		},
		{
			name: "failure from repository",
			mock: func(mockService *repo_mock.IAccountRepo) {
				mockService.On("First",
					mock.Anything,
					mock.Anything).
					Return(func(searchModel interface{}, destinationModel interface{}) error {
						return errors.New("non processing request")
					}).Once()
			},
			input:             2,
			actualResponse:    nil,
			actualErrorString: "non processing request",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			accountRepoMock := repo_mock.NewIAccountRepo(t)
			test.mock(accountRepoMock)

			accountService := services.NewAccountService(accountRepoMock)

			w := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(w)
			ginCtx.Set("logger", log)

			rCtx := request_context.GetRCtx(ginCtx)
			resp, err := accountService.GetAccount(&rCtx, test.input)
			if err != nil {
				assert.Equal(t, test.actualErrorString, err.Error())
			}
			assert.Equal(t, test.actualResponse, resp)

		})
	}

}
