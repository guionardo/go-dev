package command

import (
	"fmt"

	"github.com/guionardo/go-dev/cmd/configuration"
	"github.com/urfave/cli/v2"
)

const (
	BaseFolderArg = "basefolder"
	ConfigArg     = "config"
)

func BeforeActionLoadConfiguration(context *cli.Context) error {
	configuration.SetupEnvironmentVars(context.String(BaseFolderArg), context.String(ConfigArg))

	if !configuration.DefaultConfig.TryLoad(configuration.ConfigFileName) {
		return fmt.Errorf("failed to read configuration file %s", configuration.ConfigFileName)
	}
	return nil
}
