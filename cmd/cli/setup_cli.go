package cli_

import (
	"fmt"

	"github.com/guionardo/go-dev/cmd/actions"
	"github.com/guionardo/go-dev/cmd/commands"
	"github.com/guionardo/go-dev/cmd/configuration"
	"github.com/guionardo/go-dev/cmd/ctx"
	"github.com/guionardo/go-dev/pkg/config"
	"github.com/urfave/cli/v2"
)

func SetupCli() *cli.App {
	var metadata = configuration.MetaData
	app := &cli.App{
		Name:        metadata.AppName,
		Version:     metadata.Version,
		Compiled:    metadata.CompileTime,
		Description: fmt.Sprintf("Builder Info: %s - %s", metadata.BuilderInfo, metadata.BuildDate),
		Usage:       "Go to your projects",
		Before:      ctx.ChainedActions(ctx.SetupContext),
		Action:      actions.GoAction,
		// DefaultCommand: "go",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config_file",
				Aliases: []string{"c"},
				EnvVars: []string{"GO_DEV_CONFIG"},
				Value:   config.DefaultConfigFile,
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				EnvVars: []string{"GO_DEV_DEBUG"},
				Value:   false,
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Value:config.DefaultOutputFile,
			},
		},
		Commands: []*cli.Command{
			commands.GetGoCommand(),
			commands.GetSetupCommand(),
			commands.GetSyncCommand(),
			commands.GetInitCommand(),
		},
	}
	return app
}
