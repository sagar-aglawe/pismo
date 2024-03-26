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

func (h *HealthController) HealthStatus(ctx *gin.Context) {
	resp := map[string]string{
		"status": "working",
	}

	ctx.JSON(http.StatusOK, resp)
}
