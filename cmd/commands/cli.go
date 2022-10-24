package commands

import (
	"fmt"

	"github.com/guionardo/go-dev/cmd/actions"
	"github.com/guionardo/go-dev/cmd/ctx"
	"github.com/guionardo/go-dev/pkg/config"
	"github.com/guionardo/go-dev/pkg/consts"
	"github.com/urfave/cli/v2"
)

func SetupCli() *cli.App {

	app := &cli.App{
		Name:        consts.AppName,
		Version:     consts.Metadata.Version,
		Compiled:    consts.Metadata.BuildTime,
		Description: fmt.Sprintf("Builder Info: %s - %s", consts.Metadata.BuildRunner, consts.Metadata.BuildTime.Format("2006-01-02 15:04:05")),
		Usage:       "Go to your projects",
		Before:      ctx.ChainedActions(ctx.SetupContext),
		Action:      actions.GoAction,
		Authors: []*cli.Author{
			{
				Name:  consts.AuthorName,
				Email: consts.AuthorEmail,
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    consts.FlagConfigFile,
				EnvVars: []string{"GO_DEV_CONFIG"},
				Value:   config.DefaultConfigFile,
			},
			&cli.BoolFlag{
				Name:    consts.FlagDebug,
				EnvVars: []string{"GO_DEV_DEBUG"},
				Value:   false,
			},
			&cli.StringFlag{
				Name:  consts.FlagOutput,
				Value: config.DefaultOutputFile,
			},
		},
		Commands: []*cli.Command{
			GetGoCommand(),
			GetSetupCommand(),
			GetSyncCommand(),
			GetInitCommand(),
		},
	}
	return app
}
