package rayfile

import (
	"fmt"
	"strings"
)

type RayfileVisitor struct{}

func (rfv *RayfileVisitor) Visit(field *Field) (Visitor, error) {
	fmt.Println(field.Kind.String(), strings.Join(field.Path, "."), "|", field.Value)

	return rfv, nil
}
