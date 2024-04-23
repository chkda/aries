package main

import (
	"os"

	"github.com/chkda/aries/internal/pusher"
)

var FILE_LOC = "/config/pusher/config.json"

func main() {
	currDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config := currDir + FILE_LOC
	pusher.Start(config)
}
