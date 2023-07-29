package main

import (
	route "github.com/MAD-Goat-Project/mad-profile-service/routes"
	"github.com/MAD-Goat-Project/mad-profile-service/utils" // Update this import path
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	router := SetupRouter()

	log.Fatal(router.Run(":" + utils.GodotEnv("GO_PORT")))
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	/**
	@description Setup Mode Application
	*/
	if utils.GodotEnv("GO_ENV") != "production" && utils.GodotEnv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if utils.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// TODO: Allow origins * is not safe, change this later
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	route.InitProfileRoutes(router)

	return router
}
