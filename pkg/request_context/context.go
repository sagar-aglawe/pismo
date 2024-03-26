package request_context

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ReqCtx struct {
	Log *logrus.Entry
}

func GetRCtx(ctx *gin.Context) ReqCtx {
	log := ctx.Value("logger").(*logrus.Entry)

	return ReqCtx{
		Log: log,
	}
}
