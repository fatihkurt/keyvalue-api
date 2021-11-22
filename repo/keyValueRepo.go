package repo

import (
	"context"
	"deliveryhero/constants"
	"deliveryhero/helper"

	"github.com/go-redis/redis/v8"
)

type KeyValueRepo struct {
	Context context.Context
	Client  *redis.Client
}

func NewKeyValueRepo(client *redis.Client) *KeyValueRepo {
	return &KeyValueRepo{
		Context: context.Background(),
		Client:  client,
	}
}

func (d *KeyValueRepo) GetKey(key string) (value string, err error) {
	op := "KeyValueRepo.GetKey"

	val, err := d.Client.Get(d.Context, key).Result()

	if err == redis.Nil {
		return "", &helper.AppError{Op: op, Code: constants.ENOTFOUND}
	}

	if err != nil {
		return "", err
	}
	return val, nil
}

func (d *KeyValueRepo) SetKey(key string, value string) (err error) {
	op := "KeyValueRepo.SetKey"
	err = d.Client.Set(d.Context, key, value, 0).Err()
	if err != nil {
		return &helper.AppError{Op: op, Err: err}
	}

	return
}

func (d *KeyValueRepo) FlushDb() (err error) {
	op := "KeyValueRepo.FlushDb"
	err = d.Client.FlushDB(d.Context).Err()
	if err != nil {
		return &helper.AppError{Op: op, Err: err}
	}
	return
}
