package infrastructure

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "auth header"})
			c.Abort()
			return
		}
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unautorized access"})
			c.Abort()
			return
		}
		tokenClaim := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(authParts[1], tokenClaim, validateSigningMethod)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Set("Username", tokenClaim["username"])
		c.Set("Role", tokenClaim["role"])
		c.Next()
	}
}
