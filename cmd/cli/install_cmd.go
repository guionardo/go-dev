package command

import (
	_ "embed"
	"fmt"
	"github.com/guionardo/go-dev/cmd/configuration"
	"github.com/urfave/cli/v2"
)

//go:embed dev.sh
var dev_sh string

var (
	installRemove   bool
	devBaseFolder   string
	maximumSubLevel int
	InstallCmd      = &cli.Command{
		Name:  "install",
		Usage: "Install go-dev",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "uninstall",
				Usage:       "Remove installation",
				Destination: &installRemove,
			},
			&cli.PathFlag{
				Name:        "basefolder",
				Usage:       "Development base folder",
				Value:       configuration.DefaultDevFolder(),
				Destination: &devBaseFolder,
			},
			&cli.IntFlag{
				Name:        "max-path-level",
				Usage:       "Maximum level of paths",
				Value:       configuration.MaximumSubLevel,
				Destination: &maximumSubLevel,
			},
		},
		Action: InstallAction,
		Before: BeforeInstallAction,
	}
)

func InstallAction(context *cli.Context) error {
	newConfig := &configuration.ConfigFileType{
		DevFolder:         devBaseFolder,
		Paths:             make(configuration.Paths),
		ConfigurationFile: context.String("config"),
	}
	var err error
	if err = newConfig.Paths.ReadFolders(devBaseFolder, maximumSubLevel); err == nil {
		if err = newConfig.Save(); err == nil {
			fmt.Printf("Configuration file saved @ %s (base folder = %s)\n", newConfig.ConfigurationFile, newConfig.DevFolder)
			if err = installScript(); err == nil {
				if err = installAlias(); err == nil {
					fmt.Println("Alias setup is done")
				}
			}
		}
	}

	return err
}

func installScript() error {
	return nil
}

func installAlias() error {
	return nil
}

func BeforeInstallAction(context *cli.Context) error {
	return nil
}
