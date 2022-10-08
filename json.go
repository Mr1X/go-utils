package goutils

import (
	"encoding/json"
)

// JSONIndent converts the golang value to indent JSON string
// such as a struct, map, slice, array and so on
func JSONIndent(obj interface{}) string {
	buf, err := json.MarshalIndent(obj, "", "	")
	if err != nil {
		return err.Error()
	}
	return string(buf)
}

// JSONString converts the golang value to JSON string
// such as a struct, map, slice, array and so on
func JSONString(obj interface{}) string {
	buf, err := json.Marshal(obj)
	if err != nil {
		return err.Error()
	}
	return string(buf)
}
