package command

import (
	"fal/cli/command/base"
	"fal/shared/fs"
	"fal/shared/log"
	"fal/wrapper"
)

type Build struct {
	base.Command
}

func (c *Build) Run() error {
	log.Init(c.Verbose)
	log.DebugF("Build routine started", log.Fields{"path": c.Path})

	rootlevel := fs.NewLocation(c.Path)
	manifest, err := c.LoadManifest(rootlevel)
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
	for _, lang := range manifest.Targets {
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
