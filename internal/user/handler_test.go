package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewHandlerStoresRepository(t *testing.T) {
	repository := NewRepository()
	handler := NewHandler(repository)

	if handler.repository != repository {
		t.Fatal("expected handler to retain repository reference")
	}
}

func TestListReturnsSeedUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(NewRepository())

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)

	handler.List(context)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", recorder.Code)
	}

	var body ListResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(body.Items) != 3 {
		t.Fatalf("expected 3 users, got %d", len(body.Items))
	}
}

func TestGetReturnsUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(NewRepository())

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/v1/users/usr-002", nil)
	context.Params = gin.Params{{Key: "id", Value: "usr-002"}}

	handler.Get(context)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", recorder.Code)
	}

	var body User
	if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if body.ID != "usr-002" || body.Name != "Tran Minh Chau" {
		t.Fatalf("unexpected user response: %+v", body)
	}
}

func TestGetMissingUserReturnsCommonErrorShape(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(NewRepository())

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/v1/users/usr-missing", nil)
	context.Params = gin.Params{{Key: "id", Value: "usr-missing"}}

	handler.Get(context)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", recorder.Code)
	}

	var body struct {
		Error struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if body.Error.Code != "USER_NOT_FOUND" {
		t.Fatalf("unexpected error code: %+v", body.Error)
	}
}