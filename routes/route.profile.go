package routes

import (
	"github.com/MAD-Goat-Project/mad-profile-service/controllers"
	"github.com/MAD-Goat-Project/mad-profile-service/middlewares"
	"github.com/gin-gonic/gin"
)

func InitProfileRoutes(route *gin.Engine) {

	groupRoute := route.Group("/api/v1").Use(middlewares.Auth())
	groupRoute.GET("/", welcome)
	groupRoute.GET("profile", controllers.GetUserDetails)

}


func welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to MAD Profile Service",
	})
}