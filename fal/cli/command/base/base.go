package base

import (
	"fal/manifest"
	"fal/shared/fs"
)

type Command struct {
	Verbose bool   `short:"v" help:"Increase logging information."`
	Path    string `arg:"" name:"path" type:"path" help:"The FAL project location."`
}

func (c Command) GetConfig(rootpath *fs.Location) (*manifest.FALManifest, error) {
	manifest, err := manifest.LoadAndValidate(rootpath)
	if err != nil {
		return nil, err
	}

	return manifest, nil
}
