package service

import "deliveryhero/data"

type KeyValueService struct {
	data *data.KeyValueData
}

func NewKeyValueService() *KeyValueService {
	keyValueData := data.NewKeyValueData()
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
