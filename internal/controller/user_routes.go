package controller

import (
	"NonuNamak/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, h *UserHandler) {
	users := r.Group("/users")

	users.POST("/login", h.Login)
	users.POST("/", h.CreateUser)

	authUsers := users.Group("/")
	authUsers.Use(middleware.AuthMiddleware())
	{
		authUsers.GET("/", h.GetAllUsers)
		authUsers.GET("/:id", h.GetUserByID)
		authUsers.GET("/me", h.GetMe)

	}

	adminUsers := users.Group("/")
	adminUsers.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminUsers.PUT("/:id", h.UpdateUser)
		adminUsers.DELETE("/:id", h.DeleteUser)
		adminUsers.PATCH("/:id", h.PatchUser)
	}

}
