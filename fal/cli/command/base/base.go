package base

import (
	"fal/config"
	"fal/util"
)

type Command struct {
	Verbose bool   `short:"v" help:"Increase logging information"`
	Path    string `arg:"" name:"path" type:"path" help:"The FAL project location."`
}

func (c Command) GetConfig(rootpath *util.Location) (*config.FALConfig, error) {
	config, err := config.LoadAndValidate(rootpath)
	if err != nil {
		return nil, err
	}

	return config, nil
}
