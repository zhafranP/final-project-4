package middlewares

import (
	"finalProject4/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenClaims, err := helpers.VerifyToken(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "user is unauthorized",
				"error":   err.Error(),
			})
			return
		}
		c.Set("user", tokenClaims)
		c.Next()
	}
}
