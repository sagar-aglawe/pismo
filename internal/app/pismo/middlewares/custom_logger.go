package middlewares

import (
	"self-projects/pismo/internal/app/pismo/constants"
	"self-projects/pismo/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CustomLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestId interface{}

		if requestId = ctx.Value(constants.LoggerRequestId); requestId == nil || requestId == "" {
			requestId = uuid.New().String()
		}

		logFields := map[string]interface{}{
			constants.LoggerRequestMethod: ctx.Request.Method,
			constants.LoggerRequestPath:   ctx.Request.URL.Path,
			constants.LoggerRequestId:     requestId,
		}

		log := logger.GetLogger()
		log = log.WithFields(logFields)
		ctx.Set("logger", log)
		ctx.Next()
	}
}
