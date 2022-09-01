package discord_test

import (
	_ "embed"
	"testing"
	"testing/fstest"

	"github.com/austinvalle/dgp/pkg/discord"
)

//go:embed testdata/empty.json
var emptyJson []byte

func TestLoadChannel(t *testing.T) {

	testFs := fstest.MapFS{
		"empty.json": &fstest.MapFile{Data: emptyJson},
	}

	_, err := discord.LoadChannelsFromDir(testFs)
	if err != nil {
		t.Fatal(err)
	}
}
