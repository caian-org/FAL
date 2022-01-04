package cli

import (
	"github.com/alecthomas/kong"
)

var cli struct {
	Version cmdVersion `cmd:"" help:"Show version information."`
	Build   cmdBuild   `cmd:"" help:"Do something."`
	Create  cmdCreate  `cmd:"" help:"Do something else."`
	Review  cmdReview  `cmd:"" help:"Do another something."`
}

func ParseAndRun() {
	ctx := kong.Parse(&cli)
	err := ctx.Run()

	ctx.FatalIfErrorf(err)
}
