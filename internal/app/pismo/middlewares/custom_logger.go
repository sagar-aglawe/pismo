package middlewares

import (
	"self-projects/pismo/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CustomLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestId interface{}

		if requestId = ctx.Value("x-request-id"); requestId == nil || requestId == "" {
			requestId = uuid.New().String()
		}

		logFields := map[string]interface{}{
			"request-method": ctx.Request.Method,
			"request-path":   ctx.Request.URL.Path,
			"x-request-id":   requestId,
		}

		log := logger.GetLogger()
		log = log.WithFields(logFields)
		ctx.Set("logger", log)
		ctx.Next()
	}
}
