package main

import (
	"os"

	"github.com/guionardo/go-dev/cmd/commands"
	"github.com/guionardo/go-dev/pkg/logger"
)

func main() {
	app := commands.SetupCli()

	if err := app.Run(os.Args); err != nil {
		logger.Error("%v", err)
		os.Exit(1)
	}
}
