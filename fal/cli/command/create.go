package command

import (
	"fal/cli/command/base"
	"fmt"
)

type Create struct {
	base.Command
}

func (c *Create) Run() error {
	fmt.Println("CREATE")
	return nil
}
