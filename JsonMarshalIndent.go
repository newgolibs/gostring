package gostring

import (
	"encoding/json"
	"fmt"
	"runtime"
)

func JsonMarshalIndent(v interface{}) string {
	bytes, err := json.MarshalIndent(v, "\t", "")
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		panic(fmt.Sprintf("\033[41;36merr:%+v %+v:%+v\033[0m\n", []interface{}{err}, file, line))
	}
	return string(bytes)
}
