package rayfile

import "fmt"

type RayfileVisitor struct{}

func (rfv *RayfileVisitor) Visit(field interface{}) Visitor {
	fmt.Println(field)
	return rfv
}
