package main

import (
	"github.com/austinvalle/discord-guild-probe/pkg/discord"
)

type VersionCmd struct{}

func (cmd *VersionCmd) Run(ctx *Context) error {
	discord.Test("Hello World")
	return nil
}
