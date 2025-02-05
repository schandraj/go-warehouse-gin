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

type mockProductUsecase struct{}

func (m mockProductUsecase) InsertProduct(ctx context.Context, product entities.Product) error {
	return nil
}

func (m mockProductUsecase) GetAllProducts(ctx context.Context, limit, offset int) ([]entities.Product, int, error) {
	return []entities.Product{}, 0, nil
}

func (m mockProductUsecase) GetDetail(ctx context.Context, id int) (entities.Product, error) {
	return entities.Product{}, nil
}

func (m mockProductUsecase) Update(ctx context.Context, product entities.Product) error {
	return nil
}

func (m mockProductUsecase) Delete(ctx context.Context, id int) error {
	return nil
}

func TestInsertProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	mockUC := new(mockProductUsecase)

	productHandler := &ProductHandler{mockUC}
	r.POST("/products", productHandler.InsertProduct)

	product := &entities.Product{
		Name:     "test",
		SKU:      "lp-asd-22",
		Quantity: 2,
	}

	reqJson, _ := json.Marshal(product)

	req, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(reqJson))
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetAllProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	mockUC := new(mockProductUsecase)

	productHandler := &ProductHandler{mockUC}
	r.GET("/products", productHandler.GetAllProducts)

	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	mockUC := new(mockProductUsecase)

	productHandler := &ProductHandler{mockUC}
	r.GET("/products/:id", productHandler.GetByID)

	req, _ := http.NewRequest(http.MethodGet, "/products/1", nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestUpdateByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	mockUC := new(mockProductUsecase)

	productHandler := &ProductHandler{mockUC}
	r.PUT("/products/:id", productHandler.UpdateByID)

	product := entities.Product{
		Name:     "test",
		SKU:      "lp-asd-22",
		Quantity: 5,
	}

	reqJson, _ := json.Marshal(product)

	req, _ := http.NewRequest(http.MethodPut, "/products/1", bytes.NewBuffer(reqJson))
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestDeleteByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	mockUC := new(mockProductUsecase)

	productHandler := &ProductHandler{mockUC}
	r.DELETE("/products/:id", productHandler.UpdateByID)

	req, _ := http.NewRequest(http.MethodDelete, "/products/1", nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
