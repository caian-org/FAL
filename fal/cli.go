package main

import (
	"fmt"
	"github.com/alecthomas/kong"
)

type baseCmd struct {
	Verbose bool   `short:"v" help:"Increase logging information"`
	Path    string `arg:"" name:"path" type:"path" help:"The FAL project location."`
}

// ------- CMD: Build -------

type cmdBuild struct {
	baseCmd
}

func (c *cmdBuild) Run() error {
	config, err := LoadConfig(c.Path)
	if err != nil {
		return err
	}

	for _, f := range config.Functions {
		fmt.Printf("%s | %s\n", f.Arn, f.Name)
	}

	return nil
}

// ------- CMD: Create -------

type cmdCreate struct {
	baseCmd
}

func (c *cmdCreate) Run() error {
	fmt.Println("CREATE")
	return nil
}

// ------- CMD: Review -------

type cmdReview struct {
	baseCmd
}

func (c *cmdReview) Run() error {
	config, err := LoadConfig(c.Path)
	if err != nil {
		return err
	}

	name := config.Meta.Package.Name
	fmt.Printf("%s\n", name)
	return nil
}

// ------- CMD: Version -------

type cmdVersion struct{}

func (c *cmdVersion) Run() error {
	fmt.Println("VERSION")
	return nil
}

// >>> CLI parser

var cli struct {
	Version cmdVersion `cmd:"" help:"Show version information."`
	Build   cmdBuild   `cmd:"" help:"Do something."`
	Create  cmdCreate  `cmd:"" help:"Do something else."`
	Review  cmdReview  `cmd:"" help:"Do another something."`
}

func CliParserAndRun() {
	ctx := kong.Parse(&cli)
	err := ctx.Run()

	ctx.FatalIfErrorf(err)
}
