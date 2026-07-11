package app

import (
	"log"

	"github.com/dxo1a/obscusync/internal/cli"
)

func Run() {
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
