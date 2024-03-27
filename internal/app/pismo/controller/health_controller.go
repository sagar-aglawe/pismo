package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

type IHealthController interface {
	HealthStatus(ctx *gin.Context)
}

func NewHealthController() IHealthController {
	return &HealthController{}
}

// HealthStatus 	    godoc
// @Summary 			Get Health
// @Description 		Get the health of the application
// @Produce 			application/json
// @Tags 				Health
// @Success 			200 {object} map[string]interface{}
// @Router 				/health [get]
func (h *HealthController) HealthStatus(ctx *gin.Context) {
	resp := map[string]string{
		"status": "working",
	}

	ctx.JSON(http.StatusOK, resp)
}
