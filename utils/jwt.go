package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

func VerifyToken(ctx *gin.Context, SecretPublicKeyEnvName string) (jwt.MapClaims, error) {
	tokenHeader := ctx.GetHeader("Authorization")
	accessToken := strings.SplitAfter(tokenHeader, "Bearer")[1]
	accessToken = strings.Trim(accessToken, " ")
	jwtSecretKey := GodotEnv(SecretPublicKeyEnvName)

	publicKey, err := parseKeycloakRSAPublicKey(jwtSecretKey)
	if err != nil {
		panic(err)
	}

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// return the public key that is used to validate the token.
		return publicKey, nil
	})
	
	if err != nil {
		logrus.Error("Error parsing or validating token:", err)
	}


	if !token.Valid {
		logrus.Error("Invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok { //&& token.Valid {
		return claims, nil
	}

	return nil, errors.New("token wasn't verified correctly")
}

func parseKeycloakRSAPublicKey(base64Encoded string) (*rsa.PublicKey, error) {
	buf, err := base64.StdEncoding.DecodeString(base64Encoded)
	if err != nil {
		return nil, err
	}
	parsedKey, err := x509.ParsePKIXPublicKey(buf)
	if err != nil {
		return nil, err
	}
	publicKey, ok := parsedKey.(*rsa.PublicKey)
	if ok {
		return publicKey, nil
	}
	return nil, fmt.Errorf("unexpected key type %T", publicKey)
}

func GetJWTRealmAccessRoles(claims jwt.MapClaims) (map[string]interface{}, error) {
	realmAccess, ok := claims["realm_access"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("realm access data not found or invalid")
	}

	return realmAccess, nil
}

func GetJWTAccountEmail(claims jwt.MapClaims) (string, error) {

	email, ok := claims["email"].(string)
	if !ok {
		return "", fmt.Errorf("email not found or invalid")
	}

	return email, nil
}
