package main

import (
	"github.com/alecthomas/kong"
)

type DiscordGuildProbeCLI struct {
	Read    ReadCmd    `cmd:"" help:"Reads all exported channels (JSON format) in directory"`
	Version VersionCmd `cmd:"" help:"Print version of the dgp CLI"`
}

type Context struct{}

func main() {
	ctx := kong.Parse(&DiscordGuildProbeCLI{})
	err := ctx.Run(&Context{})

	ctx.FatalIfErrorf(err)
}
