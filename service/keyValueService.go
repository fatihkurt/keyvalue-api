package service

import (
	"deliveryhero/db"
	"deliveryhero/helper"
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
	op := "Service.GetKey"
	val, err := s.repo.GetKey(key)
	if err != nil {
		return "", &helper.AppError{Op: op, Err: err}
	}
	return val, nil
}

func (s *KeyValueService) SetKey(key string, value string) (err error) {
	op := "Service.SetKey"

	err = s.repo.SetKey(key, value)
	if err != nil {
		return &helper.AppError{Op: op, Err: err}
	}

	return
}
