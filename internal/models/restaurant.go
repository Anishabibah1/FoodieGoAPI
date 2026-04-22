package controllers

import (
	"net/http"

	"foodiego-api/internal/config"
	"foodiego-api/internal/models"

	"github.com/gin-gonic/gin"
)

func GetAllRestaurants(c *gin.Context) {
	var restaurants []models.Restaurant
	config.DB.Find(&restaurants)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   restaurants,
	})
}

func GetRestaurantByID(c *gin.Context) {
	var restaurant models.Restaurant
	if err := config.DB.First(&restaurant, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Restaurant not found.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   restaurant,
	})
}

func SearchRestaurants(c *gin.Context) {
	var restaurants []models.Restaurant
	q := c.Query("q")
	if q == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Query parameter q is required.",
		})
		return
	}
	config.DB.Where("name LIKE ? OR category LIKE ?", "%"+q+"%", "%"+q+"%").Find(&restaurants)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   restaurants,
	})
}