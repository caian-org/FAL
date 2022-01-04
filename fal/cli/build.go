package cli

import (
	"fal/builder"
	"fal/util"
)

type cmdBuild struct {
	baseCmd
}

func (c *cmdBuild) Run() error {
	config, err := c.getConfig()
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
