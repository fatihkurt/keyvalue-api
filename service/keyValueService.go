package service

import (
	"deliveryhero/db"
	"deliveryhero/repo"
)

type KeyValueService struct {
	repo *repo.KeyValueRepo
}

func NewKeyValueService() *KeyValueService {
	client := db.RedisClient()
	return &KeyValueService{
		repo: repo.NewKeyValueRepo(client),
	}
}

func (s *KeyValueService) GetKey(key string) (string, error) {
	return s.repo.GetKey(key)
}

func (s *KeyValueService) SetKey(key string, value string) error {
	return s.repo.SetKey(key, value)
}

func (s *KeyValueService) FlushDb() error {
	return s.repo.FlushDb()
}
