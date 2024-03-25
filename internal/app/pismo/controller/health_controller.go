package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (h *HealthController) HealthStatus(ctx *gin.Context) {
	resp := map[string]string{
		"status": "working",
	}

	ctx.JSON(http.StatusOK, resp)
}
