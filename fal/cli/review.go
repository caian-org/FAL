package cli

import (
	"fmt"
)

type cmdReview struct {
	baseCmd
}

func (c *cmdReview) Run() error {
	config, err := c.getConfig()
	if err != nil {
		return err
	}

	name := config.Meta.Package.Name
	fmt.Printf("%s\n", name)
	return nil
}
