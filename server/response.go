package server

import (
	"deliveryhero/constants"
	"deliveryhero/helper"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Result interface{} `json:"result"`
	Error  *string     `json:"error"`
}

func OkResponse(c *gin.Context, result interface{}) {
	response := Response{Error: nil, Result: result}
	c.JSON(http.StatusOK, response)
}

func ErrorResponse(c *gin.Context, err error) {
	errMsg := err.Error()

	c.JSON(getStatusCodeWithDefault(err), Response{Error: &errMsg, Result: nil})
}

func getStatusCodeWithDefault(err error) (respCode int) {
	var appError *helper.AppError
	ok := errors.As(err, &appError)
	if !ok {
		respCode = 500
	} else {
		respCode = getStatusCode(helper.ErrorCode(err))
	}
	return
}

func getStatusCode(code string) int {
	if code == constants.ENOTFOUND {
		return http.StatusNotFound
	} else {
		return http.StatusBadRequest
	}
}
