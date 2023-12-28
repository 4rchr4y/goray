package rayfile

type Rayfile struct {
	Version string `json:"version"`
}

func NewFromMap(data map[string]interface{}) (*Rayfile, error) {
	return nil, nil
}

func setField[T any](rf *Rayfile, key string, value T) error {
	return nil
}
