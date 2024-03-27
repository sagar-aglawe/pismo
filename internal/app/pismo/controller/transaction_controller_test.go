package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"self-projects/pismo/internal/app/pismo/controller"
	"self-projects/pismo/internal/app/pismo/dto"
	service_mocks "self-projects/pismo/internal/app/pismo/services/mocks"
	"self-projects/pismo/pkg/logger"
	"self-projects/pismo/pkg/request_context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransaction(t *testing.T) {
	log := logger.GetLogger()

	type testCase struct {
		name           string
		endpoint       string
		method         string
		mock           func(mockService *service_mocks.ITransactionService)
		input          map[string]interface{}
		actualResponse ControllerResponse
	}

	tests := []testCase{
		{
			name:     "success",
			endpoint: "/pismo/api/v1/transactions",
			method:   http.MethodPost,
			mock: func(mockService *service_mocks.ITransactionService) {
				mockService.On("CreateTransaction",
					mock.Anything,
					mock.Anything).
					Return(func(rCtx *request_context.ReqCtx, reqBody *dto.TransactionCreateRequest) (*dto.TransactionCreateResponse, error) {
						return &dto.TransactionCreateResponse{TransactionId: 1}, nil
					}).Once()
			},
			input: map[string]interface{}{
				"account_id":        1,
				"operation_type_id": 1,
				"amount":            110.1,
			},
			actualResponse: ControllerResponse{
				Data: map[string]interface{}{
					"transaction_id": float64(1),
				},
				Success:      true,
				ErrorMessage: "",
			},
		},
		{
			name:     "failure from service",
			endpoint: "/pismo/api/v1/transactions",
			method:   http.MethodPost,
			mock: func(mockService *service_mocks.ITransactionService) {
				mockService.On("CreateTransaction",
					mock.Anything,
					mock.Anything).
					Return(func(rCtx *request_context.ReqCtx, reqBody *dto.TransactionCreateRequest) (*dto.TransactionCreateResponse, error) {
						return nil, errors.New("unprocessible entity")
					}).Once()
			},
			input: map[string]interface{}{
				"account_id":        1,
				"operation_type_id": 1,
				"amount":            110.1,
			},
			actualResponse: ControllerResponse{
				Data:         "",
				Success:      false,
				ErrorMessage: "unprocessible entity",
			},
		},
		{
			name:     "validation failure",
			endpoint: "/pismo/api/v1/transactions",
			method:   http.MethodPost,
			mock:     func(mockService *service_mocks.ITransactionService) {},
			input: map[string]interface{}{
				"account_id":        1,
				"operation_type_id": 1,
			},
			actualResponse: ControllerResponse{
				Data:         "",
				Success:      false,
				ErrorMessage: "Key: 'TransactionCreateRequest.Amount' Error:Field validation for 'Amount' failed on the 'required' tag",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			transactionServiceMock := service_mocks.NewITransactionService(t)
			test.mock(transactionServiceMock)

			transactionController := controller.NewTransactionController(transactionServiceMock)

			reqBody, mErr := json.Marshal(test.input)
			if mErr != nil {
				log.WithError(mErr).Error("test-error: error while marshaling request body")
				t.Fail()
			}

			req, reqErr := http.NewRequest(test.method, test.endpoint, bytes.NewBuffer(reqBody))
			if reqErr != nil {
				log.WithError(reqErr).Error("test-error: error while creating request")
				t.Fail()
			}

			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			ginCtx, _ := gin.CreateTestContext(w)
			ginCtx.Set("logger", log)
			ginCtx.Request = req

			transactionController.CreateTransaction(ginCtx)

			response, resErr := io.ReadAll(w.Result().Body)
			if resErr != nil {
				log.WithError(resErr).Error("test-error: error while reading response")
				t.Fail()
			}

			res := ControllerResponse{}
			unErr := json.Unmarshal(response, &res)
			if unErr != nil {
				log.WithError(unErr).Error("test-error: error while unmarshaling response")
				t.Fail()
			}

			assert.Equal(t, test.actualResponse, res)

		})
	}
}
