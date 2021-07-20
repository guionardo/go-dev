package command

import (
	"bitbucket.org/kardianos/osext"
	_ "embed"
	"fmt"
	"github.com/guionardo/go-dev/cmd/configuration"
	"github.com/guionardo/go-dev/cmd/utils"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

//go:embed dev.sh
var devSh string

var (
	installRemove   bool
	devBaseFolder   string
	maximumSubLevel int
	InstallCmd      = &cli.Command{
		Name:  "install",
		Usage: "Install go-dev",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "uninstall",
				Usage:       "Remove installation",
				Destination: &installRemove,
			},
			&cli.PathFlag{
				Name:        "basefolder",
				Usage:       "Development base folder",
				Value:       configuration.DefaultDevFolder,
				Destination: &devBaseFolder,
			},
			&cli.IntFlag{
				Name:        "max-path-level",
				Usage:       "Maximum level of paths",
				Value:       configuration.MaximumSubLevel,
				Destination: &configuration.MaxFolderLevel,
			},
		},
		Action: InstallAction,
		Before: BeforeInstallAction,
	}
)

func InstallAction(context *cli.Context) error {
	newConfig := &configuration.ConfigFileType{
		DevFolder:         configuration.DevFolder,
		Paths:             make(configuration.Paths),
		ConfigurationFile: configuration.ConfigurationFileName,
	}
	var err error
	if err = newConfig.Paths.ReadFolders(configuration.DevFolder, configuration.MaxFolderLevel); err == nil {
		if err = newConfig.Save(); err == nil {
			fmt.Printf("Configuration file saved @ %s (base folder = %s)\n", newConfig.ConfigurationFile, newConfig.DevFolder)
			goDevScript, err := installScript()
			if err == nil {
				if err = installAlias(goDevScript); err == nil {
					fmt.Println("Alias setup is done")
				}
			}
		}
	}

	return err
}

func installScript() (string, error) {
	filename, err := osext.Executable()
	var scriptFile string
	if err == nil {
		devSh = strings.ReplaceAll(devSh, "{GO_DEV}", filename)
		scriptFile = path.Join(filepath.Dir(filename), "go-dev.sh")
		err = os.WriteFile(scriptFile, []byte(devSh), 0655)
	}
	return scriptFile, err
}

func installAlias(goDevShFile string) error {
	aliasDevLine := fmt.Sprintf("alias dev='source %s'", goDevShFile)
	var err error
	bashRcAliasesFile := path.Join(configuration.HomePath, ".bash_aliases")
	var lines []string
	if utils.FileExists(bashRcAliasesFile) {
		aliases, err := os.ReadFile(bashRcAliasesFile)
		if err == nil {
			lines = strings.Split(string(aliases), "\n")
			aliasIndex := -1
			for index, line := range lines {
				if strings.HasPrefix(line, "alias dev=") {
					aliasIndex = index
					break
				}
			}
			if aliasIndex > -1 {
				lines[aliasIndex] = aliasDevLine
			} else {
				lines = append(lines, aliasDevLine)
			}
		}
	} else {
		lines = append(lines, aliasDevLine)
	}
	content := strings.Join(lines, "\n")
	err = os.WriteFile(bashRcAliasesFile, []byte(content), 0655)

	return err
}

func BeforeInstallAction(context *cli.Context) error {
	configuration.SetupEnvironmentVars(context.String("basefolder"), context.String("config"))

	if !configuration.DefaultConfig.TryLoad(configuration.ConfigurationFileName) {
		log.Printf("New configuration file: %s", configuration.ConfigurationFileName)
	}
	return nil
}
