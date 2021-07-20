package command

import (
	"errors"
	"fmt"
	"github.com/guionardo/go-dev/cmd/configuration"
	"github.com/guionardo/go-dev/cmd/utils"
	"github.com/urfave/cli/v2"
	"log"
)

var (
	folders    []string
	justCD     bool
	openFolder bool
	GoCmd      = &cli.Command{
		Name:      "go",
		Usage:     "Go to folder",
		ArgsUsage: "[words for locate the folders]",
		Action:    GoAction,
		Before:    BeforeGoAction,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "just-cd",
				Destination: &justCD,
				Usage:       "Just go to folder, skip custom command",
			},
			&cli.BoolFlag{
				Name:        "open",
				Destination: &openFolder,
				Usage:       "Opens folder into file browser",
			},
		},
	}
)

func BeforeGoAction(context *cli.Context) error {
	configuration.SetupEnvironmentVars(context.String("basefolder"), context.String("config"))

	if !configuration.DefaultConfig.TryLoad(configuration.ConfigurationFileName) {
		log.Fatalf("Failed to read configuration file %s",configuration.ConfigurationFileName)
	}
	return nil
}

func GoAction(context *cli.Context) error {
	folders = context.Args().Slice()
	matches := configuration.DefaultConfig.Paths.FindFolder(folders)

	if len(matches) == 0 {
		return errors.New(fmt.Sprintf("Folder not found: %v", folders))
	}
	var match []string
	for _, m := range matches {
		match = append(match, m.Path)
	}
	var folder = utils.FolderChoice(match)
	if len(folder) == 0 {
		return errors.New("no folder choose")
	}
	path, _ := configuration.DefaultConfig.Paths.Get(folder)

	result := fmt.Sprintf("cd \"%s\"", folder)
	if openFolder {
		result = fmt.Sprintf("xdg-open \"%s\"", folder)
	} else
	if !justCD && len(path.Command) > 0 {
		result = fmt.Sprintf("%s && %s", result, path.Command)
	}
	fmt.Println(result)

	return nil
}
