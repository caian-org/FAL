package config

import (
	"fal/util"
	"gopkg.in/yaml.v2"
)

func LoadAndValidate(rootpath *util.Location) (*FALConfig, error) {
	data, err := rootpath.ReadFile(".fal.yml")
	if err != nil {
		return nil, err
	}

	config := FALConfig{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	err = validateSchema(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
