package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PageResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Total   int64       `json:"total"`
	Page    int         `json:"page"`
	Size    int         `json:"size"`
}

// Success
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

// Error
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}

// BadRequest
func BadRequest(c *gin.Context, message string) {
	Error(c, http.StatusBadRequest, message)
}

// Unauthorized
func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, message)
}

// Forbidden
func Forbidden(c *gin.Context, message string) {
	Error(c, http.StatusForbidden, message)
}

// NotFound
func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message)
}

// InternalServerError
func InternalServerError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, message)
}

// SuccessWithPagination
func SuccessWithPagination(c *gin.Context, data interface{}, total int64, page, size int) {
	c.JSON(http.StatusOK, PageResponse{
		Code:    200,
		Message: "success",
		Total:   total,
		Page:    page,
		Size:    size,
		Data:    data,
	})
}
// ErrorResponse 错误响应（带详细信息）
func ErrorResponse(c *gin.Context, statusCode int, message string, details interface{}) {
	response := Response{
		Code:    statusCode,
		Message: message,
	}

	if details != nil {
		response.Data = gin.H{"details": details}
	}

	c.JSON(statusCode, response)
}

// SuccessResponse 成功响应（带消息和数据）
func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: message,
		Data:    data,
	})
}
