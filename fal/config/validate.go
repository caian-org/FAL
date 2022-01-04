package config

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/qri-io/jsonschema"
)

var (
	//go:embed schema.json
	_schema_doc []byte
)

type schema = jsonschema.Schema

func formatJsonSchemaError(i int, e jsonschema.KeyError) string {
	m := "[JsonSchemaError #%d] at property '%s': %s"
	return fmt.Sprintf(m, i, e.PropertyPath, e.Message)
}

func getSchema() (*schema, error) {
	s := &schema{}
	if err := json.Unmarshal(_schema_doc, s); err != nil {
		return nil, err
	}

	return s, nil
}

func validateSchema(config *FALConfig) error {
	s, err := getSchema()
	if err != nil {
		return err
	}

	ctx := context.Background()
	configEncoded, err := json.Marshal(config)
	if err != nil {
		return err
	}

	if errs, _ := s.ValidateBytes(ctx, configEncoded); len(errs) > 0 {
		errorList := []string{}
		for i, err := range errs {
			errorList = append(errorList, formatJsonSchemaError(i+1, err))
		}

		return errors.New(strings.Join(errorList, "\n"))
	}

	return nil
}
