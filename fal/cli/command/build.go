package command

import (
	"fal/builder"
	"fal/cli/command/base"
	"fal/util"
)

type Build struct {
	base.Command
}

func (c *Build) Run() error {
	config, err := c.GetConfig()
	if err != nil {
		return err
	}

	buildDir := util.NewLocation(c.Path).InnerLevel("build")
	err = buildDir.CreateDir()
	if err != nil {
		return err
	}

	builder.InitSharedLib(buildDir)

	wrappersDir := buildDir.InnerLevel("wrappers")
	for _, lang := range config.Targets {
		langBuilder, err := builder.GetWrapperBuilderOf(lang)
		if err != nil {
			return nil
		}

		err = langBuilder(wrappersDir)
		if err != nil {
			return err
		}
	}

	return nil
}
