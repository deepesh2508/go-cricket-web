package util

import (
	"net/http"

	"time"

	"github.com/deepesh2508/go-cricket-web/helper/errs"
	logg "github.com/deepesh2508/go-cricket-web/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SendResponse(c *gin.Context, sucessRes interface{}, errRes *errs.ErrorResponse, err error) {
	if sucessRes == nil {
		res := GraphQLError{
			Message: errRes.GetErrorMessage(c),
			Details: ErrorDetails{
				ErrorCode: errRes.C,
			},
		}
		if err != nil {
			res.Details.Error = err.Error()
		}
		logg.Error(c, errRes.M, zap.String("uuid", c.GetString("uuid")), zap.String("errorcode", errRes.C), zap.Error(err))
		c.JSON(errRes.H, res)
		c.AbortWithStatus(errRes.H)
		return
	}
	logg.Info(c, "Success", zap.String("uuid", c.GetString("uuid")), zap.Any("response", sucessRes))
	c.JSON(http.StatusOK, sucessRes)
}

type GraphQLError struct {
	Message string       `json:"message"`
	Details ErrorDetails `json:"extensions"`
}

type ErrorDetails struct {
	ErrorCode string `json:"code"`
	Error     string `json:"internalerror"`
}

func NewHttpClient() *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	httpClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}

	return httpClient
}
