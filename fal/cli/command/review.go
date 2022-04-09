package command

import (
	"fal/cli/command/base"
	"fal/shared/fs"
	"fmt"
)

type Review struct {
	base.Command
}

func (c *Review) Run() error {
	rootpath := fs.NewLocation(c.Path)
	manifest, err := c.LoadManifest(rootpath)
	if err != nil {
		return err
	}

	name := manifest.Meta.Package.Name
	fmt.Printf("%s\n", name)
	return nil
}
