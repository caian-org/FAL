package command

import (
	"fal/cli/command/base"
	"fal/shared/fs"
	"fal/wrapper"
)

type Build struct {
	base.Command
}

func (c *Build) Run() error {
	rootlevel := fs.NewLocation(c.Path)
	config, err := c.GetConfig(rootlevel)
	if err != nil {
		return err
	}

	buildlevel := rootlevel.InnerLevel("_fal")
	err = buildlevel.CreateDir()
	if err != nil {
		return err
	}

	wrapper.InitSharedLib(buildlevel)

	wrapperlevel := buildlevel.InnerLevel("target")
	for _, lang := range config.Targets {
		langBuilder, err := wrapper.GetWrapperBuilderOf(lang)
		if err != nil {
			return nil
		}

		err = langBuilder(wrapperlevel)
		if err != nil {
			return err
		}
	}

	return nil
}
