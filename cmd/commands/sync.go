package commands

import (
	"github.com/guionardo/go-dev/cmd/actions"
	"github.com/guionardo/go-dev/cmd/ctx"
	"github.com/urfave/cli/v2"
)

func GetSyncCommand() *cli.Command {
	return &cli.Command{
		Name:   "sync",
		Usage:  "Sync go-dev",
		Before: ctx.ChainedActions(ctx.AssertConfigExists),
		Action: actions.SyncAction,
	}
}
