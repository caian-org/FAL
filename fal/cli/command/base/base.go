package base

import (
	"fal/config"
)

type Command struct {
	Verbose bool   `short:"v" help:"Increase logging information"`
	Path    string `arg:"" name:"path" type:"path" help:"The FAL project location."`
}

func (c Command) GetConfig() (*config.FALConfig, error) {
	config, err := config.Load(c.Path)
	if err != nil {
		return nil, err
	}

	return config, nil
}
