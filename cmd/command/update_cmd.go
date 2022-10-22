package command

import (
	"github.com/guionardo/go-dev/cmd/configuration"
	"github.com/urfave/cli/v2"
	"log"
)

var (
	UpdateCmd = &cli.Command{
		Name:   "update",
		Usage:  "Update folders",
		Action: UpdateAction,
		Before: BeforeActionLoadConfiguration,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "max-path-level",
				Usage:       "Maximum level of paths",
				Value:       configuration.MaximumSubLevel,
				Destination: &configuration.MaxFolderLevel,
			},
		},
	}
)

func UpdateAction(*cli.Context) error {
	log.Println("Updating folders")
	err := configuration.DefaultConfig.Paths.ReadFolders(configuration.DevFolder, configuration.MaxFolderLevel)
	if err == nil {
		err = configuration.DefaultConfig.Save()
	}

	return err
}
