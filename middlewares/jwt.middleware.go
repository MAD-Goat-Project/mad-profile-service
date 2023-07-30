package middlewares

import (
	"github.com/MAD-Goat-Project/mad-profile-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UnauthorizedError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Method  string `json:"method"`
	Message string `json:"message"`
}

func Auth() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var errorResponse UnauthorizedError

		errorResponse.Status = "Forbidden"
		errorResponse.Code = http.StatusForbidden
		errorResponse.Method = ctx.Request.Method
		errorResponse.Message = "Authorization is required for this endpoint"

		if ctx.GetHeader("Authorization") == "" {
			ctx.JSON(http.StatusForbidden, errorResponse)
			defer ctx.AbortWithStatus(http.StatusForbidden)
		}

		claims, err := utils.VerifyToken(ctx, "JWT_SECRET")

		errorResponse.Status = "Unauthorized"
		errorResponse.Code = http.StatusUnauthorized
		errorResponse.Method = ctx.Request.Method
		errorResponse.Message = "accessToken invalid or expired"

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, errorResponse)
			defer ctx.AbortWithStatus(http.StatusUnauthorized)
		} else {
			// global value result
			ctx.Set("user", claims)
			// return to next method if token is valid
			ctx.Next()
		}
	}
}
