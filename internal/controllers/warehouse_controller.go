package controllers

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
	"wh-gin/internal/entities"
	"wh-gin/internal/usecases"
	"wh-gin/utils"
)

type WarehouseHandler struct {
	uc usecases.WarehouseUsecase
}

func NewWarehouseHandler(r *gin.Engine, uc usecases.WarehouseUsecase) {
	warehouseHandler := WarehouseHandler{uc}

	locGroup := r.Group("/locations")
	locGroup.Use(utils.AuthMiddleware())
	locGroup.POST("", utils.AdminMiddleware(), warehouseHandler.CreateWL)
	locGroup.GET("", warehouseHandler.GetAll)
}

func (w *WarehouseHandler) CreateWL(c *gin.Context) {
	var input entities.WarehouseLocation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if w.uc.CreateWarehouse(c, input) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "CreateWarehouse failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (w *WarehouseHandler) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		size = 10
	}

	page, size, offset := utils.DecodePage(page, size)

	locations, total, err := w.uc.GetAll(c, size, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": locations,
		"meta": gin.H{
			"total":        total,
			"current_page": page,
			"total_page":   int(math.Ceil(float64(total) / float64(size))),
		},
	})
}
