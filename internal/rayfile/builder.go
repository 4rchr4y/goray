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
	// if o.refval.Type().Elem() == reflect.TypeOf(field.Value) {
	// 	return true
	// }

	if field.Kind == reflect.Map {
		return true
	}
	// fmt.Println(strings.Join(field.Path, "."), o.refval.Type().Elem().Kind().String(), reflect.TypeOf(field.Value).Kind().String())
	return assignable(o.refval.Type().Elem(), reflect.TypeOf(field.Value))
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

// TODO refactor this shit
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

// TODO refactor this shit
func (b *Builder) buildTargetCache(v reflect.Value, path string) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem() // get the structure pointed to by the pointer
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		typeField := v.Type().Field(i)

		// defining a field tag
		fieldTag := strings.Split(typeField.Tag.Get(b.tag), ",")[0]
		if fieldTag == "" {
			fieldTag = typeField.Name
		}

		// building a path to the current field
		currentPath := path
		if currentPath != "" {
			currentPath += "."
		}
		currentPath += fieldTag

		// handling pointers to structures
		var effectiveField reflect.Value
		if field.Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.Struct {
			if field.IsNil() {
				// if the pointer is nil, create a new structure and set the pointer
				newStruct := reflect.New(field.Type().Elem())
				field.Set(newStruct)
			}
			effectiveField = field.Elem()
		} else {
			effectiveField = field
		}

		b.targetCache[currentPath] = &origin{
			name:   typeField.Name,
			path:   currentPath,
			refval: &effectiveField,
		}

		// processing slices of structures
		if field.Kind() == reflect.Slice {
			elemType := field.Type().Elem()
			isStructPtr := elemType.Kind() == reflect.Ptr && elemType.Elem().Kind() == reflect.Struct

			if elemType.Kind() == reflect.Struct || isStructPtr {
				sce := b.sourceCache[currentPath]

				if sce == nil {
					continue
				}

				for j := 0; j < sce.refval.Elem().Len(); j++ {
					if j == 0 {
						amount := sce.refval.Elem().Len()
						field.Set(reflect.MakeSlice(reflect.SliceOf(elemType), amount, amount))
					}

					// working with pointers to structures in a slice
					var elem reflect.Value
					if isStructPtr {
						elem = reflect.New(elemType.Elem())
						field.Index(j).Set(elem)
					} else {
						elem = reflect.New(elemType).Elem()
						field.Index(j).Set(elem)
					}

					b.buildTargetCache(field.Index(j), fmt.Sprintf("%s.[%d]", currentPath, j))
				}
			}
			continue
		}

		if effectiveField.Kind() != reflect.Struct {
			continue
		}

		// recursively build cache for nested structures
		b.buildTargetCache(effectiveField, currentPath)
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
	case assignable(origin.refval.Type(), fieldVal.Type()):

		fmt.Println(path)
		assign(*origin.refval, fieldVal)

	case b.mode == Autocomplete && origin.refval.Kind() == reflect.Slice && origin.isSupportedType(field):
		sliceType := origin.refval.Type()
		slice := reflect.MakeSlice(sliceType, 1, 1)

		assign(slice.Index(0), fieldVal)

		origin.refval.Set(slice)
		return
	}
}

func assignable(destType, valType reflect.Type) bool {
	// checking for pointer compatibility
	if destType.Kind() == reflect.Ptr && destType.Elem().Kind() == valType.Kind() {
		return true
	}

	if valType.Kind() == reflect.Ptr && valType.Elem().Kind() == destType.Kind() {
		return true
	}

	// checking for forward type compatibility
	if destType.Kind() != reflect.Ptr && valType.AssignableTo(destType) {
		return true
	}

	// checking for pointer and value compatibility
	if destType.Kind() == reflect.Ptr && valType.AssignableTo(destType.Elem()) {
		return true
	}

	if destType.Kind() == reflect.Slice && valType.Kind() == reflect.Slice {
		return assignable(destType.Elem(), valType.Elem())
	}

	return false
}

func assign(dest, val reflect.Value) {
	if !dest.CanSet() || !val.IsValid() {
		return
	}

	if dest.Kind() == reflect.Ptr && val.Kind() != reflect.Ptr {
		newPtr := reflect.New(dest.Type().Elem())
		newPtr.Elem().Set(val)
		dest.Set(newPtr)
		return
	}

	if dest.Kind() != reflect.Ptr && val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Type().AssignableTo(dest.Type()) {
		dest.Set(val)
		return
	}

	if dest.Type().Kind() == reflect.Slice && val.Type().Kind() == reflect.Slice {
		destIsPtrSlice := dest.Type().Elem().Kind() == reflect.Ptr
		valIsPtrSlice := val.Type().Elem().Kind() == reflect.Ptr

		// Подготавливаем новый слайс для результатов.
		newSlice := reflect.MakeSlice(dest.Type(), val.Len(), val.Cap())

		for i := 0; i < val.Len(); i++ {
			elem := val.Index(i)

			// Преобразовываем элементы в соответствии с типами слайсов.
			if destIsPtrSlice && !valIsPtrSlice {
				// Преобразуем строку в указатель на строку.
				newPtr := reflect.New(elem.Type())
				newPtr.Elem().Set(elem)
				newSlice.Index(i).Set(newPtr)
			} else if !destIsPtrSlice && valIsPtrSlice {
				// Преобразуем указатель на строку в строку.
				if !elem.IsNil() {
					newSlice.Index(i).Set(elem.Elem())
				}
			} else {
				// Типы совпадают, просто копируем.
				newSlice.Index(i).Set(elem)
			}
		}

		// Присваиваем новый слайс целевому значению.
		dest.Set(newSlice)
	}

	// fmt.Println(dest.Type().Kind().String(), val.Type().Kind().String(), val.Type().AssignableTo(dest.Type()))

}
