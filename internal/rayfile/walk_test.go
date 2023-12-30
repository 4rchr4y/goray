package rayfile

import (
	"fmt"
	"reflect"
	"testing"
)

type testStruct struct {
	Version string `json:"version"`
	User    string `json:"user"`
	Data    struct {
		Term      string   `json:"term"`
		Set       []string `json:"set"`
		Workspace []struct {
			Path    string   `json:"path" group/item:"#"`
			Targets []string `json:"targets"`
		} `json:"workspace" group:"workspace"`
	} `json:"data"`
}

func TestWalk(t *testing.T) {
	ts := &testStruct{}
	v, err := NewBuilder[*testStruct](ts, "json", Autocomplete)
	if err != nil {
		fmt.Println(err)
	}

	m := map[string]interface{}{
		"version": "1.1",
		"user":    "g10z3r",

		"data": map[string]interface{}{
			"term": "xterm-256color",
			// "set":  "foo",
			"path": "src/dir",
			"set":  []string{"foo", "bar", "zip", "zap"},
			"workspace": []map[string]interface{}{
				{
					"path":    "root",
					"targets": []string{"t1", "t2", "t3"},
				},
				{
					"path": "not-root",
				},
			},
		},
	}

	Walk(v, &Field{
		Path:  make([]string, 0),
		Value: m,
		Kind:  reflect.ValueOf(m).Kind(),
	})

	fmt.Println(len(ts.Data.Workspace))

	t.Fail()
}
