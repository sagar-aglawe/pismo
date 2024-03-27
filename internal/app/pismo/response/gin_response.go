package response

import "github.com/gin-gonic/gin"

type Response gin.H

func FormatResponse(ctx *gin.Context, success bool, statusCode int, data interface{}, error_message string) {

	res := &Response{
		"success":       success,
		"data":          data,
		"error_message": error_message,
	}

	ctx.JSON(statusCode, res)
}
