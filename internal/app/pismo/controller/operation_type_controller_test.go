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

func TestCreateOperationType(t *testing.T) {
	log := logger.GetLogger()

	type testCase struct {
		name           string
		endpoint       string
		method         string
		mock           func(mockService *service_mocks.IOperationTypeService)
		input          map[string]interface{}
		actualResponse ControllerResponse
	}

	tests := []testCase{
		{
			name:     "success",
			endpoint: "/pismo/api/v1/operations",
			method:   http.MethodPost,
			mock: func(mockService *service_mocks.IOperationTypeService) {
				mockService.On("CreateOperationType",
					mock.Anything,
					mock.Anything).
					Return(func(rCtx *request_context.ReqCtx, reqBody *dto.OperationTypeCreateRequest) (*dto.OperationTypeCreateResponse, error) {
						return &dto.OperationTypeCreateResponse{OperationTypeId: 1}, nil
					}).Once()
			},
			input: map[string]interface{}{
				"description": "Normal Purchase",
			},
			actualResponse: ControllerResponse{
				Data: map[string]interface{}{
					"operation_type_id": float64(1),
				},
				Success:      true,
				ErrorMessage: "",
			},
		},
		{
			name:     "failure from service",
			endpoint: "/pismo/api/v1/operations",
			method:   http.MethodPost,
			mock: func(mockService *service_mocks.IOperationTypeService) {
				mockService.On("CreateOperationType",
					mock.Anything,
					mock.Anything).
					Return(func(rCtx *request_context.ReqCtx, reqBody *dto.OperationTypeCreateRequest) (*dto.OperationTypeCreateResponse, error) {
						return nil, errors.New("unprocessible entity")
					}).Once()
			},
			input: map[string]interface{}{
				"description": "Normal Purchase",
			},
			actualResponse: ControllerResponse{
				Data:         "",
				Success:      false,
				ErrorMessage: "unprocessible entity",
			},
		},
		{
			name:     "validation failure",
			endpoint: "/pismo/api/v1/operations",
			method:   http.MethodPost,
			mock:     func(mockService *service_mocks.IOperationTypeService) {},
			input:    map[string]interface{}{},
			actualResponse: ControllerResponse{
				Data:         "",
				Success:      false,
				ErrorMessage: "Key: 'OperationTypeCreateRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			operationTypeServiceMock := service_mocks.NewIOperationTypeService(t)
			test.mock(operationTypeServiceMock)

			operationTypeController := controller.NewOperationTypeController(operationTypeServiceMock)

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

			operationTypeController.CreateOperationType(ginCtx)

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
