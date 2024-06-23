package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSingleResponseCreated(ctx *gin.Context, data any, descriptionMsg string) {
	ctx.JSON(http.StatusOK, &SingleResponse{
		Status: Status{
			Code:        http.StatusCreated,
			Description: descriptionMsg,
		},
		Data: data,
	})
}

func SendSingleResponse(ctx *gin.Context, data any, descriptionMsg string) {
	ctx.JSON(http.StatusOK, &SingleResponse{
		Status: Status{
			Code:        http.StatusOK,
			Description: descriptionMsg,
		},
		Data: data,
	})
}

func SendSingleResponseError(ctx *gin.Context, code int, errorMessage string) {
	ctx.AbortWithStatusJSON(code, &Status{
		Code:        code,
		Description: errorMessage,
	})
}
