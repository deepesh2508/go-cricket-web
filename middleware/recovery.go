package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GinPanicRecovery() gin.RecoveryFunc {
	return func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			log.Fatal(c, "panic", zap.Error(errors.New(err)))
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
