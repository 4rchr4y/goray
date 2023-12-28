package rayfile

import "fmt"

type RayfileVisitor struct{}

func (rfv *RayfileVisitor) Visit(field *Field) (Visitor, error) {
	fmt.Println(field.Path, field.Value)

	return rfv, nil
}
