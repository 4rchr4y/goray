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
	Handle(field *Field) (Handler, error)
}

func Walk(h Handler, field *Field) error {
	val := reflect.ValueOf(field.Value)

	switch val.Kind() {
	case reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Bool, reflect.Float32, reflect.Float64:
		field.Value = val.Interface()
		field.Kind = val.Kind()

	case reflect.Map:
		for _, key := range val.MapKeys() {
			child := val.MapIndex(key)
			childField := &Field{
				Value: child.Interface(),
				Path:  append(field.Path, key.String()),
				Kind:  child.Kind(),
			}
			if err := Walk(h, childField); err != nil {
				return err
			}
		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			childField := &Field{
				Value: val.Index(i).Interface(),
				Path:  append(field.Path, fmt.Sprintf("[%d]", i)),
				Kind:  val.Index(i).Kind(),
			}
			if err := Walk(h, childField); err != nil {
				return err
			}
		}
	}

	_, err := h.Handle(field)
	return err
}
