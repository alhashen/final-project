package middleware

import (
	"finalproj/helper"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helper.ValidateToken(c)

		if err != nil {
			helper.Unauthorized(c)
			return
		}

		c.Set("token", verifyToken)
		c.Next()
	}
}
