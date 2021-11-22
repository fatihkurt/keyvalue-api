package handler

import (
	"deliveryhero/constants"
	"deliveryhero/helper"
	"deliveryhero/model"
	"deliveryhero/service"
)

type KeyValueHttpHandler struct {
	Service *service.KeyValueService
}

func NewKeyValueHttpHandler() *KeyValueHttpHandler {
	keyValueService := service.NewKeyValueService()
	return &KeyValueHttpHandler{
		Service: keyValueService,
	}
}

func (h *KeyValueHttpHandler) HanldeGetKey(key string) (result *model.KeyValue, err error) {
	op := "HandleGetKey"

	if key == "" {
		return nil, &helper.AppError{Code: constants.EINVALID, Op: op, Message: "No key selected"}
	}

	value, err := h.Service.GetKey(key)

	if err != nil {
		return nil, err
	}

	result = &model.KeyValue{Key: key, Value: value}
	return
}

func (h *KeyValueHttpHandler) HanldeSetKey(data model.KeyValue) (*bool, error) {
	op := "HandleSetKey"

	if data.Key == "" {
		return nil, &helper.AppError{Code: constants.EINVALID, Op: op}
	}

	err := h.Service.SetKey(data.Key, data.Value)

	if err != nil {
		return nil, err
	}

	res := true

	return &res, nil
}

func (h *KeyValueHttpHandler) HandleFlushDb() (*bool, error) {
	err := h.Service.FlushDb()

	if err != nil {
		return nil, err
	}

	res := true

	return &res, nil
}
