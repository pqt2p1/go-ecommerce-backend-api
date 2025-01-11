package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pqt2p1/go-ecommerce-backend-api/pkg/response"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorReponse(c, response.ErrInvalidToken)
			c.Abort()
		}
		c.Next()
	}
}