package middleware

import (
	"car-rental/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helper.TokenValid(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "You are unauthorized to access this resource, login to your account with your email and password",
			})
			return
		}
		c.Next()
	}
}
