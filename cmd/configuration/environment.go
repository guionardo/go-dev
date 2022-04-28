package configuration

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/guionardo/go-dev/cmd/debug"
	"github.com/guionardo/go-dev/cmd/utils"
)

const MaximumAge = 86400
const MaximumSubLevel = 4

var (
	HomePath                string
	MaxFolderLevel          int
	DevFolder               string
	ConfigFileName          string
	DefaultDevFolder        string
	DefaultFolderConfigFile string
	DefaultOutputFile       string
)

func CurrentDir() string {
	dir, err := os.Getwd()
	if err == nil {
		return dir
	}
	return "."
}

func SetupBaseEnvironment() error {
	utils.SetupLogging()
	var err error
	HomePath, err = os.UserHomeDir()
	if err == nil {
		DefaultFolderConfigFile = path.Join(HomePath, ".dev_folders_go.yaml")
		DefaultDevFolder = path.Join(HomePath, "dev")
		DefaultOutputFile = path.Join(HomePath, ".dev_folders_go.out")
		debug.Debug(fmt.Sprintf("Defaults: FolderConfigFile=%s DevFolder=%s OutputFile=%s",
			DefaultFolderConfigFile, DefaultDevFolder, DefaultOutputFile))
	}
	utils.SetOutput(DefaultOutputFile)

	return err
}

func SetupEnvironmentVars(devFolder string, configurationFile string) {
	if len(devFolder) == 0 {
		devFolder = DefaultDevFolder
	}
	DevFolder = devFolder
	if len(configurationFile) == 0 {
		configurationFile = DefaultFolderConfigFile
	}
	ConfigFileName = configurationFile
}

func NeedUpdateConfigFile(filename string, force bool) bool {
	if force {
		debug.Debug("Forced update config file")
		return true
	}
	info, err := os.Stat(filename)
	if err != nil {
		return true
	}
	fileAge := time.Now().Unix() - info.ModTime().Unix()
	fileAgeDur, _ := time.ParseDuration(fmt.Sprintf("%ds", fileAge))
	maxDur, _ := time.ParseDuration(fmt.Sprintf("%ds", MaximumAge))
	debug.Debug(fmt.Sprintf("Configuration File: %s (%v | %v | %v)", filename, info.ModTime(), fileAge, fileAgeDur))

	if fileAgeDur > maxDur {
		log.Printf("Reloading paths after %v (%v)", fileAgeDur, maxDur)
		return true
	}
	debug.Debug(fmt.Sprintf("File age (%v) less than (%v). No need to update", fileAge, maxDur))
	return false
}
