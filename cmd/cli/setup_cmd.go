package command

import (
	"errors"
	"fmt"
	"github.com/guionardo/go-dev/cmd/utils"
	"github.com/urfave/cli/v2"
	"log"
	"path/filepath"
)

const folderArg = "folder"
const enableArg = "enable"
const disableArg = "disable"
const disableSubFolders = "disable-subs"

var (
	SetupFolder      string
	setupEnable      bool
	setupDisable     bool
	setupDisableSubs bool
	SetupCmd         = &cli.Command{
		Name:    "setup",
		Usage:   "Setup configuration for folder",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        folderArg,
				Aliases:     []string{"f"},
				Value:       ".",
				Usage:       "Folder",
				Destination: &SetupFolder,
			},
			&cli.BoolFlag{
				Name:        enableArg,
				Destination: &setupEnable,
				Usage:       "Enable folder",
			},
			&cli.BoolFlag{
				Name:        disableArg,
				Destination: &setupDisable,
				Usage:       "Disable folder",
			},
			&cli.BoolFlag{
				Name:        disableSubFolders,
				Destination: &setupDisableSubs,
				Usage:       "Disable all sub folders",
			},
		},
		Action: SetupAction,
		Before: BeforeAction,
	}
)

func BeforeAction(context *cli.Context) error {
	if !utils.PathExists(SetupFolder) {
		return errors.New(fmt.Sprintf("Path not found: %s", SetupFolder))
	}

	folder, err := filepath.Abs(SetupFolder)
	if err == nil {
		SetupFolder = folder
		return nil
	}
	return errors.New(fmt.Sprintf("Failed to get absolute path of %s : %v", SetupFolder, err))
}

func SetupAction(context *cli.Context) error {
	err := parseEnableDisable()
	if err != nil {
		return err
	}
	err = parseDisableSubFolders()

	return nil
}

func parseEnableDisable() error {
	if setupEnable && setupDisable {
		return errors.New(fmt.Sprintf("%s and %s flags are mutually exclusive", enableArg, disableArg))
	}
	if setupEnable {
		log.Printf("Enabling folder %s\n", SetupFolder)
	} else if setupDisable {
		log.Printf("Disabling folder %s\n", SetupFolder)
	}
	return nil
}

func parseDisableSubFolders() error {
	if !setupDisableSubs {
		return nil
	}
	log.Printf("Disabling sub-folders for %s\n", SetupFolder)
	return nil
}
