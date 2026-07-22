package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response[T interface{}] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
	Error   string `json:"error"`
}

type PaginatedResponse[T interface{}] struct {
	Response[T]
	Meta PaginationMeta `json:"meta"`
}

type PaginationMeta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

func SuccessResponse[T interface{}](c *gin.Context, message string, data T) {
	c.JSON(http.StatusOK, Response[T]{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func CreatedResponse[T interface{}](c *gin.Context, message string, data T) {
	c.JSON(http.StatusCreated, Response[T]{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	response := Response[string]{
		Success: false,
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
	}

	c.JSON(statusCode, response)
}

func BadRequestResponse(c *gin.Context, message string, err error) {
	ErrorResponse(c, http.StatusBadRequest, message, err)
}

func UnauthorizedResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, message, nil)
}

func ForbiddenResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusForbidden, message, nil)
}

func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message, nil)
}

func InternalServerErrorResponse(c *gin.Context, message string, err error) {
	ErrorResponse(c, http.StatusInternalServerError, message, err)
}

func PaginatedSuccessResponse[T interface{}](c *gin.Context, message string, data T, meta PaginationMeta) {
	c.JSON(http.StatusOK, PaginatedResponse[T]{
		Response: Response[T]{
			Success: true,
			Message: message,
			Data:    data,
		},
		Meta: meta,
	})
}
