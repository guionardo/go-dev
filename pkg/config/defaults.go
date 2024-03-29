package config

import (
	"os"
	"path"

	"github.com/guionardo/go-dev/pkg/consts"
	"github.com/guionardo/go-dev/pkg/folders"
)

var (
	DefaultConfigFile string
	DefaultOutputFile string
	home              string
)

func init() {
	var err error
	home, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	DefaultConfigFile = path.Join(home, ".config", consts.AppName, consts.AppName+".yaml")
	DefaultOutputFile = path.Join(home, ".config", consts.AppName, consts.AppName+".sh")
}

func GetDefaultConfig() *Config {
	return &Config{
		DevFolders: make(map[string]*folders.FolderCollection, 0),
	}
}
