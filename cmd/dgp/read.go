package main

import (
	"fmt"
	"os"

	"github.com/austinvalle/dgp/pkg/discord"
)

type ReadCmd struct {
	Path string `arg:"" type:"existingdir"`
}

func (cmd *ReadCmd) Run(ctx *Context) error {
	dir := os.DirFS(cmd.Path)

	exports, err := discord.LoadChannelsFromDir(dir)
	if err != nil {
		return err
	}
	for _, export := range exports {
		fmt.Printf("Message Count: '%d'\n", export.MessageCount)
	}

	return nil
}
