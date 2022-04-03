package manifest

import (
	"fal/shared/fs"
	"gopkg.in/yaml.v2"
)

func LoadAndValidate(rootpath *fs.Location) (*FALManifest, error) {
	data, err := rootpath.ReadFile(".fal.yml")
	if err != nil {
		return nil, err
	}

	manifest := FALManifest{}
	err = yaml.Unmarshal(data, &manifest)
	if err != nil {
		return nil, err
	}

	err = validateSchema(&manifest)
	if err != nil {
		return nil, err
	}

	return &manifest, nil
}
