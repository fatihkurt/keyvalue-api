package service

import (
	"deliveryhero/db"
	"deliveryhero/repo"
)

type KeyValueService struct {
	repo *repo.KeyValueData
}

func NewKeyValueService() *KeyValueService {
	client := db.RedisClient()
	return &KeyValueService{
		repo: repo.NewKeyValueRepo(client),
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
