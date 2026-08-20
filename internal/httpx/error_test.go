package httpx

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"nexuscart/user-service/internal/requestid"
)

func TestWriteErrorIncludesRequestID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(requestid.Middleware())
	router.GET("/boom", func(c *gin.Context) {
		WriteError(c, http.StatusTeapot, "TEAPOT", "short and stout")
	})

	request := httptest.NewRequest(http.MethodGet, "/boom", nil)
	request.Header.Set(requestid.HeaderName, "httpx-test")
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusTeapot {
		t.Fatalf("expected status 418, got %d", response.Code)
	}
	if got := response.Header().Get(requestid.HeaderName); got != "httpx-test" {
		t.Fatalf("expected request ID header %q, got %q", "httpx-test", got)
	}

	var body ErrorResponse
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if body.Error.Code != "TEAPOT" || body.Error.Message != "short and stout" || body.Error.RequestID != "httpx-test" {
		t.Fatalf("unexpected error response: %+v", body.Error)
	}
}