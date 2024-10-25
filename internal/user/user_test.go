package user

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"initial_project_go/pkg/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserService adalah struct yang digunakan untuk mock UserService
type MockUserService struct {
	mock.Mock
}

// Mock method untuk ServiceAddUsers
func (m *MockUserService) ServiceAddUsers(ctx context.Context, request *RequestAddUsers) (any, int, error) {
	args := m.Called(ctx, request)
	return args.Get(0), args.Int(1), args.Error(2)
}

// Mock method untuk ServiceGetUsers
func (m *MockUserService) ServiceGetUsers(ctx context.Context, request *RequestGetUsers) (any, int, error) {
	args := m.Called(ctx, request)
	return args.Get(0), args.Int(1), args.Error(2)
}

// Setup function to initialize routes with mocks
func setupTestRouter(mockUserService *MockUserService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	userController := NewUserController(mockUserService)
	v1 := r.Group("/v1")
	{
		v1.POST("/add/users", userController.AddUsers)
		v1.POST("/get/users", userController.GetUsers)
	}
	return r
}

func TestAddUsersEndpoint(t *testing.T) {
	mockUserService := new(MockUserService)
	router := setupTestRouter(mockUserService)
	url := "/v1/add/users"

	t.Run("onSuccess", func(t *testing.T) {
		// Sample request data
		requestAddUsers := RequestAddUsers{
			Name: "Test User",
		}
		// Convert request to JSON
		requestBody, _ := json.Marshal(requestAddUsers)

		// Mocking the service response
		mockUserService.On("ServiceAddUsers", mock.Anything, &requestAddUsers).Return("Add Users Success", 200, nil)

		// Create a new HTTP POST request
		req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		// Record the response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert the response status code and body
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Add Users Success")

		mockUserService.AssertExpectations(t)
	})

	t.Run("onEmptyNameRequest", func(t *testing.T) {
		// Sample request data
		requestAddUsers := RequestAddUsers{
			Name: "",
		}
		// Convert request to JSON
		requestBody, _ := json.Marshal(requestAddUsers)

		// Mocking the service response
		mockUserService.On("ServiceAddUsers", mock.Anything, mock.Anything).Return(utils.Response{}, 400, errors.New(errParams))

		// Create a new HTTP POST request
		req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		// Record the response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert the response status code and body
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), `"name":"The Name field is required."`)

		mockUserService.AssertExpectations(t)
	})
}

func TestGetUsersEndpoint(t *testing.T) {
	mockUserService := new(MockUserService)
	router := setupTestRouter(mockUserService)
	url := "/v1/get/users"

	t.Run("onSuccessWithKeyword", func(t *testing.T) {
		// Sample request data
		requestGetUsers := RequestGetUsers{
			Keyword: "Test",
			Limit:   nil,
			Offset:  nil,
		}
		// Convert request to JSON
		requestBody, _ := json.Marshal(requestGetUsers)

		// Sample response data
		mockUserResponse := []Users{
			{
				Id:   "123",
				Name: "Test User",
			},
		}
		// Mocking the service response
		mockUserService.On("ServiceGetUsers", mock.Anything, &requestGetUsers).Return(mockUserResponse, 200, nil)

		// Create a new HTTP POST request
		req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		// Record the response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert the response status code and body
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Test User")

		mockUserService.AssertExpectations(t)
	})

	t.Run("onSuccessNoKeyValue", func(t *testing.T) {
		// Sample request data
		requestGetUsers := RequestGetUsers{}

		// Convert request to JSON
		requestBody, _ := json.Marshal(requestGetUsers)

		// Sample response data
		mockUserResponse := []Users{
			{
				Id:   "123",
				Name: "Test User",
			},
		}
		// Mocking the service response
		mockUserService.On("ServiceGetUsers", mock.Anything, &requestGetUsers).Return(mockUserResponse, 200, nil)

		// Create a new HTTP POST request
		req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		// Record the response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert the response status code and body
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Test User")

		mockUserService.AssertExpectations(t)
	})
}
