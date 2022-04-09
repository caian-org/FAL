//go:build _beforebuild

package main

import "fal/manifest"

func main() {
	manifest.CreateJsonSchemaFiles()
}
