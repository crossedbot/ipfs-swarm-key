package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/crossedbot/common/golang/logger"
	"github.com/crossedbot/ipfs-swarm-key/swarmkey"
)

const (
	FatalCode = 1
)

func main() {
	keyPath := flag.String(
		"key-path",
		"swarm.key",
		"The path the key will be saved",
	)
	flag.Parse()
	k, err := swarmkey.New()
	fatal(err, "failed to generate new key")
	fatal(
		ioutil.WriteFile(*keyPath, k.Bytes(), 0644),
		"failed to write file",
	)
}

// fatal prints the given message and exits on the given error
func fatal(err error, msg string) {
	if err != nil {
		if msg != "" {
			err = fmt.Errorf("%s: %s", msg, err)
		}
		logger.Error(err)
		os.Exit(FatalCode)
	}
}
