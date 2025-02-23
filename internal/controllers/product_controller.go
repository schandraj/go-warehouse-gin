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

type ProductHandler struct {
	uc usecases.ProductUsecase
}

func NewProductHandler(r *gin.Engine, uc usecases.ProductUsecase) {
	productHandler := &ProductHandler{
		uc: uc,
	}

	prodGroup := r.Group("/products")
	prodGroup.Use(utils.AuthMiddleware())
	prodGroup.POST("", utils.AdminMiddleware(), productHandler.InsertProduct)
	prodGroup.GET("", productHandler.GetAllProducts)
	prodGroup.GET("/:id", productHandler.GetByID)
	prodGroup.PUT("/:id", utils.AdminMiddleware(), productHandler.UpdateByID)
	prodGroup.DELETE("/:id", utils.AdminMiddleware(), productHandler.DeleteByID)

}

func (p *ProductHandler) InsertProduct(c *gin.Context) {
	var input entities.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := p.uc.InsertProduct(c, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (p *ProductHandler) GetAllProducts(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		size = 10
	}

	page, size, offset := utils.DecodePage(page, size)

	products, total, err := p.uc.GetAllProducts(c, size, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
		"meta": gin.H{
			"total":        total,
			"current_page": page,
			"total_page":   int(math.Ceil(float64(total) / float64(size))),
		},
	})
}

func (p *ProductHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	detail, err := p.uc.GetDetail(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": detail})
}

func (p *ProductHandler) UpdateByID(c *gin.Context) {
	var product entities.Product
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = uint(id)

	err = p.uc.Update(c, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (p *ProductHandler) DeleteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = p.uc.Delete(c, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
