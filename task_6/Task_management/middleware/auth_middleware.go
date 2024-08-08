package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"task_management/data"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func validateSigningMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return data.SecretKey, nil
}

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

// creting amiddleware that check the wether the user is admin or not
func AdminMidleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// uses_name, user_exist := c.Get("userame")
		role, role_exist := c.Get("Role")

		if !role_exist || role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unautorized access"})
			c.Abort()
			return
		}
		c.Next()
	}
}
