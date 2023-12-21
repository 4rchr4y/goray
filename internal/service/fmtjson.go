package fmtjson

import (
	"encoding/json"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
)

func Raw(b []byte) (string, error) {
	return string(b), nil
}

func Pretty(b []byte) (string, error) {
	var obj map[string]interface{}

	err := json.Unmarshal(b, &obj)
	if err != nil {
		return "", err
	}

	f := colorjson.NewFormatter()
	f.Indent = 2
	f.KeyColor = color.New(color.FgHiBlue)

	s, err := f.Marshal(obj)

	return string(s), err
}
