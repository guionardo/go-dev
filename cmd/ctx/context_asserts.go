package ctx

import (
	"fmt"

	"github.com/guionardo/go-dev/cmd/update"
	"github.com/guionardo/go-dev/pkg/config"
	"github.com/guionardo/go-dev/pkg/logger"
	"github.com/urfave/cli/v2"
)

func AssertConfigExists(c *cli.Context) error {
	if GetContext(c).Config == nil {
		return fmt.Errorf("Config file %s  not found", GetContext(c).ConfigFile)
	}
	return nil
}

func AssertAtLeastDefault(c *cli.Context) error {
	c2 := GetContext(c)
	if c2.Config != nil {
		return nil
	}
	c2.Config = config.GetDefaultConfig()
	return c2.Config.Save(c2.ConfigFile)
}

func AssertSaveIfNotError(c *cli.Context) error {
	c2 := GetContext(c)
	if c2.LastErr == nil {
		return c2.Config.Save(c2.ConfigFile)
	}
	return c2.LastErr
}

func AssertAutoSync(c *cli.Context) error {
	c2 := GetContext(c)
	if c2.Config.AutoSync.ShouldRun() {
		for _, folder := range c2.Config.DevFolders {
			if err := folder.Sync(); err != nil {
				c2.LastErr = err
				return err
			}
		}
		logger.Info("AutoSync completed")
		c2.Config.AutoSync.Run()
		return c2.Config.Save(c2.ConfigFile)
	}
	return nil
}

func AssertAutoUpdate(c *cli.Context) error {
	c2 := GetContext(c)
	if c2.Config.AutoUpdate.ShouldRun() {
		if err := update.RunGitUpdate(); err != nil {
			c2.LastErr = err
			return err
		}
		c2.Config.AutoUpdate.Run()
		return c2.Config.Save(c2.ConfigFile)
	}
	return nil
}
