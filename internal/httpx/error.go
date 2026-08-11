package httpx

import (
	"github.com/gin-gonic/gin"

	"nexuscart/user-service/internal/requestid"
)

type ErrorBody struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}

type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}

func WriteError(c *gin.Context, status int, code, message string) {
	c.JSON(status, ErrorResponse{Error: ErrorBody{
		Code:      code,
		Message:   message,
		RequestID: requestid.From(c),
	}})
}
