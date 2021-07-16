package command

import (
	"github.com/guionardo/go-dev/cmd/configuration"
	"github.com/urfave/cli/v2"
)

var (
	UpdateCmd = &cli.Command{
		Name:   "update",
		Usage:  "Update folders",
		Action: UpdateAction,
		Before: BeforeUpdateAction,
	}
)

func UpdateAction(context *cli.Context) error {
	err := configuration.Config.ReadFolders(configuration.DevFolder, configuration.MaximumSubLevel)
	if err != nil {
		err = configuration.Config.Save(configuration.DevFolderConfig)
	}

	return err
}

func BeforeUpdateAction(context *cli.Context) error {
	return nil
}
