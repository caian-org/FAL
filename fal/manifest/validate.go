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

	//go:embed manifest.json/target/schema.json
	_schema_doc_target string

	//go:embed manifest.json/function/schema.json
	_schema_doc_func string

	//go:embed manifest.json/function/cache-policy.json
	_schema_doc_func_cache string

	//go:embed manifest.json/assembly/schema.json
	_schema_doc_assembly string

	//go:embed manifest.json/assembly/step/call.json
	_schema_doc_assembly_step_call string

	//go:embed manifest.json/assembly/step/map.json
	_schema_doc_assembly_step_map string

	//go:embed manifest.json/assembly/step/series.json
	_schema_doc_assembly_step_series string
)

func loadJSON(c string) gjs.JSONLoader {
	return gjs.NewStringLoader(c)
}

func validateSchema(manifest *FALManifest) error {
	loader := gjs.NewSchemaLoader()
	loader.AddSchemas(
        loadJSON(_schema_doc_meta),
        loadJSON(_schema_doc_target),
        loadJSON(_schema_doc_func),
        loadJSON(_schema_doc_func_cache),
        loadJSON(_schema_doc_assembly),
        loadJSON(_schema_doc_assembly_step_call),
        loadJSON(_schema_doc_assembly_step_series),
        loadJSON(_schema_doc_assembly_step_map),
    )

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
