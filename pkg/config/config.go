package config

import (
	"os"
	"path"

	"github.com/guionardo/go-dev/pkg/folders"
	"gopkg.in/yaml.v2"
)

type Config struct {
	DevFolders map[string]*folders.FolderCollection `yaml:"dev_folders"`
	AutoSync   IntervalRunner                       `yaml:"auto_sync"`
	AutoUpdate IntervalRunner                       `yaml:"auto_update"`
}

func LoadConfigFile(filename string) (config *Config, err error) {
	var content []byte
	content, err = os.ReadFile(filename)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(content, &config)
	if err == nil {
		for _, collection := range config.DevFolders {
			collection.FixPathsLoad()
		}
	}
	return
}

func (c *Config) Save(filename string) error {
	content, err := yaml.Marshal(c)
	if err == nil {
		configDir := path.Dir(filename)
		var stat os.FileInfo
		if stat, err = os.Stat(configDir); os.IsNotExist(err) || !stat.IsDir() {
			err = os.MkdirAll(configDir, 0755)
		}
		if err == nil {
			err = os.WriteFile(filename, content, 0644)
		}
	}
	return err
}
