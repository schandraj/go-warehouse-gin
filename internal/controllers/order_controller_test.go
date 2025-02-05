package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"wh-gin/internal/entities"
)

type mockOrderUsecase struct{}

func (m *mockOrderUsecase) GetByID(ctx context.Context, id int) (entities.Order, error) {
	return entities.Order{
		ID:     1,
		Status: "completed",
	}, nil
}

func (m *mockOrderUsecase) ExecuteOrder(ctx context.Context) {

}

func (m *mockOrderUsecase) CreateOrder(c context.Context, order entities.Order) error {
	return nil
}

func (m *mockOrderUsecase) GetAllOrders(c context.Context, limit, offset int) ([]entities.Order, int, error) {
	return []entities.Order{{ID: 1, Status: "pending"}}, 1, nil
}

func (m *mockOrderUsecase) GetOrderById(c *gin.Context, id int) (*entities.Order, error) {
	return &entities.Order{ID: uint(id), Status: "pending"}, nil
}

func TestCreateNewOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	mockUC := new(mockOrderUsecase)

	orderHandler := &OrderHandler{uc: mockUC}
	r.POST("/orders/:type", orderHandler.CreateNewOrder)

	order := entities.Order{ID: 1, Status: "pending"}
	orderJSON, _ := json.Marshal(order)

	req, _ := http.NewRequest(http.MethodPost, "/orders/receive", bytes.NewBuffer(orderJSON))
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetAllOrders(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockUC := &mockOrderUsecase{}

	orderHandler := &OrderHandler{uc: mockUC}
	r.GET("/orders", orderHandler.GetAll) // No AuthMiddleware

	req, _ := http.NewRequest(http.MethodGet, "/orders", nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetOrderByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockUC := &mockOrderUsecase{}
	orderHandler := &OrderHandler{uc: mockUC}

	r.GET("/orders/:id", orderHandler.GetByID)

	req, _ := http.NewRequest(http.MethodGet, "/orders/1", nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
