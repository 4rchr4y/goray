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
	name   string
	path   string         // field path in the struct, e.g. data.embed.some
	refval *reflect.Value // original reflect value
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

type Builder struct {
	value       any
	tag         string             // e.g. 'json', 'xml', 'toml'
	targetCache map[string]*origin // cache of paths to structure fields
	sourceCache map[string]*origin // cache of paths in source map
	mode        Mode
}

type BuilderOptFn func(*Builder)

func NewBuilder(value any, source map[string]any, options ...BuilderOptFn) (*Builder, error) {
	v := reflect.ValueOf(value).Elem()

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, errors.New("invalid type")
	}

	b := &Builder{
		value:       value,
		tag:         "json",
		targetCache: make(map[string]*origin),
		sourceCache: make(map[string]*origin),
		mode:        0,
	}

	for i := range options {
		options[i](b)
	}

	b.buildSourceCache(reflect.ValueOf(source), "")

	// started creating a cache from the root
	b.buildTargetCache(v, "")

	d, _ := json.Marshal(b.sourceCache)

	fmt.Println(string(d))

	d, _ = json.Marshal(b.targetCache)

	fmt.Println(string(d))

	return b, nil
}

func WithTag(tag string) BuilderOptFn {
	return func(b *Builder) {
		b.tag = tag
	}
}

func WithMode(mode Mode) BuilderOptFn {
	return func(b *Builder) {
		b.mode = mode
	}
}

func (b *Builder) buildSourceCache(v reflect.Value, path string) {
	for _, key := range v.MapKeys() {
		val := v.MapIndex(key)

		currentPath := path
		if currentPath != "" {
			currentPath += "."
		}

		currentPath += key.String()

		switch {
		case val.Elem().Kind() == reflect.Map:
			b.buildSourceCache(val.Elem(), currentPath)

		case val.Elem().Kind() == reflect.Slice:
		SVL:
			for i := 0; i < val.Elem().Len(); i++ {
				elem := val.Elem().Index(i)
				if elem.Kind() == reflect.Map {
					b.buildSourceCache(elem, fmt.Sprintf("%s.[%d]", currentPath, i))
					continue SVL
				}
			}
		}

		b.sourceCache[currentPath] = &origin{
			name:   key.String(),
			path:   currentPath,
			refval: &val,
		}
	}
}

func (b *Builder) buildTargetCache(v reflect.Value, path string) {
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
		b.targetCache[currentPath] = &origin{
			name:   typeField.Name,
			path:   currentPath,
			refval: &field,
		}

		if field.Kind() == reflect.Slice && field.Type().Elem().Kind() == reflect.Struct {
			elemType := field.Type().Elem()
			sce := b.sourceCache[currentPath] // source cache element

			if sce == nil {
				continue
			}

		SVL:
			for j := 0; j < sce.refval.Elem().Len(); j++ {
				if j == 0 {
					amount := sce.refval.Elem().Len()
					field.Set(reflect.MakeSlice(reflect.SliceOf(elemType), amount, amount))
				}

				elem := reflect.New(elemType).Elem()
				field.Index(j).Set(elem)
				b.buildTargetCache(field.Index(j), fmt.Sprintf("%s.[%d]", currentPath, j))
				continue SVL
			}

			continue
		}

		if !(field.Kind() == reflect.Struct || (field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct)) {
			continue // field is not a structure or a pointer to a structure
		}

		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}

		// building a cache for embed structs
		b.buildTargetCache(field, currentPath)
	}
}

func (b *Builder) Handle(field *Field) {
	path := strings.Join(field.Path, ".")
	origin, exists := b.targetCache[path]
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
		return
	}
}

func trimLastOccurrence(s, substr string) string {
	idx := strings.LastIndex(s, substr)
	if idx == -1 {
		return s
	}

	return s[:idx]
}
