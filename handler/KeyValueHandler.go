package handler

import (
	"github.com/gin-gonic/gin"

	"deliveryhero/constants"
	"deliveryhero/helper"
	"deliveryhero/model"
)

type KeyValueHandler struct {
	Context *gin.Context
}

func NewKeyValueHandler(ctx *gin.Context) *KeyValueHandler {
	return &KeyValueHandler{
		Context: ctx,
	}
}

func (h *KeyValueHandler) HanldeGetKey() (*model.KeyValue, error) {
	op := "GetKey"
	key := h.Context.Query("key")

	if key == "" {
		return nil, &helper.AppError{Code: constants.EINVALID, Op: op}
	}

	return &model.KeyValue{Key: "dump", Value: "test"}, nil
}

func (h *KeyValueHandler) HanldeSetKey() (*bool, error) {
	op := "SetKey"
	key := h.Context.Query("key")
	// value := h.Context.Query("value")

	if key == "" {
		return nil, &helper.AppError{Code: constants.EINVALID, Op: op}
	}

	res := true
	return &res, nil
}
