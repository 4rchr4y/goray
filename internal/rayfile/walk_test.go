package rayfile

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	v := new(RayfileVisitor)

	m := map[string]interface{}{
		"user": "g10z3r",
		"term": "xterm-256color",
		"data": map[string]interface{}{
			"version": 1.0,
			"workspace": map[string]interface{}{
				"root": ".",
			},
			"set": []string{"foo", "bar", "zip", "zap"},
		},
	}

	Walk(v, &Field{
		Path:  make([]string, 0),
		Value: m,
		Kind:  reflect.ValueOf(m).Kind(),
	})

	t.Fail()
}
