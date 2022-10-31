package config

import (
	"fmt"
	"os"
	"path"

	"github.com/guionardo/go-dev/pkg/folders"
	"github.com/guionardo/go-dev/pkg/logger"
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
	logger.Debug("Saving config to %s - error=%v", filename, err)
	return err
}

func (c *Config) Find(folderName string) (coll *folders.FolderCollection, folder *folders.Folder, err error) {
	for _, collection := range c.DevFolders {
		if folder, err = collection.Get(folderName); folder != nil {
			coll = collection
			return
		}
	}
	err = fmt.Errorf("Folder %s not found", folderName)
	return
}