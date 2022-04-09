//go:build _beforebuild

package manifest

import (
	"os"
	"path/filepath"
	"strings"

	"fal/shared/fs"
	"sigs.k8s.io/yaml"
)

func convertYamlFileToJson(path string, info os.FileInfo, err error) error {
	if err != nil {
		panic(err)
	}

	if info.IsDir() || !strings.HasSuffix(path, ".yml") {
		return nil
	}

	schemaLevel := fs.NewLocation(filepath.Dir(path))
	schemaFilename := info.Name()

	yamlSchema, err := schemaLevel.ReadFile(schemaFilename)
	if err != nil {
		panic(err)
	}

	jsonSchema, err := yaml.YAMLToJSON(yamlSchema)
	if err != nil {
		panic(err)
	}

	safeFL := []rune(schemaFilename)
	jsonSchemaFilename := string(safeFL[:len(safeFL)-4]) + ".json"

	_, err = schemaLevel.CreateFile(jsonSchemaFilename, jsonSchema)
	if err != nil {
		panic(err)
	}

	return nil
}

func CreateJsonSchemaFiles() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filepath.Walk(filepath.Join(pwd, "manifest", "manifest.json"), convertYamlFileToJson)
}
