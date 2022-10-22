package commands

import (
	"github.com/guionardo/go-dev/cmd/actions"
	"github.com/guionardo/go-dev/cmd/ctx"
	"github.com/urfave/cli/v2"
)

func GetGoCommand() *cli.Command {
	return &cli.Command{
		Name:   "go",
		Usage:  "Execute go command",
		Before: ctx.ChainedActions(ctx.AssertConfigExists, ctx.AssertAutoSync),
		Action: actions.GoAction,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "just-cd",
				Usage: "Just go to folder, skip custom command",
			},
			&cli.BoolFlag{
				Name:  "open",
				Usage: "Opens folder into file browser",
			},
		},
	}
}
