package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

func Error(
	ctx *gin.Context,
	statusCode int,
	message string,
	errors interface{},
) {
	ctx.JSON(statusCode, ErrorResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

func BadRequest(ctx *gin.Context, detail string) {
	Error(ctx, http.StatusBadRequest, "bad request", detail)
}

func Unauthorized(ctx *gin.Context, detail string) {
	Error(ctx, http.StatusUnauthorized, "unauthorized", detail)
}

func InternalServerError(ctx *gin.Context, detail string) {
	Error(ctx, http.StatusInternalServerError, "internal server error", detail)
}
