package commands

import (
	"os"
	"strings"

	"github.com/guionardo/go-dev/cmd/actions"
	"github.com/guionardo/go-dev/cmd/ctx"
	"github.com/guionardo/go-dev/pkg/consts"
	"github.com/guionardo/go-dev/pkg/folders"
	"github.com/urfave/cli/v2"
)

func GetSetupCommand() *cli.Command {
	return &cli.Command{
		Name:   "setup",
		Usage:  "Setup go-dev",
		Before: ctx.ChainedActions(ctx.AssertAtLeastDefault),
		Subcommands: []*cli.Command{
			GetSetupAddFolderCommand(),
			GetSetupUpdateFolderCommand(),
			GetSetupAutoSyncCommand(),
			GetSetupShellCommand(),
			GetSetupAutoUpdateCommand(),
			GetSetupDoGitUpdateCommand(),
			GetSetupShowCommand(),
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
				Name:    consts.FlagMaxDept,
				Aliases: []string{"d"},
				Usage:   "Max depth to search for subfolders",
				Value:   3,
			},
			&cli.BoolFlag{
				Name:    consts.FlagIgnoreChildren,
				Aliases: []string{"i"},
				Usage:   "Ignore children of this folder",
				Value:   false,
			},
			&cli.StringFlag{
				Name:    consts.FlagCommand,
				Aliases: []string{"c"},
				Usage:   "Command to execute when folder is selected",
			},
		},
	}
}

func GetSetupUpdateFolderCommand() *cli.Command {
	currentFolder, _ := os.Getwd()
	allowedCommands := strings.Join(folders.AllowedCommands, ", ")
	return &cli.Command{
		Name:      "update-folder",
		Usage:     "Add a folder to go-dev",
		Action:    actions.SetupUpdateFolderAction,
		After:     ctx.ChainedActions(ctx.AssertSaveIfNotError),
		ArgsUsage: "[folder default=" + currentFolder + "]",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    consts.FlagIgnoreChildren,
				Aliases: []string{"i"},
				Usage:   "Ignore children of this folder",
				Value:   false,
			},
			&cli.StringFlag{
				Name:    consts.FlagCommand,
				Aliases: []string{"c"},
				Usage:   "Command to execute when folder is selected [" + allowedCommands + "]",
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
				Name:  consts.FlagDisable,
				Usage: "Disable auto-sync",
				Value: true,
			},
			&cli.IntFlag{
				Name:    consts.FlagInterval,
				Aliases: []string{"i"},
				Usage:   "Interval (minutes) to run auto-sync",
				Value:   360,
			},
		},
	}
}

func GetSetupAutoUpdateCommand() *cli.Command {
	return &cli.Command{
		Name:   "auto-update",
		Usage:  "Enable or disable auto-update",
		Action: actions.SetupAutoUpdateAction,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  consts.FlagDisable,
				Usage: "Disable auto-update",
				Value: true,
			},
			&cli.IntFlag{
				Name:    consts.FlagInterval,
				Aliases: []string{"i"},
				Usage:   "Interval (minutes) to run auto-autoupdate",
				Value:   360,
			},
		},
	}
}

func GetSetupShellCommand() *cli.Command {
	return &cli.Command{
		Name:  "shell",
		Usage: "Setup shell",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    consts.FlagDisable,
				Aliases: []string{"d"},
				Usage:   "Disable shell",
				Value:   false,
			},
		},
		Action: actions.SetupShellAction,
	}
}

func GetSetupDoGitUpdateCommand() *cli.Command {
	return &cli.Command{
		Name:   "do-git-update",
		Usage:  "Update go-dev",
		Action: actions.SetupDoGitUpdateAction,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    consts.FlagNoFolders,
				Aliases: []string{"n"},
			},
		},
	}
}

func GetSetupShowCommand() *cli.Command {
	return &cli.Command{
		Name:   "show",
		Usage:  "Show go-dev configuration",
		Action: actions.SetupShowAction,
	}
}
