package controllers

import (
	keycloak "github.com/MAD-Goat-Project/mad-profile-service/services"
	"github.com/MAD-Goat-Project/mad-profile-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func GetUserDetails(c *gin.Context) {

	claims, exists := c.Get("user")

	realmAccessRoles, err := utils.GetJWTRealmAccessRoles(claims.(jwt.MapClaims))
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	email, err := utils.GetJWTAccountEmail(claims.(jwt.MapClaims))
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	if !exists {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	details, err := keycloak.GetUserDetails(c, email, realmAccessRoles)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, details)
}
