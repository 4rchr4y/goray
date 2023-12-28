package rayfile

import (
	"github.com/4rchr4y/goray/internal/service/toml"
	"github.com/4rchr4y/goray/internal/syswrap"
)

// ToDos:
// Implement Validate function

type LookupFn func(key string) (string, bool)

type Config struct {
	Workspace *workspace               `toml:"workspace"`
	Analysis  map[string]*AnalysisConf `toml:"analysis"`
}

type workspace struct {
	Version    string   `toml:"version"`
	RootDir    string   `toml:"root"`
	PolicyDir  string   `toml:"policies"`
	IgnoreList []string `toml:"ignore-list"`
	GoArch     string   `toml:"go-arch"`
}

type AnalysisConf struct {
	Level  string
	Target string
}

type ConfOptFn func(*Config)

func NewConfig(options ...ConfOptFn) *Config {
	conf := &Config{
		Workspace: &workspace{
			Version:    defaultVersion,
			RootDir:    defaultRoot,
			PolicyDir:  defaultPolicies,
			GoArch:     defaultGoArch,
			IgnoreList: defaultIgnoreList,
		},
	}

	for i := 0; i < len(options); i++ {
		options[i](conf)
	}

	return conf
}

func WithRootDir(dirPath string) ConfOptFn {
	return func(c *Config) {
		c.Workspace.RootDir = dirPath
	}
}

func WithPolicyDir(dirPath string) ConfOptFn {
	return func(c *Config) {
		c.Workspace.PolicyDir = dirPath
	}
}

func WithGoArch(goArch string) ConfOptFn {
	return func(c *Config) {
		c.Workspace.GoArch = goArch
	}
}

func WithAnalysis(analysis map[string]*AnalysisConf) ConfOptFn {
	return func(c *Config) {
		c.Analysis = analysis
	}
}

type ConfReadFileOptFn func(*ReadConf)

type fsClient interface {
	ReadFile(name string) ([]byte, error)
}

type tomlService interface {
	Decode(data string, v interface{}) error
}

type ReadConf struct {
	fsClient
	tomlService
}

func NewFromFile(options ...ConfReadFileOptFn) *ReadConf {
	readConf := &ReadConf{
		fsClient:    syswrap.FsClient{},
		tomlService: &toml.TomlService{},
	}

	for i := 0; i < len(options); i++ {
		options[i](readConf)
	}

	return readConf
}

func NewConfigFromFile(filePath string, options ...ConfReadFileOptFn) (*Config, error) {
	readConf := NewFromFile(options...)

	conf := NewConfig()

	data, err := readConf.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// dataWithEnv, err := readConf.Interpolate(string(data))
	// if err != nil {
	// 	return nil, err
	// }

	if err := readConf.Decode(string(data), &conf); err != nil {
		return nil, err
	}

	return conf, nil
}

func (cfg *Config) Validate() error {
	return nil
}
