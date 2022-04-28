package main

import (
	"fmt"
	"log"
	"os"

	command "github.com/guionardo/go-dev/cmd/cli"
	"github.com/guionardo/go-dev/cmd/configuration"
	"github.com/guionardo/go-dev/cmd/debug"
	"github.com/guionardo/go-dev/cmd/utils"

	"github.com/urfave/cli/v2"
)

func main() {
	debug.Debug("Starting")
	err := configuration.SetupBaseEnvironment()
	if err != nil {
		log.Fatalf("Failed to start: %v", err)
	}

	var metadata = configuration.MetaData
	app := &cli.App{
		Name:        metadata.AppName,
		Version:     metadata.Version,
		Compiled:    metadata.CompileTime,
		Description: fmt.Sprintf("Builder Info: %s - %s", metadata.BuilderInfo, metadata.BuildDate),

		Usage: "Go to your projects",
		Commands: []*cli.Command{
			command.GoCmd,
			command.SetupCmd,
			command.ListCmd,
			command.UpdateCmd,
			command.InstallCmd,
		},
		Authors: []*cli.Author{
			{
				Name:  metadata.AuthorName,
				Email: metadata.AuthorEmail,
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    command.ConfigArg,
				EnvVars: []string{"GO_DEV_CONFIG"},
				Value:   configuration.DefaultFolderConfigFile,
				Usage:   "Configuration file",
			},
			&cli.StringFlag{
				Name:        "output",
				EnvVars:     []string{"GO_DEV_OUTPUT"},
				Value:       configuration.DefaultOutputFile,
				Usage:       "Output file for command execution",
				Destination: &utils.OutputFileName,
			},
			&cli.BoolFlag{
				Name:        "debug",
				Value:       false,
				Usage:       "Enable debug logging",
				Destination: &debug.Enabled,
			},
		},
		Before: BeforeMainAction,
		Action: command.GoAction,
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func BeforeMainAction(*cli.Context) error {
	debug.Debug("")
	return nil
}
