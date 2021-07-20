package main

import (
	command "github.com/guionardo/go-dev/cmd/cli"
	"github.com/guionardo/go-dev/cmd/configuration"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	err := configuration.SetupBaseEnvironment()
	if err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
	//configuration.Setup()
	var metadata = configuration.LoadMetaData()
	app := &cli.App{
		Name:        metadata.AppName,
		Version:     metadata.Version,
		Compiled:    metadata.CompileTime,
		Description: "Go to your projects",
		Commands: []*cli.Command{
			command.SetupCmd,
			command.GoCmd,
			command.ListCmd,
			command.UpdateCmd,
			command.InstallCmd,
		},
		Authors: []*cli.Author{
			{
				Name:  "Guionardo Furlan",
				Email: "guionardo@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				EnvVars: []string{"GO-DEV-CONFIG"},
				Value:   configuration.DefaultFolderConfigFile,
				Usage:   "Configuration file",
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

func BeforeMainAction(context *cli.Context) error {
	return nil
}
