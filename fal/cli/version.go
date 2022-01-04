package cli

import (
	"fmt"
)

type cmdVersion struct{}

func (c *cmdVersion) Run() error {
	fmt.Println("VERSION")
	return nil
}
