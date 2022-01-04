package cli

import (
	"fal/cli/command"
	"github.com/alecthomas/kong"
)

var cli struct {
	Version command.Version `cmd:"" help:"Show version information."`
	Build   command.Build   `cmd:"" help:"Do something."`
	Create  command.Create  `cmd:"" help:"Do something else."`
	Review  command.Review  `cmd:"" help:"Do another something."`
}

func ParseAndRun() {
	ctx := kong.Parse(&cli)
	err := ctx.Run()

	ctx.FatalIfErrorf(err)
}
