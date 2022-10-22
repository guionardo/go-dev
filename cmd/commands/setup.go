package commands

import (
	"os"

	"github.com/guionardo/go-dev/cmd/actions"
	"github.com/guionardo/go-dev/cmd/ctx"
	"github.com/urfave/cli/v2"
)

func GetSetupCommand() *cli.Command {
	return &cli.Command{
		Name:   "setup",
		Usage:  "Setup go-dev",
		Before: ctx.ChainedActions(ctx.AssertAtLeastDefault),
		Subcommands: []*cli.Command{
			GetSetupAddFolderCommand(),
			GetSetupAutoSyncCommand(),
		},
	}
}

func GetSetupAddFolderCommand() *cli.Command {
	currentFolder, _ := os.Getwd()
	return &cli.Command{
		Name:      "add-folder",
		Usage:     "Add a folder to go-dev",
		Action:    actions.SetupAddFolderAction,
		ArgsUsage: "[folder default=" + currentFolder + "]",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "max-depth",
				Aliases: []string{"d"},
				Usage:   "Max depth to search for subfolders",
				Value:   3,
			},
			&cli.BoolFlag{
				Name:    "ignore-children",
				Aliases: []string{"i"},
				Usage:   "Ignore children of this folder",
				Value:   false,
			},
		},
	}
}

func GetSetupAutoSyncCommand() *cli.Command {
	return &cli.Command{
		Name:   "auto-sync",
		Usage:  "Enable or disable auto-sync",
		Action: actions.SetupAutoSyncAction,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "enable",
				Aliases: []string{"e"},
				Usage:   "Enable auto-sync",
				Value:   true,
			},
			&cli.IntFlag{
				Name:    "interval",
				Aliases: []string{"i"},
				Usage:   "Interval (minutes) to run auto-sync",
				Value:   360,
			},
		},
	}
}

func GetSetupShell() *cli.Command {
	return &cli.Command{
		Name:   "shell",
		Usage:  "Setup shell",		
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "disable",
				Aliases: []string{"d"},
				Usage:   "Disable shell",
				Value:   false,
			},
		},
		Action: actions.SetupShellAction,
	}
}

