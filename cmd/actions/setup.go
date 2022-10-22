package actions

import (
	"fmt"
	"os"
	"time"

	"github.com/guionardo/go-dev/cmd/ctx"
	"github.com/guionardo/go-dev/pkg/folders"
	"github.com/guionardo/go-dev/pkg/logger"
	"github.com/urfave/cli/v2"
)

func SetupAddFolderAction(c *cli.Context) error {
	c2 := ctx.GetContext(c)
	if c.NArg() < 1 {
		return fmt.Errorf("Missing folder name")
	}
	folderName := c.Args().First()
	if folderName == "" {
		folderName, _ = os.Getwd()
	}
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		return fmt.Errorf("Folder %s not found", folderName)
	}
	if _, ok := c2.Config.DevFolders[folderName]; ok {
		return fmt.Errorf("Folder %s already added", folderName)
	}
	maxDepth := c.Int("max-depth")
	if maxDepth <= 0 {
		return fmt.Errorf("Max depth must be greater than 0")
	}
	collection := folders.CreateCollection(folderName, maxDepth)
	if err := collection.Sync(); err != nil {
		return err
	}
	c2.Config.DevFolders[folderName] = collection
	return c2.Config.Save(c2.ConfigFile)

}

func SetupAutoSyncAction(c *cli.Context) error {
	c2 := ctx.GetContext(c)
	interval := c.Int("interval")
	if interval == 0 {
		logger.Info("AutoSync disabled")
	} else {
		logger.Info("AutoSync enabled with interval %s", time.Duration(interval*int(time.Minute)))
	}
	c2.Config.AutoSync.Interval = time.Duration(interval * int(time.Minute))

	return c2.Config.Save(c2.ConfigFile)
}

func SetupShellAction(c *cli.Context) error {
//	source <(./go-dev init)

	c2 := ctx.GetContext(c)
	if c.NArg() < 1 {
		return fmt.Errorf("Missing shell name")
	}
	shellName := c.Args().First()
	if shellName == "" {
		return fmt.Errorf("Missing shell name")
	}
	c2.Config.Shell = shellName
	return c2.Config.Save(c2.ConfigFile)
}