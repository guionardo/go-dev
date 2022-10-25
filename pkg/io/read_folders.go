package io

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/guionardo/go-dev/pkg/logger"
	"github.com/schollz/progressbar/v3"
)

func ReadFolders(root string, maxSubLevel int) (subFolders []string, err error) {
	intTerm := IsRunningFromInteractiveTerminal() && !logger.IsDebugMode()
	var bar *progressbar.ProgressBar
	logger.Info("Reading folders from %s (depth=%d)", root, maxSubLevel)
	noIntLog := "Reading"
	startTime := time.Now()
	if intTerm {
		bar = progressbar.Default(-1, noIntLog)
	}
	defer func() {
		if bar != nil {
			bar.Finish()
		}
		logger.Info("%s took %v to get %d folders", noIntLog, time.Since(startTime).String(), len(subFolders))
	}()

	subFolders, err = FolderReaderReadDir(root, maxSubLevel,
		func(name string) bool {
			if intTerm {
				bar.Add(1)
			}
			return !strings.HasPrefix(name, ".") && !strings.HasPrefix(name, "_")
		},
		func(name string) {
			logger.Debug("%s", name)
		})

	return

}

func ReadFolders_(root string, maxSubLevel int) ([]string, error) {
	intTerm := IsRunningFromInteractiveTerminal() && !logger.IsDebugMode()
	var bar *progressbar.ProgressBar
	logger.Info("Reading folders from %s (depth=%d)", root, maxSubLevel)
	noIntLog := "Reading"
	startTime := time.Now()

	devFolderLevel := len(strings.Split(root, string(os.PathSeparator)))
	_subFolders := make([]string, 0, 1000)

	if intTerm {
		bar = progressbar.Default(-1, noIntLog)
	}
	defer func() {
		if bar != nil {
			bar.Finish()
		}
		logger.Info("%s took %v to get %d folders", noIntLog, time.Since(startTime).String(), len(_subFolders))
	}()
	err := filepath.WalkDir(path.Join(root, "."),
		func(path string, info os.DirEntry, err error) error {
			if intTerm {
				bar.Add(1)
			}
			if err == nil && info.IsDir() && !DirectoryHasHiddenFolder(path) {
				folderLevel := len(strings.Split(path, "/"))
				if folderLevel-devFolderLevel <= maxSubLevel {
					_subFolders = append(_subFolders, path)
					logger.Debug("%s", path)
				}
			}
			return nil
		})

	return _subFolders, err

}

func DirectoryHasHiddenFolder(directory string) bool {
	folders := strings.Split(directory, string(os.PathSeparator))
	for _, folder := range folders {
		if len(folder) > 0 && (folder[0] == '.' || folder[0] == '_') {
			return true
		}
	}
	return false
}
