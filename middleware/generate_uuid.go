package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GenerateUUID() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.NewString()
		c.Set("uuid", uuid)
	}
}
