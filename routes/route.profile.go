package routes

import (
	"github.com/MAD-Goat-Project/mad-profile-service/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitProfileRoutes(route *gin.Engine) {

	groupRoute := route.Group("/api/v1/profile").Use(middlewares.Auth())
	groupRoute.GET("/", getUserDetails)

}

func getUserDetails(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ""+
		"{"+
		"\"id\": \"1\","+
		"\"name\": \"John Coltrane\","+
		"\"email\": \""+
		"}")
}
