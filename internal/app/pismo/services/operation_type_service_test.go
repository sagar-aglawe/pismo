package services_test

import (
	"errors"
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

func TestServiceCreateOperationType(t *testing.T) {
	log := logger.GetLogger()

	type testCase struct {
		name              string
		mock              func(mockService *repo_mock.IOperationTypeRepo)
		input             *dto.OperationTypeCreateRequest
		actualResponse    *dto.OperationTypeCreateResponse
		actualErrorString string
	}

	tests := []testCase{
		{
			name: "success",
			mock: func(mockService *repo_mock.IOperationTypeRepo) {
				mockService.On("Create",
					mock.Anything,
					mock.Anything).
					Return(func(operationTypeModel interface{}, tableName string) error {
						ac := operationTypeModel.(*model.OperationType)
						ac.Id = 2
						operationTypeModel = ac
						return nil
					}).Once()
			},
			input: &dto.OperationTypeCreateRequest{
				Description: "Normal Purchase",
			},
			actualResponse:    &dto.OperationTypeCreateResponse{OperationTypeId: 2},
			actualErrorString: "",
		},
		{
			name: "failure from repository",
			mock: func(mockService *repo_mock.IOperationTypeRepo) {
				mockService.On("Create",
					mock.Anything,
					mock.Anything).
					Return(func(accountModel interface{}, tableName string) error {
						return errors.New("non processing request")
					}).Once()
			},
			input: &dto.OperationTypeCreateRequest{
				Description: "Normal Purchase",
			},
			actualResponse:    nil,
			actualErrorString: "non processing request",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			operationTypeRepoMock := repo_mock.NewIOperationTypeRepo(t)
			test.mock(operationTypeRepoMock)

			operationTypeService := services.NewOperationTypeService(operationTypeRepoMock)

			w := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(w)
			ginCtx.Set("logger", log)

			rCtx := request_context.GetRCtx(ginCtx)
			resp, err := operationTypeService.CreateOperationType(&rCtx, test.input)
			if err != nil {
				assert.Equal(t, test.actualErrorString, err.Error())
			}
			assert.Equal(t, test.actualResponse, resp)

		})
	}

}
