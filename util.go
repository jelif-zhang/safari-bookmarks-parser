package main

import (
	"encoding/json"
)

func ToJson(data interface{}) []byte {
	b, _ := json.Marshal(data)
	return b
}