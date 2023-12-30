package rayfile

import (
	"encoding/json"
	"errors"
	"fmt"

	"reflect"
	"strings"
)

type Mode int

const (
	// Autocomplete is a mode used in the Builder to indicate that if a single value is provided
	// for a field that expects a slice, the single value should be automatically converted
	// into a slice with that single value as its element. This is particularly useful when you
	// want to simplify the setting of fields that are expected to be slices but are common to
	// have a single value, avoiding the need to manually wrap the value in a slice.
	Autocomplete Mode = iota + 1
)

type origin struct {
	refval reflect.Value // original reflect value
}

func (o *origin) isSupportedType(field *Field) bool {
	if o.refval.Type().Elem() == reflect.TypeOf(field.Value) {
		return true
	}

	if field.Kind == reflect.Map {
		return true
	}

	return false
}

// func (c *cache) load(key string) (*origin, bool) {

// }

type Builder[T any] struct {
	val   T
	tag   string             // e.g. 'json', 'xml', 'toml'
	cache map[string]*origin // cache of paths to structure fields
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

	// started creating a cache from the root
	b.buildCache(v, "")

	d, _ := json.Marshal(b.cache)

	fmt.Println(string(d))

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

		if field.Kind() == reflect.Slice && field.Type().Elem().Kind() == reflect.Struct {
			elemType := field.Type().Elem()
			elem := reflect.New(elemType).Elem()
			b.buildCache(elem, fmt.Sprintf("%s.[#]", currentPath))
			continue
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
	fmt.Println(strings.Join(field.Path, "."), "|", field.Value, field.Kind.String())

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

	case b.mode == Autocomplete && origin.refval.Kind() == reflect.Slice && origin.isSupportedType(field) && field.Kind != reflect.Map:
		sliceType := origin.refval.Type()
		slice := reflect.MakeSlice(sliceType, 1, 1)
		slice.Index(0).Set(fieldVal)
		origin.refval.Set(slice)
		return

		// case origin.refval.Kind() == reflect.Slice && field.Kind == reflect.Map:
		// 	fmt.Println(strings.Join(field.Path, "."), field.Kind)
	}
}
