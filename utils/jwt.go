package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strings"
)

func VerifyToken(ctx *gin.Context, SecretPublicKeyEnvName string) (jwt.MapClaims, error) {
	tokenHeader := ctx.GetHeader("Authorization")
	accessToken := strings.SplitAfter(tokenHeader, "Bearer")[1]
	//Trim
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
		return nil, err
	}

	if !token.Valid {
		logrus.Error("Invalid token")
		return nil, errors.New("Token wasn't verified correctly")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok { //&& token.Valid {
		return claims, nil
	}

	return nil, errors.New("Token wasn't verified correctly")
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
