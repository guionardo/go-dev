package ctx

import (
	"context"
	"errors"
	"os"

	"github.com/guionardo/go-dev/pkg/config"
	"github.com/guionardo/go-dev/pkg/consts"
	"github.com/guionardo/go-dev/pkg/logger"
	"github.com/urfave/cli/v2"
)

const (
	configFileKey = consts.FlagConfigFile
	contextKey    = "ctx"
)

type Context struct {
	Config     *config.Config
	ConfigFile string
	LastErr    error
	Debug      bool
}

func SetupContext(c *cli.Context) error {
	c2 := &Context{
		ConfigFile: c.String(configFileKey),
		Debug:      c.Bool(consts.FlagDebug),
	}
	if c2.ConfigFile == "" {
		return errors.New("config_file not found")
	}

	c2.Config, c2.LastErr = config.LoadConfigFile(c2.ConfigFile)
	c.Context = context.WithValue(c.Context, contextKey, c2)
	if c2.Debug {
		logger.SetDebugMode(true)
		logger.Debug("Starting: %v", os.Args)
		logger.Debug("ConfigFile: %v", c2.ConfigFile)
	}
	return nil
}

func GetContext(c *cli.Context) *Context {
	return c.Context.Value(contextKey).(*Context)
}

func ChainedActions(actions ...func(c *cli.Context) error) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		for _, action := range actions {
			err := action(c)
			if err != nil {
				GetContext(c).LastErr = err
				return err
			}
		}
		return nil
	}
}
