package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type FALConfig struct {
	Meta       sectMeta         `yaml:"meta"`
	Targets    []string         `yaml:"targets"`
	Functions  []sectFunctions  `yaml:"functions"`
	Assemblies []sectAssemblies `yaml:"assemblies"`
}

func Load(projpath string) (*FALConfig, error) {
	cfile := filepath.Join(projpath, ".fal.yml")
	cdata, err := os.ReadFile(cfile)
	if err != nil {
		return nil, err
	}

	config := FALConfig{}
	err = yaml.Unmarshal(cdata, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
