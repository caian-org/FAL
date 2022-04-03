package manifest

// ~~~ meta

type metaPackage struct {
	Name        string `yaml:"name" json:"name"`
	Version     string `yaml:"version" json:"version"`
	Author      string `yaml:"author" json:"author"`
	Description string `yaml:"description" json:"description"`
}

type sectMeta struct {
	Package metaPackage `yaml:"package" json:"package"`
}

// ~~~ functions

type cacheStrategy struct {
	Filesystem           string `yaml:"filesystem" json:"filesystem"`
	Redis                string `yaml:"redis" json:"redis"`
	AutoCollectFsGarbage bool   `yaml:"auto-collect-fs-garbage" json:"auto-collect-fs-garbage"`
}

type cachePolicy struct {
	Enabled                bool          `yaml:"enabled" json:"enabled"`
	MaxSize                string        `yaml:"max-size" json:"max-size"`
	MaxRetained            int           `yaml:"max-retained" json:"max-retained"`
	MaxRetainedBuffer      int           `yaml:"max-retained-buffer" json:"max-retained-buffer"`
	RetainedLifespan       string        `yaml:"retained-lifespan" json:"retained-lifespan"`
	RetainedBufferLifespan string        `yaml:"retained-buffer-lifespan" json:"retained-buffer-lifespan"`
	Strategy               cacheStrategy `yaml:"strategy" json:"strategy"`
}

type funcOptions struct {
	Input        bool        `yaml:"input" json:"input"`
	Output       bool        `yaml:"output" json:"output"`
	DefaultInput string      `yaml:"default-input" json:"default-input"`
	CachePolicy  cachePolicy `yaml:"cache-policy" json:"cache-policy"`
}

type sectFunctions struct {
	Name string      `yaml:"name" json:"name"`
	Arn  string      `yaml:"arn" json:"arn"`
	With funcOptions `yaml:"with,omitempty" json:"with,omitempty"`
}

// ~~~ assemblies

type assemblySteps struct {
	Call string `yaml:"call" json:"call"`
}

type sectAssemblies struct {
	Name  string          `yaml:"name" json:"name"`
	Steps []assemblySteps `yaml:"steps" json:"steps"`
}

// ~~~ root manifest object

type FALManifest struct {
	Meta       sectMeta         `yaml:"meta" json:"meta"`
	Targets    []string         `yaml:"targets" json:"targets"`
	Functions  []sectFunctions  `yaml:"functions" json:"functions"`
	Assemblies []sectAssemblies `yaml:"assemblies" json:"assemblies"`
}
