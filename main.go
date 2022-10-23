package main

import (
	"os"

	"github.com/guionardo/go-dev/cmd/commands"
	"github.com/guionardo/go-dev/pkg/consts"
	"github.com/guionardo/go-dev/pkg/logger"
)

var (
	build_date   string = "1970-01-01T00:00:00"
	build_runner string = "unknown"
	release      string = "0.0.0"
)

func main() {
	consts.SetupBuildInfo(build_date, build_runner, release)
	app := commands.SetupCli()

	if err := app.Run(os.Args); err != nil {
		logger.Error("%v", err)
		os.Exit(1)
	}
}
