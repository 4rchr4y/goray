package rayfile

import (
	"fmt"
	"strings"
)

type Builder[T any] struct {
	value T
}

func (rfv *Builder[T]) Handle(field *Field) (Handler, error) {
	fmt.Println(field.Kind.String(), strings.Join(field.Path, "."), "|", field.Value)

	return rfv, nil
}
