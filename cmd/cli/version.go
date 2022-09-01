package main

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/austinvalle/discord-guild-probe/pkg/discord"
)

type VersionCmd struct{}

var commit = "local"

func init() {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, s := range info.Settings {
			if s.Key == "vcs.revision" {
				commit = s.Value
			}
		}
	}
}

func (cmd *VersionCmd) Run(ctx *Context) error {
	fmt.Printf("dgp - commit: %s - %s/%s\n", commit, runtime.GOOS, runtime.GOARCH)

	discord.Test("Hello World")
	return nil
}
