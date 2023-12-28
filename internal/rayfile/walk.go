package rayfile

import (
	"errors"
	"fmt"
	"reflect"
)

type Field struct {
	Path  []string
	Value interface{}
	Kind  reflect.Kind
}

type Visitor interface {
	Visit(field *Field) (Visitor, error)
}

func Walk(v Visitor, field *Field) error {
	visitor, err := v.Visit(field)
	if err != nil {
		return err
	}

	if visitor == nil {
		return errors.New("visitor is nil")
	}

	val := reflect.ValueOf(field.Value)
	switch val.Kind() {
	case reflect.String:
		return nil

	case reflect.Map:
		for _, key := range val.MapKeys() {
			child := val.MapIndex(key)
			childField := &Field{
				Value: child.Interface(),
				Path:  append(field.Path, key.String()),
				Kind:  child.Kind(),
			}
			if err := Walk(v, childField); err != nil {
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
			if err := Walk(v, childField); err != nil {
				return err
			}
		}
	}

	return nil
}
