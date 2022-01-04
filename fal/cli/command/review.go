package command

import (
	"fal/cli/command/base"
	"fmt"
)

type Review struct {
	base.Command
}

func (c *Review) Run() error {
	config, err := c.GetConfig()
	if err != nil {
		return err
	}

	name := config.Meta.Package.Name
	fmt.Printf("%s\n", name)
	return nil
}
