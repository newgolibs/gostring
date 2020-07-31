package gostring

import (
	"encoding/json"
)

func JsonMarshalIndent(v interface{}) string {
	bytes, _ := json.MarshalIndent(v, "\t", "")
	return string(bytes)
}
