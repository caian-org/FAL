package command

import (
	"fal/cli/command/base"
	"fal/util"
	"fmt"
)

type Review struct {
	base.Command
}

func (c *Review) Run() error {
	rootpath := util.NewLocation(c.Path)
	config, err := c.GetConfig(rootpath)
	if err != nil {
		return err
	}

	name := config.Meta.Package.Name
	fmt.Printf("%s\n", name)
	return nil
}
