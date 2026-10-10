package response

import (
	"log"

	"github.com/gin-gonic/gin"

	"todo-backend/internal/apperror"
)

// ErrorDetail 定義前端收到的錯誤資訊
type ErrorDetail struct {
	Code    apperror.Code `json:"code"`
	Message string        `json:"message"`
}

// ErrorResponse 定義統一錯誤 JSON 格式
type ErrorResponse struct {
	Success bool        `json:"success"`
	Error   ErrorDetail `json:"error"`
}

// Error 統一處理 HTTP Error Response
func Error(c *gin.Context, err error) {
	appErr := apperror.FromError(err)

	// 真正的錯誤原因保留在後端 log
	log.Printf(
		"API error: code=%s method=%s path=%s cause=%v",
		appErr.Code,
		c.Request.Method,
		c.Request.URL.Path,
		err,
	)

	// 前端只收到公開、安全的錯誤資訊
	c.JSON(appErr.HTTPStatus, ErrorResponse{
		Success: false,
		Error: ErrorDetail{
			Code:    appErr.Code,
			Message: appErr.Message,
		},
	})
}
