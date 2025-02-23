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

type OrderHandler struct {
	uc usecases.OrderUsecase
}

func NewOrderHandler(r *gin.Engine, uc usecases.OrderUsecase) {
	orderHandler := &OrderHandler{uc}

	orderGroup := r.Group("/orders").Use(utils.AuthMiddleware())
	orderGroup.POST("/:type", utils.StaffMiddleware(), orderHandler.CreateNewOrder)
	orderGroup.GET("", orderHandler.GetAll)
	orderGroup.GET("/:id", orderHandler.GetByID)
}

func (o *OrderHandler) CreateNewOrder(c *gin.Context) {
	var input entities.Order
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.Param("type") != "receive" && c.Param("type") != "ship" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type must be 'receive' or 'ship'"})
		return
	}

	input.Status = "pending"
	input.Type = c.Param("type")

	if err := o.uc.CreateOrder(c, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (o *OrderHandler) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		size = 10
	}

	page, size, offset := utils.DecodePage(page, size)

	data, total, err := o.uc.GetAllOrders(c, size, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
		"meta": gin.H{
			"total":        total,
			"current_page": page,
			"total_page":   int(math.Ceil(float64(total) / float64(size))),
		},
	})
}

func (o *OrderHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := o.uc.GetByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": data})
}
