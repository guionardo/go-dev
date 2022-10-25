package commands

import (
	"github.com/guionardo/go-dev/cmd/actions"
	"github.com/guionardo/go-dev/cmd/ctx"
	"github.com/guionardo/go-dev/pkg/consts"
	"github.com/urfave/cli/v2"
)

func GetSyncCommand() *cli.Command {
	return &cli.Command{
		Name:   "sync",
		Usage:  "Sync go-dev",
		Before: ctx.ChainedActions(ctx.AssertConfigExists),
		After:  ctx.ChainedActions(ctx.AssertSaveIfNotError),
		Action: actions.SyncAction,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     consts.FlagMaxDept,
				Aliases:  []string{"d"},
				Usage:    "Max depth to search for subfolders",
				Required: false,
			},
		},
	}
}
