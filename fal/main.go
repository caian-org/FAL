//go:build !_beforebuild

package main

import "fal/cli"

func main() {
	cli.ParseAndRun()
}
