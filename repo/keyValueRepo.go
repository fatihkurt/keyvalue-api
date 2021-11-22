package repo

import (
	"context"
	"deliveryhero/constants"
	"deliveryhero/helper"

	"github.com/go-redis/redis/v8"
)

type KeyValueData struct {
	Context context.Context
	Client  *redis.Client
}

func NewKeyValueRepo(client *redis.Client) *KeyValueData {
	return &KeyValueData{
		Context: context.Background(),
		Client:  client,
	}
}

func (d *KeyValueData) GetKey(key string) (value string, err error) {

	val, err := d.Client.Get(d.Context, key).Result()

	if err == redis.Nil {
		return "", &helper.AppError{Code: constants.ENOTFOUND}
	}

	if err != nil {
		return "", err
	}
	return val, nil
}

func (d *KeyValueData) SetKey(key string, value string) (err error) {
	err = d.Client.Set(d.Context, key, value, 0).Err()
	if err != nil {
		return err
	}

	return
}
