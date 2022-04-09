//go:build !_beforebuild

package manifest

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	gjs "github.com/xeipuuv/gojsonschema"
)

var (
	//go:embed manifest.json/schema.json
	_schema_doc_root string

	//go:embed manifest.json/meta/schema.json
	_schema_doc_meta string
)

func loadJSON(c string) gjs.JSONLoader {
	return gjs.NewStringLoader(c)
}

func validateSchema(manifest *FALManifest) error {
	loader := gjs.NewSchemaLoader()
	loader.AddSchemas(loadJSON(_schema_doc_meta))

	schema, err := loader.Compile(loadJSON(_schema_doc_root))
	if err != nil {
		return err
	}

	manifestEncoded, err := json.Marshal(manifest)
	if err != nil {
		return err
	}

	result, err := schema.Validate(loadJSON(string(manifestEncoded)))
	if err != nil {
		panic(err)
	}

	if !result.Valid() {
		errorList := []string{}

		for i, e := range result.Errors() {
			errMsg := fmt.Sprintf("[ManifestSchemaError #%d] at property '%s' --> %s", i, e.Field(), e.Description())
			errorList = append(errorList, errMsg)
		}

		return errors.New(strings.Join(errorList, "\n"))
	}

	return nil
}
