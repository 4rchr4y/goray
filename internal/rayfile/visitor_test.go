package rayfile

import (
	"reflect"
	"testing"
)

func TestVis(t *testing.T) {
	v := new(RayfileVisitor)

	m := map[string]interface{}{
		"USER": "g10z3r",
		"TERM": "xterm-256color",
		"MAP": map[string]interface{}{
			"version": 1.0,
			"workspace": map[string]interface{}{
				"root": ".",
			},
		},
	}

	Walk(v, &Field{
		Path:  make([]string, 0),
		Value: m,
		Kind:  reflect.ValueOf(m).Kind(),
	})

	t.Fail()
}
