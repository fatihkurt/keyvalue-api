package helper

import (
	"bytes"
	"deliveryhero/constants"
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	Code    string
	Message string
	Op      string
	Err     error
}

func ErrorCode(err error) string {
	if err == nil {
		return ""
	}
	var appError *AppError
	ok := errors.As(err, &appError)

	if ok && appError.Code != "" {
		return appError.Code
	} else if ok && appError.Err != nil {
		return ErrorCode(appError.Err)
	}
	return constants.EINTERNAL
}

func ErrorMessage(err error) string {
	if err == nil {
		return ""
	}
	var appError *AppError
	ok := errors.As(err, &appError)
	if ok && appError.Message != "" {
		return appError.Message
	} else if ok && appError.Err != nil {
		return ErrorMessage(appError.Err)
	}
	return "An internal error has occurred. Please contact technical support."
}

func (e *AppError) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&buf, "(%s) ", e.Code)
		}
		buf.WriteString(e.Message)
	}
	return buf.String()
}

func GetStatusCodeWithDefault(err error) (respCode int) {
	var appError *AppError
	ok := errors.As(err, &appError)
	if !ok {
		respCode = 500
	} else {
		respCode = getStatusCode(ErrorCode(err))
	}
	return
}

func getStatusCode(code string) int {
	if code == constants.ENOTFOUND {
		return http.StatusNotFound
	} else if code == constants.ECONFLICT {
		return http.StatusConflict
	} else {
		return http.StatusBadRequest
	}
}
