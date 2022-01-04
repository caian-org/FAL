package config

// ~~~ meta

type metaPackage struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Author      string `yaml:"author"`
	Description string `yaml:"description"`
}

type sectMeta struct {
	Package metaPackage `yaml:"package"`
}

// ~~~ functions

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

// ~~~ assemblies

type assemblySteps struct {
	Call string `yaml:"call"`
}

type sectAssemblies struct {
	Name  string          `yaml:"name"`
	Steps []assemblySteps `yaml:"steps"`
}
