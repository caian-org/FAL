package command

import (
	"fmt"
)

type Version struct{}

func (c *Version) Run() error {
	fmt.Println("VERSION")
	return nil
}
