package actions

import (
	"errors"
	"fmt"
	"log"
	"sort"

	"github.com/guionardo/go-dev/cmd/ctx"
	"github.com/guionardo/go-dev/cmd/utils"
	"github.com/guionardo/go-dev/pkg/folders"
	"github.com/urfave/cli/v2"
)

func GoAction(c *cli.Context) error {
	if c.Args().Len() == 0 {
		return errors.New("No arguments provided")
	}
	c2 := ctx.GetContext(c)
	found := make([]*folders.Folder, 0)
	devFoldersFound := make([]string, 0, len(c2.Config.DevFolders))
	for df, collection := range c2.Config.DevFolders {
		f := collection.Find(c.Args().Slice())
		if len(f) > 0 {
			found = append(found, f...)
			devFoldersFound = append(devFoldersFound, df)
		}
	}
	if len(found) == 0 {
		return fmt.Errorf("No folders found for %s", c.Args().Slice())
	}

	match := make([]string, len(found))
	i := 0
	for _, m := range found {
		match[i] = m.Path
		i++
	}
	var folder *folders.Folder
	// devFolder := ""
	// ok := false
	maxDevFolder := 0
	for _, df := range devFoldersFound {
		if len(df) > maxDevFolder {
			maxDevFolder = len(df)
		}
	}
	// for df, coll := range c2.Config.DevFolders {
	// 	if folder, ok = coll.Folders[choosed_folder]; ok {
	// 		devFolder = df
	// 		break
	// 	}
	// }
	sort.Strings(match)
	choosed_folder := utils.FolderChoice(match, maxDevFolder)
	if len(choosed_folder) == 0 {
		return errors.New("no folder choose")
	}
	for _, df := range devFoldersFound {
		if folder, _ = c2.Config.DevFolders[df].Folders[choosed_folder]; folder != nil {
			break
		}
	}

	openFolder := c.Bool("open-folder")
	justCD := c.Bool("just-cd")
	output := c.String("output")
	utils.SetOutput(output)
	command := parseCommand(folder, openFolder, justCD)

	utils.WriteOutput(command)

	return nil
}

func parseCommand(folder *folders.Folder, justOpenFolder bool, justCD bool) string {

	if justOpenFolder {
		return fmt.Sprintf("xdg-open \"%s\"", folder.Path)
	}
	command := fmt.Sprintf("cd \"%s\"", folder.Path)

	if (!justCD) && len(folder.Command) > 0 {
		command = fmt.Sprintf("%s && %s", command, folder.Command)
	}
	log.Printf("Running command: %s\n", command)
	return command
}
