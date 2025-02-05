package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"wh-gin/config"
	"wh-gin/internal/controllers"
	"wh-gin/internal/repositories"
	"wh-gin/internal/usecases"
	"wh-gin/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.InitDB()

	router := gin.Default()

	if os.Getenv("RATE_LIMIT") == "true" {
		router.Use(utils.RateLimiter)
	}

	userRepo := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	controllers.NewUserHandler(router, userUsecase)

	prodRepo := repositories.NewProductRepository(db)
	prodUsecase := usecases.NewProductUsecase(prodRepo)
	controllers.NewProductHandler(router, prodUsecase)

	locRepo := repositories.NewWarehouseRepository(db)
	locUsecase := usecases.NewWarehouseUsecase(locRepo)
	controllers.NewWarehouseHandler(router, locUsecase)

	orderRepo := repositories.NewOrderRepository(db)
	orderUsecase := usecases.NewOrderUsecase(orderRepo, prodRepo, locRepo)
	controllers.NewOrderHandler(router, orderUsecase)

	go orderUsecase.ExecuteOrder(context.Background())

	router.Run(":5689")
}
