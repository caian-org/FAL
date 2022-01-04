package cli

import (
	"fmt"
)

type cmdCreate struct {
	baseCmd
}

func (c *cmdCreate) Run() error {
	fmt.Println("CREATE")
	return nil
}
