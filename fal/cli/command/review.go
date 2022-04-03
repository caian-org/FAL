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
	config, err := c.GetConfig(rootpath)
	if err != nil {
		return err
	}

	name := config.Meta.Package.Name
	fmt.Printf("%s\n", name)
	return nil
}
