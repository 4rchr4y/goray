package rayfile

import (
	"fmt"
	"reflect"
)

type Field struct {
	Path  []string
	Value interface{}
	Kind  reflect.Kind
}

type Handler interface {
	Handle(field *Field)
}

func Walk(h Handler, field *Field) {
	val := reflect.ValueOf(field.Value)

	switch val.Kind() {
	case reflect.Map:
		for _, key := range val.MapKeys() {
			child := val.MapIndex(key)
			Walk(h, &Field{
				Value: child.Interface(),
				Path:  append(append([]string(nil), field.Path...), key.String()),
				Kind:  val.Kind(),
			})
		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			Walk(h, &Field{
				Value: val.Index(i).Interface(),
				Path:  append(append([]string(nil), field.Path...), fmt.Sprintf("[%d]", i)),
				Kind:  val.Index(i).Kind(),
			})
		}
	}

	if val.Kind() != reflect.Map && val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		field.Value = val.Interface()
		field.Kind = val.Kind()
	}

	h.Handle(field)
}
