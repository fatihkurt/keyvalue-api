package service

type KeyValueService struct {
	db interface{}
}

func NewKeyValueService() *KeyValueService {
	return &KeyValueService{
		db: "dump",
	}
}

func (h *KeyValueService) GetKey(key string) (value string, err error) {
	// op := "Service.GetKey"
	// TODO data logic
	return
}

func (h *KeyValueService) SetKey(key string, value string) (err error) {
	// op := "Service.SetKey"

	// TODO data logic

	return
}
