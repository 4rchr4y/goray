package rfs

import (
	"bufio"
	"io"

	"github.com/4rchr4y/goray/internal/rayfile"
)

type tomlService interface {
	Decode(data string, v interface{}) error
}

type RayfileService struct {
	toml tomlService
}

func (rs *RayfileService) Parse(reader io.Reader) (*rayfile.Rayfile, error) {
	content := map[string]interface{}{}
	scanner := bufio.NewScanner(reader)

	if err := rs.toml.Decode(scanner.Text(), content); err != nil {
		return nil, err
	}

	rayfile, err := rayfile.NewFromMap(content)
	if err != nil {
		return nil, err
	}

	return rayfile, nil
}

func NewRayfileService(ts tomlService) *RayfileService {
	return &RayfileService{
		toml: ts,
	}
}
