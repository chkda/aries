package main

import (
	"os"

	"github.com/chkda/aries/internal/engine"
)

var FILE_LOC = "/config/engine/config.json"

func main() {
	currDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config := currDir + FILE_LOC
	engine.Start(config)
}
