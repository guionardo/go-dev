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

func (c *Config) GetFolders(onlyNotIgnored bool) chan *folders.Folder {
	ch := make(chan *folders.Folder)
	go func() {
		for _, collection := range c.DevFolders {
			for _, folder := range collection.Folders {
				if onlyNotIgnored && folder.Ignore {
					continue
				}
				ch <- folder
			}
		}
		close(ch)
	}()
	return ch
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
	err = fmt.Errorf("folder %s not found", folderName)
	return
}
