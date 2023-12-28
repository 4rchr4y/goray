package rayfile

import "testing"

func TestVis(t *testing.T) {
	v := new(RayfileVisitor)

	Walk(v, map[string]string{
		"USER": "g10z3r",
		"TERM": "xterm-256color",
	})

	t.Fail()
}
