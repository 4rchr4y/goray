package rayfile

import (
	"fmt"
)

type Field struct {
	Key   string
	Value interface{}
}

type Visitor interface {
	Visit(field interface{}) Visitor
}

func Walk(v Visitor, field interface{}) {
	if v = v.Visit(field); v == nil {
		return
	}

	switch f := field.(type) {
	case string:

	case []string:
		Walk(v, field)

	case int:
		fmt.Println(f)
	}

}
