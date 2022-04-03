package command

import (
	"fal/builder"
	"fal/cli/command/base"
	"fal/shared/fs"
)

type Build struct {
	base.Command
}

func (c *Build) Run() error {
	rootlvl := fs.NewLocation(c.Path)
	config, err := c.GetConfig(rootlvl)
	if err != nil {
		return err
	}

	buildlvl := rootlvl.InnerLevel("build")
	err = buildlvl.CreateDir()
	if err != nil {
		return err
	}

	builder.InitSharedLib(buildlvl)

	wrapperlvl := buildlvl.InnerLevel("wrapper")
	for _, lang := range config.Targets {
		langBuilder, err := builder.GetWrapperBuilderOf(lang)
		if err != nil {
			return nil
		}

		err = langBuilder(wrapperlvl)
		if err != nil {
			return err
		}
	}

	return nil
}
