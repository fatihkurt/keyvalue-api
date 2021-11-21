package data

import "github.com/go-redis/redis/v8"

type KeyValueData struct {
	Client *redis.Client
}

func NewKeyValueData(client *redis.Client) *KeyValueData {
	return &KeyValueData{
		Client: client,
	}
}

func (d *KeyValueData) GetKey(key string) (value string, err error) {
	return
}

func (d *KeyValueData) SetKey(key string, value string) (err error) {
	return
}
