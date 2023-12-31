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
			Path    string   `json:"path"`
			Targets []string `json:"targets"`
			Emm     []struct {
				K1 string `json:"k1"`
			} `json:"emm"`
		} `json:"workspace"`
	} `json:"data"`
}

func TestWalk(t *testing.T) {
	m := map[string]interface{}{
		"version": "1.1",
		"user":    "g10z3r",

		"data": map[string]interface{}{
			"term": "xterm-256color",
			// "set":  "foo",
			"path": "src/dir",
			// "set":  []string{"foo", "bar", "zip", "zap"},
			"workspace": []map[string]interface{}{
				{
					"path":    "root",
					"targets": []string{"t1", "t2", "t3"},
					"emm": []map[string]interface{}{
						{
							"k1": "v1",
						},
					},
				},
				{
					"path": "not-root",
				},
				{
					"path": "foo",
				},
			},
		},
	}

	ts := &testStruct{}
	v, err := NewBuilder(ts, m, WithMode(Autocomplete))
	if err != nil {
		fmt.Println(err)
	}

	Walk(v, &Field{
		Path:  nil,
		Value: m,
		Kind:  reflect.ValueOf(m).Kind(),
	})

	fmt.Println(ts.Data.Workspace[0].Emm)

	t.Fail()
}
