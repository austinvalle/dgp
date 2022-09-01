package main

import (
	"github.com/alecthomas/kong"
)

type DiscordGuildProbeCLI struct {
	Version VersionCmd `cmd:"" help:"Print version of the discord-guild-probe CLI"`
}

type Context struct{}

func main() {
	ctx := kong.Parse(&DiscordGuildProbeCLI{})
	err := ctx.Run(&Context{})

	ctx.FatalIfErrorf(err)
}
