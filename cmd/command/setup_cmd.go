package command

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/guionardo/go-dev/cmd/configuration"
	"github.com/urfave/cli/v2"
)

const folderArg = "folder"
const enableArg = "enable"
const disableArg = "disable"
const disableSubFolders = "disable-subs"
const commandArg = "command"

var (
	SetupFolder      string
	setupEnable      bool
	setupDisable     bool
	setupDisableSubs bool
	setupCommand     string
	SetupCmd         = &cli.Command{
		Name:  "setup",
		Usage: "Setup configuration for folder",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        folderArg,
				Aliases:     []string{"f"},
				Value:       configuration.CurrentDir(),
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
			&cli.StringFlag{
				Name:        commandArg,
				Destination: &setupCommand,
				Usage:       fmt.Sprintf("Command to be executed after change to path [%s]", AllowedCommands),
			},
		},
		Action: SetupAction,
		Before: BeforeActionLoadConfiguration,
	}
)

func SetupAction(*cli.Context) error {
	err := parseEnableDisable()
	if err == nil {
		err = parseDisableSubFolders()
		if err == nil {
			err = parseSetCommand()
		}
	}

	return err
}

func parseSetCommand() error {
	path, err := configuration.DefaultConfig.Paths.Get(SetupFolder)
	if err == nil {
		command := "_"
		for key, value := range AllowedCommandsFunctions {
			if strings.EqualFold(key, setupCommand) {
				command = value(path)
				break
			}
		}
		if command == "_" {
			return fmt.Errorf("invalid command %s, expected %v", setupCommand, AllowedCommands)
		}

		path.Command = command
		if len(setupCommand) > 0 {
			log.Printf("Command: %s = %s\n", path.Path, command)
		} else {
			log.Printf("Command: %s [DISABLED]\n", path.Path)
		}
		if path.Command != setupCommand {
			configuration.DefaultConfig.Paths.Set(path)
			err = configuration.DefaultConfig.Save()
		}
	}
	return err
}

func setEnableDisable(enable bool) error {
	path, err := configuration.DefaultConfig.Paths.Get(SetupFolder)
	if err != nil {
		return err
	}
	if path.Ignore != enable {
		return nil
	}
	path.Ignore = !enable
	err = configuration.DefaultConfig.Save()
	return err
}

func parseEnableDisable() error {
	if setupEnable && setupDisable {
		return fmt.Errorf("%s and %s flags are mutually exclusive", enableArg, disableArg)
	}
	var err error
	if setupEnable {
		log.Printf("Enabling folder %s\n", SetupFolder)
		err = setEnableDisable(true)
	} else if setupDisable {
		log.Printf("Disabling folder %s\n", SetupFolder)
		err = setEnableDisable(false)
	}
	return err
}

func parseDisableSubFolders() error {
	if !setupDisableSubs {
		return nil
	}
	log.Printf("Disabling sub-folders for %s\n", SetupFolder)
	var changed = false
	var setupFolder = SetupFolder + string(os.PathSeparator)
	for _, folder := range configuration.DefaultConfig.Paths.FolderList() {

		if strings.HasPrefix(folder, setupFolder) && folder != SetupFolder {
			path, err := configuration.DefaultConfig.Paths.Get(folder)
			if err == nil && !path.Ignore {
				path.Ignore = true
				err = configuration.DefaultConfig.Paths.Set(path)
				if err == nil {
					changed = true
					log.Printf("+ %s\n", folder)
				} else {
					log.Printf("! %s (%v)\n", folder, err)
				}

			}

		}
	}
	if changed {
		if err := configuration.DefaultConfig.Save(); err == nil {
			log.Println("Updated configuration")
		} else {
			log.Fatalf("Failed to update configuration: %v", err)
		}
	}
	return nil
}
