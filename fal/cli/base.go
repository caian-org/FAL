package cli

import "fal/config"

type baseCmd struct {
	Verbose bool   `short:"v" help:"Increase logging information"`
	Path    string `arg:"" name:"path" type:"path" help:"The FAL project location."`
}

func (b baseCmd) getConfig() (*config.FALConfig, error) {
	config, err := config.Load(b.Path)
	if err != nil {
		return nil, err
	}

	return config, nil
}
