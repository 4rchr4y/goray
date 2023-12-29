package rayfile

import (
	"errors"

	"reflect"
	"strings"
)

type Mode int

const (
	Autocomplete Mode = iota + 1
)

type origin struct {
	refval reflect.Value // original reflect value
}

func (o *origin) isSupportedType(field *Field) bool {
	return o.refval.Type().Elem() == reflect.TypeOf(field.Value)
}

type Builder[T any] struct {
	val   T
	tag   string
	cache map[string]*origin
	mode  Mode
}

func NewBuilder[T any](value T, tag string, mode Mode) (*Builder[T], error) {
	v := reflect.ValueOf(value).Elem()

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, errors.New("invalid type")
	}

	b := &Builder[T]{
		val:   value,
		tag:   tag,
		cache: make(map[string]*origin),
		mode:  mode,
	}

	b.buildCache(v, "")

	return b, nil
}

func (b *Builder[T]) buildCache(v reflect.Value, path string) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		typeField := v.Type().Field(i)

		fieldTag := strings.Split(typeField.Tag.Get(b.tag), ",")[0]
		if fieldTag == "" {
			fieldTag = typeField.Name
		}

		// building the path to the current field
		currentPath := path
		if currentPath != "" {
			currentPath += "."
		}

		currentPath += fieldTag
		b.cache[currentPath] = &origin{
			refval: field,
		}

		if !(field.Kind() == reflect.Struct || (field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct)) {
			continue // field is not a structure or a pointer to a structure
		}

		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}

		// building a cache for embed structs
		b.buildCache(field, currentPath)
	}
}

func (b *Builder[T]) Handle(field *Field) {
	path := strings.Join(field.Path, ".")
	origin, exists := b.cache[path]
	if !exists {
		return
	}

	if !origin.refval.CanSet() {
		return
	}

	fieldVal := reflect.ValueOf(field.Value)

	switch {
	case fieldVal.Type().AssignableTo(origin.refval.Type()):
		origin.refval.Set(fieldVal)

	case b.mode == Autocomplete && origin.refval.Kind() == reflect.Slice && origin.isSupportedType(field):
		sliceType := origin.refval.Type()
		slice := reflect.MakeSlice(sliceType, 1, 1)
		slice.Index(0).Set(fieldVal)
		origin.refval.Set(slice)
	}
}
