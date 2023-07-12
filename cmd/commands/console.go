package commands

import (
	"github.com/guionardo/go-dev/cmd/actions"
	"github.com/guionardo/go-dev/cmd/ctx"
	"github.com/guionardo/go-dev/pkg/consts"
	"github.com/urfave/cli/v2"
)

func GetConsoleCommand() *cli.Command {
	return &cli.Command{
		Name:   "console",
		Usage:  "Open console",
		Before: ctx.ChainedActions(ctx.AssertAutoUpdate, ctx.AssertConfigExists, ctx.AssertAutoSync),
		Action: actions.ConsoleAction,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  consts.FlagJustCD,
				Usage: "Just go to folder, skip custom command",
			},
			&cli.BoolFlag{
				Name:  consts.FlagOpen,
				Usage: "Opens folder into file browser",
			},
			&cli.StringFlag{
				Name:     consts.FlagOutput,
				Usage:    "File to save command output",
				Required: false,
			},
			&cli.StringFlag{
				Name:  consts.FlagChoiceType,
				Usage: "Type of choice to be used when multiple folders are found [browse, index]",
				Value: "browse",
			},
		},
	}
}
