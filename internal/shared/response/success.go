package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(
	ctx *gin.Context,
	statusCode int,
	message string,
	data interface{},
) {
	ctx.JSON(statusCode, SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Created(
	ctx *gin.Context,
	message string,
	data interface{},
) {
	Success(ctx, http.StatusCreated, message, data)
}

func OK(
	ctx *gin.Context,
	message string,
	data interface{},
) {
	Success(ctx, http.StatusOK, message, data)
}
