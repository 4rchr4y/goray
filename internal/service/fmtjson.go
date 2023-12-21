package fmtjson

import (
	"encoding/json"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
)

type FormatterConfig struct {
	KeyColor        *color.Color
	StringColor     *color.Color
	BoolColor       *color.Color
	NumberColor     *color.Color
	NullColor       *color.Color
	StringMaxLength int
	Indent          int
	DisabledColor   bool
	RawStrings      bool
}

func JsonFormatter() *FormatterConfig {
	return &FormatterConfig{
		KeyColor:        color.New(color.FgHiBlue),
		StringColor:     color.New(color.FgGreen),
		BoolColor:       color.New(color.FgYellow),
		NumberColor:     color.New(color.FgCyan),
		NullColor:       color.New(color.FgMagenta),
		StringMaxLength: 0,
		DisabledColor:   false,
		Indent:          2,
		RawStrings:      false,
	}
}

func (j *FormatterConfig) Raw(b []byte) (string, error) {
	return string(b), nil
}

func (j *FormatterConfig) Pretty(b []byte) (string, error) {
	var obj map[string]interface{}

	err := json.Unmarshal(b, &obj)
	if err != nil {
		return "", err
	}

	f := colorjson.NewFormatter()
	f.KeyColor = j.KeyColor
	f.StringColor = j.StringColor
	f.BoolColor = j.BoolColor
	f.NumberColor = j.NumberColor
	f.NullColor = j.NullColor
	f.StringMaxLength = j.StringMaxLength
	f.DisabledColor = j.DisabledColor
	f.Indent = j.Indent
	f.RawStrings = j.RawStrings

	s, err := f.Marshal(obj)

	return string(s), err
}
