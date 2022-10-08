package goutils

import (
	"bytes"
	"encoding/json"
)

// JSONIndent converts the golang value to indent JSON string
// such as a struct, map, slice, array and so on
func JSONIndent(obj interface{}) string {
	bs, err := json.Marshal(obj)
	if err != nil {
		return err.Error()
	}
	buf := new(bytes.Buffer)
	err = json.Indent(buf, bs, "", "\t")
	if err != nil {
		return err.Error()
	}
	return buf.String()
}

// JSONString converts the golang value to JSON string
// such as a struct, map, slice, array and so on
func JSONString(obj interface{}) string {
	bs, err := json.Marshal(obj)
	if err != nil {
		return err.Error()
	}
	return string(bs)
}
