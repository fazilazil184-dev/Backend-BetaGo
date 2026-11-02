package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, statusCode int, message string, data any) {
	ctx.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func SuccessWithMeta(ctx *gin.Context, statusCode int, message string, data any, meta *Meta) {
	ctx.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func Created(ctx *gin.Context, message string, data any) {
	Success(ctx, http.StatusCreated, message, data)
}

func Error(ctx *gin.Context, statusCode int, message string, code string, details []ErrorDetail) {
	ctx.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Error: &ErrorResponse{
			Code:    code,
			Details: details,
		},
	})
}

func NoContent(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func BadRequest(ctx *gin.Context, message string) {
	Error(ctx, http.StatusBadRequest, message, CodeBadRequest, nil)
}

func ValidationError(ctx *gin.Context, details []ErrorDetail) {
	Error(ctx, http.StatusUnprocessableEntity, "Validation failed", CodeValidation, details)
}

func Unauthorized(ctx *gin.Context, message string) {
	Error(ctx, http.StatusUnauthorized, message, CodeUnauthorized, nil)
}

func Forbidden(ctx *gin.Context, message string) {
	Error(ctx, http.StatusForbidden, message, CodeForbidden, nil)
}

func NotFound(ctx *gin.Context, message string) {
	Error(ctx, http.StatusNotFound, message, CodeNotFound, nil)
}

func Conflict(ctx *gin.Context, message string) {
	Error(ctx, http.StatusConflict, message, CodeConflict, nil)
}

func InternalServerError(ctx *gin.Context) {
	Error(ctx, http.StatusInternalServerError, "Internal server error", CodeInternalServerError, nil)
}
