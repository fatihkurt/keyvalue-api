package repo

import (
	"context"

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

	// val := d.Client.Get(d.Context, key)

	return
}

func (d *KeyValueData) SetKey(key string, value string) (err error) {
	return
}
