package user

import (
	"encoding/json"

	"github.com/kye-gregory/koicards-api/pkg/util"
)

type ID struct {
	value string
}

func NewID() *ID {
	return &ID{value: util.GenerateRandomString(24, 24, "abcdefghijklmnopqrstuvwxyz0123456789")}
}

func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.value)
}

func (id ID) String() string {
	return id.value
}