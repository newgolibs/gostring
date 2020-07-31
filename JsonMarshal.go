package gostring

import (
	"encoding/json"
)

func JsonMarshal(v interface{}) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}
