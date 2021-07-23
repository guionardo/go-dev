package main

import (
	command "github.com/guionardo/go-dev/cmd/cli"
	"github.com/guionardo/go-dev/cmd/configuration"
	"github.com/guionardo/go-dev/cmd/utils"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	err := configuration.SetupBaseEnvironment()
	if err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
	
	var metadata = configuration.LoadMetaData()
	app := &cli.App{
		Name:        metadata.AppName,
		Version:     metadata.Version,
		Compiled:    metadata.CompileTime,
		//Description: "Go to your projects",
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
				Name:  "Guionardo Furlan",
				Email: "guionardo@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    command.ConfigArg,
				EnvVars: []string{"GO-DEV-CONFIG"},
				Value:   configuration.DefaultFolderConfigFile,
				Usage:   "Configuration file",
			},
			&cli.StringFlag{
				Name:        "output",
				EnvVars:     []string{"GO-DEV-OUTPUT"},
				Value:       configuration.DefaultOutputFile,
				Usage:       "Output file for command execution",
				Destination: &utils.OutputFileName,
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
	return nil
}
