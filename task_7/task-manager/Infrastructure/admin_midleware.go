package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
