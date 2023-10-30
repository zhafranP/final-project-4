package handler

import (
	"finalProject4/infrastructure/config"
	"finalProject4/infrastructure/database"
	middlewares "finalProject4/pkg/middleware"
	"finalProject4/repository/user_repository/user_pg"
	"finalProject4/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	config.LoadAppConfig()
	database.InitiliazeDatabase()
	db := database.GetDatabaseInstance()
	database.SeedAdmin(db)

	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	r := gin.Default()
	users := r.Group("/users")
	{
		users.POST("/register", userHandler.CreateUser)
		users.POST("/login", userHandler.Login)
		users.Use(middlewares.Authentication())
		{
			users.PATCH("/topup", userHandler.Topup)
		}
	}

	r.Run(":" + config.GetAppConfig().Port)
}
