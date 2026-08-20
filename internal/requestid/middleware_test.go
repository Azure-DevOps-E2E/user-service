package requestid

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMiddlewareUsesProvidedRequestID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(Middleware())
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, From(c))
	})

	request := httptest.NewRequest(http.MethodGet, "/ping", nil)
	request.Header.Set(HeaderName, "test-request-id")
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
	if got := response.Header().Get(HeaderName); got != "test-request-id" {
		t.Fatalf("expected response header %q, got %q", "test-request-id", got)
	}
	if body := response.Body.String(); body != "test-request-id" {
		t.Fatalf("expected response body %q, got %q", "test-request-id", body)
	}
}

func TestMiddlewareGeneratesAndPropagatesRequestID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(Middleware())
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, From(c))
	})

	request := httptest.NewRequest(http.MethodGet, "/ping", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
	if got := response.Header().Get(HeaderName); got == "" {
		t.Fatal("expected generated request ID in response header")
	} else if body := response.Body.String(); body != got {
		t.Fatalf("expected response body to match generated request ID, got body %q and header %q", body, got)
	}
}

func TestFromReturnsUnknownForMissingOrInvalidValue(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	if got := From(c); got != "unknown" {
		t.Fatalf("expected unknown for missing request ID, got %q", got)
	}

	c.Set(contextKey, 123)
	if got := From(c); got != "unknown" {
		t.Fatalf("expected unknown for invalid request ID type, got %q", got)
	}
}

func TestGenerateFallsBackToTimestampWhenRandFails(t *testing.T) {
	originalRead := randRead
	t.Cleanup(func() {
		randRead = originalRead
	})

	randRead = func([]byte) (int, error) {
		return 0, errors.New("rand failure")
	}

	got := generate()
	if !regexp.MustCompile(`^\d{14}\.\d{9}$`).MatchString(got) {
		t.Fatalf("expected timestamp fallback, got %q", got)
	}
}
