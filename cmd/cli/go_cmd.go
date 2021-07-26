package command

import (
	"errors"
	"fmt"
	"log"
	"sort"

	"github.com/guionardo/go-dev/cmd/configuration"
	"github.com/guionardo/go-dev/cmd/utils"
	"github.com/urfave/cli/v2"
)

var AllowedCommandsFunctions = map[string]func(configuration.PathSetup) string{
	"vscode":  func(p configuration.PathSetup) string { return fmt.Sprintf("code \"%s\"", p.Path) },
	"disable": func(p configuration.PathSetup) string { return "" },
}

var AllowedCommands = func() []string {
	keys := make([]string, 0, len(AllowedCommandsFunctions))
	for k := range AllowedCommandsFunctions {
		keys = append(keys, k)
	}
	return keys
}()

var (
	folders    []string
	justCD     bool
	openFolder bool
	GoCmd      = &cli.Command{
		Name:      "go",
		Usage:     "Go to folder",
		ArgsUsage: "[words for locate the folders]",
		Action:    GoAction,
		Before:    BeforeActionLoadConfiguration,
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

func GoAction(context *cli.Context) error {
	folders = context.Args().Slice()
	matches := configuration.DefaultConfig.Paths.FindFolder(folders)

	if len(matches) == 0 {
		return fmt.Errorf("folder not found: %v", folders)
	}
	var match []string
	for _, m := range matches {
		match = append(match, m.Path)
	}
	sort.Strings(match)
	var folder = utils.FolderChoice(match, len(configuration.DevFolder))
	if len(folder) == 0 {
		return errors.New("no folder choose")
	}
	path, _ := configuration.DefaultConfig.Paths.Get(folder)

	command := parseCommand(path, openFolder, justCD)

	utils.WriteOutput(command)

	return nil
}

func parseCommand(path configuration.PathSetup, justOpenFolder bool, justCD bool) string {

	if justOpenFolder {
		return fmt.Sprintf("xdg-open \"%s\"", path.Path)
	}
	command := fmt.Sprintf("cd \"%s\"", path.Path)

	if (!justCD) && len(path.Command) > 0 {
		command = fmt.Sprintf("%s && %s", command, path.Command)
	}
	log.Printf("Running command: %s\n", command)
	return command
}
