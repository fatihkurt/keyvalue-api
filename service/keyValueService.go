package service

import (
	"deliveryhero/data"

	"github.com/go-redis/redis/v8"
)

type KeyValueService struct {
	data *data.KeyValueData
}

func NewKeyValueService(client *redis.Client) *KeyValueService {
	keyValueData := data.NewKeyValueData(client)
	return &KeyValueService{
		data: keyValueData,
	}
}

func (s *KeyValueService) GetKey(key string) (value string, err error) {
	// op := "Service.GetKey"
	// TODO data logic
	return
}

func (s *KeyValueService) SetKey(key string, value string) (err error) {
	// op := "Service.SetKey"

	// TODO data logic

	return
}
