package controller_test

import (
	"bytes"
	"encoding/json"
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

type ControllerResponse struct {
	Data         map[string]interface{} `json:"data"`
	ErrorMessage string                 `json:"error_message"`
	Success      bool                   `json:"success"`
}

func TestCreateAccount(t *testing.T) {
	log := logger.GetLogger()

	type testCase struct {
		name           string
		endpoint       string
		method         string
		mock           func(mockService *service_mocks.IAccountService)
		input          map[string]interface{}
		actualResponse ControllerResponse
	}

	tests := []testCase{
		{
			name:     "success",
			endpoint: "/pismo/api/v1/accounts",
			method:   http.MethodPost,
			mock: func(mockService *service_mocks.IAccountService) {
				mockService.On("CreateAccount",
					mock.Anything,
					mock.Anything).
					Return(func(rCtx *request_context.ReqCtx, reqBody *dto.AccountCreateRequest) (*dto.AccountCreateResponse, error) {
						return &dto.AccountCreateResponse{AccountId: 1}, nil
					}).Once()
			},
			input: map[string]interface{}{
				"document_number": "1234",
			},
			actualResponse: ControllerResponse{
				Data: map[string]interface{}{
					"account_id": float64(1),
				},
				Success:      true,
				ErrorMessage: "",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			accountServiceMock := service_mocks.NewIAccountService(t)
			test.mock(accountServiceMock)

			accountController := controller.NewAccountController(accountServiceMock)

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

			accountController.CreateAccount(ginCtx)

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
