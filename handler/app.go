package handler

import (
	"finalProject4/infrastructure/config"
	"finalProject4/infrastructure/database"
	middlewares "finalProject4/pkg/middleware"
	"finalProject4/repository/category_repository/category_pg"
	"finalProject4/repository/product_repository/product_pg"
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

	categoryRepo := category_pg.NewCategoryPG(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := NewCategoryHandler(categoryService)

	productRepo := product_pg.NewProductPG(db)
	productService := service.NewProductService(productRepo)
	productHandler := NewProductHandler(productService)

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

	products := r.Group("/products")
	{
		products.Use(middlewares.Authentication())
		{
			products.POST("", middlewares.AdminAuthorization(), productHandler.CreateProduct)
			products.PUT("/:productId", middlewares.AdminAuthorization(), productHandler.UpdateProduct)
			products.GET("", productHandler.GetProduct)
			products.DELETE("/:productId", middlewares.AdminAuthorization(), productHandler.DeleteProduct)
		}
	}

	r.Run(":" + config.GetAppConfig().Port)
}
