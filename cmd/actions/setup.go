package actions

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/guionardo/go-dev/cmd/ctx"
	"github.com/guionardo/go-dev/pkg/arrays"
	"github.com/guionardo/go-dev/pkg/consts"
	"github.com/guionardo/go-dev/pkg/folders"
	"github.com/guionardo/go-dev/pkg/logger"
	"github.com/guionardo/go-dev/pkg/shell"
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
	maxDepth := c.Int(consts.FlagMaxDept)
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

func SetupUpdateFolderAction(c *cli.Context) (err error) {
	c2 := ctx.GetContext(c)
	if c.NArg() < 1 {
		return fmt.Errorf("Missing folder name")
	}
	folderName := c.Args().First()
	if folderName == "" {
		folderName, _ = os.Getwd()
	}
	if _, err = os.Stat(folderName); os.IsNotExist(err) {
		return fmt.Errorf("Folder %s not found", folderName)
	}
	var folder *folders.Folder
	for _, collection := range c2.Config.DevFolders {
		if folder, err = collection.Get(folderName); folder != nil {
			break
		}
	}
	if folder == nil {
		return fmt.Errorf("Folder %s not found - run a sync command", folderName)
	}
	changed := false
	if c.IsSet(consts.FlagIgnoreChildren) && c.Bool(consts.FlagIgnoreChildren) != folder.IgnoreSubFolders {
		folder.IgnoreSubFolders = c.Bool(consts.FlagIgnoreChildren)
		changed = true
	}
	if c.IsSet(consts.FlagCommand) && c.String(consts.FlagCommand) != folder.Command {
		folder.SetCommand(c.String(consts.FlagCommand))
		changed = true
	}

	if !changed {
		return errors.New("Nothing to update")
	}

	return c2.Config.Save(c2.ConfigFile)
}

func SetupAutoSyncAction(c *cli.Context) error {
	c2 := ctx.GetContext(c)
	interval := c.Int(consts.FlagInterval)
	if c.Bool(consts.FlagDisable) {
		interval = 0
	}
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
	executableName, err := os.Executable()
	if err != nil {
		return err
	}
	if strings.HasSuffix(executableName, "__debug_bin") {
		// Running from vscode
		return errors.New("Not supported from vscode")
	}
	sourceLine := fmt.Sprintf("source <(%s init)", executableName)

	disable := c.Bool(consts.FlagDisable)

	shellInfo, err := shell.NewShellInfo()
	if err != nil {
		return err
	}
	lines, err := arrays.LoadFromFile(shellInfo.RCFile)
	if err != nil {
		return err
	}

	lineNo, line, err := lines.FindByLine(func(l string) bool {
		l = strings.TrimSpace(l)
		return strings.HasPrefix(l, "source <(") && (strings.HasSuffix(l, "go-dev init)") || strings.HasSuffix(l, executableName+" init"))
	})
	operation := ""
	if disable {
		if err != nil {
			return fmt.Errorf("Shell action was just disabled in %s", shellInfo.RCFile)
		}
		lines.RemoveItem(lineNo)
		operation = fmt.Sprintf("Disabled shell action in %s", shellInfo.RCFile)
	} else {
		if err == nil {
			if line == sourceLine {
				return fmt.Errorf("Shell action already enabled in %s at line %d: %s", shellInfo.RCFile, lineNo, line)
			}
			lines.UpdateItem(lineNo, sourceLine)
			operation = fmt.Sprintf("Updated shell action in %s at line %d: %s\nNew line: %s", shellInfo.RCFile, lineNo, line, sourceLine)
		} else {
			lines.AppendItem(sourceLine)
			operation = fmt.Sprintf("Added shell action in %s", shellInfo.RCFile)
		}
	}
	if err = lines.SaveToFile(shellInfo.RCFile); err == nil {
		logger.Info(operation)
	} else {
		logger.Error(fmt.Sprintf("Failed to save %s - %s - %v", shellInfo.RCFile, operation, err))
	}
	return err
}
