package commands

import (
	"github.com/guionardo/go-dev/cmd/actions"
	"github.com/urfave/cli/v2"
)

func GetInitCommand() *cli.Command {
	return &cli.Command{
		Name:   "init",
		Usage:  "Initialize go-dev for bash",
		Action: actions.InitAction,
	}
}
