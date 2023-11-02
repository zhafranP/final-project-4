package handler

import (
	"finalProject4/infrastructure/config"
	"finalProject4/infrastructure/database"
	middlewares "finalProject4/pkg/middleware"
	"finalProject4/service"
	"finalProject4/repository/user_repository/user_pg"
	"finalProject4/repository/category_repository/category_pg"

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

	categoryRepo := category_pg.NewCategoryPG(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := NewCategoryHandler(categoryService)

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

	categories := r.Group("/categories")
	{
		categories.Use(middlewares.Authentication())
		{
			categories.Use(middlewares.AdminAuthorization())
			categories.GET("/", categoryHandler.GetCategories)
			categories.POST("/", categoryHandler.CreateCategory)
		}
	}

	r.Run(":" + config.GetAppConfig().Port)
}
