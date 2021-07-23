package configuration

import (
	"github.com/guionardo/go-dev/cmd/utils"
	"log"
	"os"
	"path"
	"time"
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
		DefaultFolderConfigFile = path.Join(HomePath, ".dev_folders_go.json")
		DefaultDevFolder = path.Join(HomePath, "dev")
		DefaultOutputFile = path.Join(HomePath, ".dev_folders_go.out")
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
		return true
	}
	info, err := os.Stat(filename)
	if err != nil {
		return true
	}
	if time.Now().Unix()-info.ModTime().Unix() > MaximumAge {
		log.Printf("Reloading paths after %v", time.Duration(MaximumAge))
		return true
	}
	return false
}
