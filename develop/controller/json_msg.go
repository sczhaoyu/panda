package controller

import (
	"encoding/json"
)

func ToJson(retOrError interface{}) []byte {
	ret := struct {
		Success bool        `json:"success"`
		Ret     interface{} `json:"ret,omitempty"`
		Message string      `json:"msg,omitempty"`
	}{
		Success: true,
	}

	switch err := retOrError.(type) {
	case error:
		ret.Message = err.Error()
		ret.Success = false
	default:
		ret.Ret = retOrError
	}

	b, _ := json.Marshal(ret)
	return b
}
