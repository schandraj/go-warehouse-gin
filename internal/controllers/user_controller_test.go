package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"wh-gin/internal/entities"
)

type mockUserUsecase struct {
}

func (m mockUserUsecase) Register(ctx context.Context, input entities.User) error {
	return nil
}

func (m mockUserUsecase) Login(ctx context.Context, username, password string) (string, error) {
	return "ini token", nil
}

func (m mockUserUsecase) Me(ctx context.Context, username string) (entities.User, error) {
	return entities.User{Username: "test", ID: 1, Role: "admin"}, nil
}

func (m mockUserUsecase) GetAll(ctx context.Context, limit int, offset int) ([]entities.User, int, error) {
	return nil, 0, nil
}

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	mockUC := new(mockUserUsecase)

	userHandler := &UserHandler{mockUC}
	r.POST("/register", userHandler.Register)

	request := entities.UserInput{
		Username:        "test",
		Password:        "test",
		ConfirmPassword: "test",
		Role:            "admin",
	}

	reqJson, _ := json.Marshal(request)

	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(reqJson))
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockUC := new(mockUserUsecase)
	userHandler := &UserHandler{mockUC}
	r.POST("/login", userHandler.Login)
	request := json.RawMessage(`{"username":"test","password":"test"}`)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(request))
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestMe(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUC := new(mockUserUsecase)
	userHandler := &UserHandler{mockUC}
	req, _ := http.NewRequest(http.MethodGet, "/users/me", nil)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = req
	c.Set("claims", jwt.MapClaims{
		"username": "test",
		"role":     "admin",
	})
	userHandler.GetMe(c)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetAllUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockUC := new(mockUserUsecase)
	userHandler := &UserHandler{mockUC}
	r.GET("/users", userHandler.GetAllUsers)
	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
}
