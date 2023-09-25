package utils

import "encoding/json"

func ToJson(v interface{}, indent ...bool) string {
	if len(indent) > 0 && indent[0] {
		marshalIndent, _ := json.MarshalIndent(v, "", "  ")
		return string(marshalIndent)
	}
	marshal, _ := json.Marshal(v)
	return string(marshal)
}
