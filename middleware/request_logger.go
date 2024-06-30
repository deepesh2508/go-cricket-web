package middleware

import (
	"bytes"
	"io"
	"os"
	"time"

	"github.com/deepesh2508/go-cricket-web/env"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic("failed to initialize zap logger: " + err.Error())
	}
}

func RequestLogger() gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	process := env.ENV.PROCESS_NAME
	nodeip := env.ENV.SERVER_IP

	return func(c *gin.Context) {
		start := time.Now()

		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		c.Request.Body = io.NopCloser(&buf)
		body, _ := io.ReadAll(tee)
		requestBody := string(body)

		c.Next()

		stop := time.Since(start)
		uuid := c.GetString("uuid")
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		clientReferer := c.Request.Referer()
		requestPath := c.Request.URL.Path
		statusCode := c.Writer.Status()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		logger.Info("Request",
			zap.String("uuid", uuid),
			zap.String("clientIP", clientIP),
			zap.Int("dataLength", dataLength),
			zap.String("hostname", hostname),
			zap.String("serverip", nodeip),
			zap.String("latency", stop.String()),
			zap.String("request", requestBody),
			zap.String("method", c.Request.Method),
			zap.String("path", requestPath),
			zap.String("process", process),
			zap.String("referer", clientReferer),
			zap.Int("statusCode", statusCode),
			zap.String("timestamp", start.Format(time.RFC3339Nano)),
			zap.String("userAgent", clientUserAgent),
		)
	}
}
