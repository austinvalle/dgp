package discord

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
)

func LoadChannelsFromDir(dir fs.FS) ([]*DceExport, error) {
	matches, err := fs.Glob(dir, "*.json")
	if err != nil {
		return nil, err
	}

	exports := make([]*DceExport, 0)
	for _, match := range matches {
		json, err := dir.Open(match)
		if err != nil {
			fmt.Printf("skipping json '%s' because of err: %v\n", match, err)
			continue
		}
		defer json.Close()

		export, err := parseJson(json)
		if err != nil {
			fmt.Printf("skipping json '%s' because of err: %v\n", match, err)
			continue
		}

		exports = append(exports, export)
	}

	return exports, nil
}

func parseJson(jsonFile fs.File) (*DceExport, error) {
	bytes, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var export DceExport
	err = json.Unmarshal(bytes, &export)
	if err != nil {
		return nil, err
	}

	return &export, nil
}
