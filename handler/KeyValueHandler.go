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

func (h *KeyValueHandler) HanldeGetKey() (err error) {
	op := "GetKey"
	key := h.Context.Query("key")

	if key == "" {
		return &helper.AppError{Code: constants.EINVALID, Op: op}
	}

	result := &model.KeyValue{Key: "dump", Value: "test"}
	helper.OkResponse(h.Context, result)
	return
}

func (h *KeyValueHandler) HanldeSetKey() (err error) {
	op := "SetKey"
	key := h.Context.Query("key")
	// value := h.Context.Query("value")

	if key == "" {
		return &helper.AppError{Code: constants.EINVALID, Op: op}
	}

	helper.OkResponse(h.Context, true)
	return
}
