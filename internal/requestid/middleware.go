package requestid

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	HeaderName = "X-Request-ID"
	contextKey = "requestId"
)

var randRead = rand.Read

// Middleware accepts a request ID from the gateway or creates one for direct calls.
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(HeaderName)
		if requestID == "" {
			requestID = generate()
		}

		c.Set(contextKey, requestID)
		c.Header(HeaderName, requestID)
		c.Next()
	}
}

func From(c *gin.Context) string {
	if requestID, ok := c.Get(contextKey); ok {
		if value, valid := requestID.(string); valid {
			return value
		}
	}
	return "unknown"
}

func generate() string {
	buffer := make([]byte, 8)
	if _, err := randRead(buffer); err == nil {
		return hex.EncodeToString(buffer)
	}
	return time.Now().UTC().Format("20060102150405.000000000")
}
