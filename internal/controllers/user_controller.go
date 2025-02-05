package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"math"
	"net/http"
	"strconv"
	"wh-gin/internal/entities"
	"wh-gin/internal/usecases"
	"wh-gin/utils"
)

type UserHandler struct {
	uc usecases.UserUsecase
}

func NewUserHandler(r *gin.Engine, usecase usecases.UserUsecase) {
	handler := &UserHandler{
		uc: usecase,
	}

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	userGroup := r.Group("/users")
	userGroup.Use(utils.AuthMiddleware())
	userGroup.GET("/me", handler.GetMe)
	userGroup.GET("", utils.AdminMiddleware(), handler.GetAllUsers)
}

func (h *UserHandler) Register(c *gin.Context) {
	var input entities.UserInput
	var data entities.User
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Password != input.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password does not match"})
		return
	}

	if input.Role != "admin" && input.Role != "staff" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role"})
	}

	hashPass, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	data.Password = hashPass
	data.Role = input.Role
	data.Username = input.Username

	if err = h.uc.Register(c, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "User registered"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	token, err := h.uc.Login(c, request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *UserHandler) GetMe(c *gin.Context) {
	claims := c.MustGet("claims").(jwt.MapClaims)
	username := claims["username"].(string)

	me, err := h.uc.Me(c, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": me.Username,
		"role":     me.Role,
		"id":       me.ID,
	})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		size = 10
	}

	page, size, offset := utils.DecodePage(page, size)

	all, total, err := h.uc.GetAll(c, size, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": all,
		"meta": gin.H{
			"total":        total,
			"current_page": page,
			"total_page":   int(math.Ceil(float64(total) / float64(size))),
		},
	})
}
