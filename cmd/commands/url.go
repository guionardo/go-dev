package commands

import (
	"github.com/guionardo/go-dev/cmd/actions"
	"github.com/urfave/cli/v2"
)

func GetUrlCommand() *cli.Command {
	return &cli.Command{
		Name:    "url",
		Aliases: []string{"u"},
		Usage:   "Open URL from repository into current browser",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "just-print",
				Aliases: []string{"j"},
				Usage:   "Just show the URL",
			},
		},
		Action: actions.UrlAction,
	}
}