package config

type Def interface{}

type PolicyDef struct {
	Path    string
	Target  []string
	Include []string
}