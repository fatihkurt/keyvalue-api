package handler

import (
	"github.com/gin-gonic/gin"

	"deliveryhero/constants"
	"deliveryhero/helper"
	"deliveryhero/model"
	"deliveryhero/service"
)

type KeyValueHandler struct {
	Context *gin.Context
	Service *service.KeyValueService
}

func NewKeyValueHandler(ctx *gin.Context) *KeyValueHandler {
	keyValueService := service.NewKeyValueService()
	return &KeyValueHandler{
		Context: ctx,
		Service: keyValueService,
	}
}

func (h *KeyValueHandler) HanldeGetKey() (err error) {
	op := "HandleGetKey"
	key := h.Context.Query("key")

	if key == "" {
		return &helper.AppError{Code: constants.EINVALID, Op: op}
	}

	value, err := h.Service.GetKey(key)

	if err != nil {
		return &helper.AppError{Code: constants.EINTERNAL, Op: op, Err: err}
	}

	result := &model.KeyValue{Key: "dump", Value: value}
	helper.OkResponse(h.Context, result)
	return
}

func (h *KeyValueHandler) HanldeSetKey() (err error) {
	op := "HandleSetKey"
	key := h.Context.Query("key")
	value := h.Context.Query("value")

	if key == "" {
		return &helper.AppError{Code: constants.EINVALID, Op: op}
	}

	err = h.Service.SetKey(key, value)

	if err != nil {
		return err
	}

	helper.OkResponse(h.Context, true)
	return
}
