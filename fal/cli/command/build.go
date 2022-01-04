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
	rootpath := util.NewLocation(c.Path)
	config, err := c.GetConfig(rootpath)
	if err != nil {
		return err
	}

	buildDir := rootpath.InnerLevel("build")
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
