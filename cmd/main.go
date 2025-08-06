package main

import (
	"NonuNamak/internal/controller"
	"NonuNamak/internal/model"
	"NonuNamak/internal/repository"
	"NonuNamak/internal/service"
	"NonuNamak/pkg/config"
	"NonuNamak/pkg/database"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("NonuNamak — backend запускается...")
	config.LoadEnv()

	database.Connect()
	db := database.DB

	db.AutoMigrate(&model.User{})

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := controller.NewUserHandler(userService)

	r := gin.Default()
	controller.RegisterUserRoutes(r, userHandler)

	r.Run(":8080")
}
