package main

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

// ------- SECT: Meta -------

type metaPackage struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Author      string `yaml:"author"`
	Description string `yaml:"description"`
}

type sectMeta struct {
	Package metaPackage `yaml:"package"`
}

// ------- SECT: Functions -------

type cacheStrategy struct {
	Filesystem           string `yaml:"filesystem"`
	Redis                string `yaml:"redis"`
	AutoCollectFsGarbage bool   `yaml:"auto-collect-fs-garbage"`
}

type cachePolicy struct {
	Enabled                bool          `yaml:"enabled"`
	MaxSize                string        `yaml:"max-size"`
	MaxRetained            int           `yaml:"max-retained"`
	MaxRetainedBuffer      int           `yaml:"max-retained-buffer"`
	RetainedLifespan       string        `yaml:"retained-lifespan"`
	RetainedBufferLifespan string        `yaml:"retained-buffer-lifespan"`
	Strategy               cacheStrategy `yaml:"strategy"`
}

type funcOptions struct {
	Input        bool        `yaml:"input"`
	Output       bool        `yaml:"output"`
	DefaultInput string      `yaml:"default-input"`
	CachePolicy  cachePolicy `yaml:"cache-policy"`
}

type sectFunctions struct {
	Name string      `yaml:"name"`
	Arn  string      `yaml:"arn"`
	With funcOptions `yaml:"with,omitempty"`
}

// ------- SECT: Assemblies -------

type assemblySteps struct {
	Call string `yaml:"call"`
}

type sectAssemblies struct {
	Name  string          `yaml:"name"`
	Steps []assemblySteps `yaml:"steps"`
}

// >>> FAL Config

type FALConfig struct {
	Meta       sectMeta         `yaml:"meta"`
	Targets    []string         `yaml:"targets"`
	Functions  []sectFunctions  `yaml:"functions"`
	Assemblies []sectAssemblies `yaml:"assemblies"`
}

func LoadConfig(projpath string) (*FALConfig, error) {
	configFile := filepath.Join(projpath, ".fal.yml")
	configData, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	config := FALConfig{}
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
