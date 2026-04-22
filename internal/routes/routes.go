package routes

import (
	"foodiego-api/internal/controllers"
	"foodiego-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	restaurants := api.Group("/restaurants")
	restaurants.Use(middleware.AuthMiddleware())
	{
		restaurants.GET("/", controllers.GetAllRestaurants)
		restaurants.GET("/search", controllers.SearchRestaurants)
		restaurants.GET("/:id", controllers.GetRestaurantByID)
	}
}