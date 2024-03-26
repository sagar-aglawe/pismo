package response

import "github.com/gin-gonic/gin"

type Response gin.H

func FormatResponse(ctx *gin.Context, success bool, statusCode int, data interface{}) {

	res := &Response{
		"success":       true,
		"data":          data,
		"error_message": "",
	}

	ctx.JSON(statusCode, res)
}
