package keycloak

import (
	"fmt"

	"github.com/MAD-Goat-Project/mad-profile-service/utils"
	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
)

type Keycloak struct {
	Token *gocloak.JWT
}

func GetUserDetails(ctx *gin.Context, email string, roles map[string]interface{}) ([]map[string]interface{}, error) {
	keycloakURL := utils.GodotEnv("KC_SERVER_URL")
	realm := utils.GodotEnv("KC_REALM")
	clientID := utils.GodotEnv("KC_CLIENT_ID")
	clientSecret := utils.GodotEnv("KC_SECRET")

	keycloakClient := gocloak.NewClient(keycloakURL)
	contextValue := ctx.Request.Context()

	token, err := keycloakClient.LoginClient(contextValue, clientID, clientSecret, realm)
	if err != nil {
		return nil, fmt.Errorf("failed to obtain admin token: %s", err)
	}

	// Check if "app-user" and "app-admin" roles exist in realm_access
	hasAppUser := false
	hasAppAdmin := false

	if roleList, ok := roles["roles"].([]interface{}); ok {
		for _, role := range roleList {
			if roleName, ok := role.(string); ok {
				if roleName == "app-user" {
					hasAppUser = true
				} else if roleName == "app-admin" {
					hasAppAdmin = true
				}
			}
		}
	}

	if !hasAppUser && !hasAppAdmin {
		return nil, fmt.Errorf("user does not have required roles")
	}

	userParams := gocloak.GetUsersParams{}

	if hasAppAdmin {
		userParams = gocloak.GetUsersParams{}
	}

	if hasAppUser {
		userParams = gocloak.GetUsersParams{
			Email: &email,
		}
	}

	users, err := keycloakClient.GetUsers(contextValue, token.AccessToken, realm, userParams)

	var usersWithAttributes []map[string]interface{}
	for _, user := range users {
		userWithAttributes := map[string]interface{}{
			"email":     *user.Email,
			"username":  *user.Username,
			"firstname": *user.FirstName,
			"lastname":  *user.LastName,
		}
		usersWithAttributes = append(usersWithAttributes, userWithAttributes)
	}

	return usersWithAttributes, nil
}
