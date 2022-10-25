package commands

import (
	"github.com/guionardo/go-dev/cmd/actions"
	"github.com/guionardo/go-dev/cmd/ctx"
	"github.com/guionardo/go-dev/pkg/consts"
	"github.com/urfave/cli/v2"
)

func GetGoCommand() *cli.Command {
	return &cli.Command{
		Name:   "go",
		Usage:  "Execute go command",
		Before: ctx.ChainedActions(ctx.AssertAutoUpdate, ctx.AssertConfigExists, ctx.AssertAutoSync),
		Action: actions.GoAction,
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
		},
	}
}
