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

type mockWarehouseUsecase struct {
}

func (m mockWarehouseUsecase) CreateWarehouse(ctx context.Context, warehouse entities.WarehouseLocation) error {
	return nil
}

func (m mockWarehouseUsecase) GetAll(ctx context.Context, limit, offset int) ([]entities.WarehouseLocation, int, error) {
	return nil, 0, nil
}

func TestCreateWL(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	mockUC := new(mockWarehouseUsecase)
	whHandler := &WarehouseHandler{mockUC}
	r.POST("/locations", whHandler.CreateWL)

	request := entities.WarehouseLocation{
		Name:     "CGK",
		Capacity: 1000,
	}

	reqJson, _ := json.Marshal(request)

	req, _ := http.NewRequest(http.MethodPost, "/locations", bytes.NewReader(reqJson))
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockUC := new(mockWarehouseUsecase)
	whHandler := &WarehouseHandler{mockUC}
	r.GET("/locations", whHandler.GetAll)
	req, _ := http.NewRequest(http.MethodGet, "/locations", nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
}
