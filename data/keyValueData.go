package data

type KeyValueData struct{}

func NewKeyValueData() *KeyValueData {
	return &KeyValueData{}
}

func (d *KeyValueData) GetKey(key string) (value string, err error) {
	return
}

func (d *KeyValueData) SetKey(key string, value string) (err error) {
	return
}
