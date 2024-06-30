package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func BasicAuthMiddleware(userName string, password string) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			c.Header("WWW-Authenticate", "Basic realm=Restricted")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		authHeader = strings.TrimPrefix(authHeader, "Basic ")
		decodedAuth, err := decodeBase64(authHeader)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// split the decoded credentials into username and password
		uname, pwd, ok := strings.Cut(decodedAuth, ":")

		// check if the username and password match the expected values
		if !ok || uname != userName || pwd != password {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// if the credentials are valid, allow the request to proceed
		c.Next()
	}
}

// decodeBase64 decodes a base64 string
func decodeBase64(input string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(decodedBytes), nil
}
