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

func ReadFolders(root string, maxSubLevel int) ([]string, error) {
	intTerm := IsRunningFromInteractiveTerminal() && !logger.IsDebugMode()
	var bar *progressbar.ProgressBar
	noIntLog := "Reading folders " + root
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

// func (pc *Paths) ReadFolders(devFolder string, maxSubLevel int) error {
// 	devFolder, err := filepath.Abs(devFolder)
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf("Reading folders: %s\n", devFolder)
// 	devFolderLevel := len(strings.Split(devFolder, string(os.PathSeparator)))
// 	var _subFolders []string
// 	err = filepath.WalkDir(path.Join(devFolder, "."),
// 		func(path string, info os.DirEntry, err error) error {
// 			if err == nil && info.IsDir() && !DirectoryHasHiddenFolder(path) {
// 				folderLevel := len(strings.Split(path, "/"))
// 				if folderLevel-devFolderLevel < maxSubLevel {
// 					_subFolders = append(_subFolders, path)
// 				}
// 			}
// 			return nil
// 		})
// 	if err != nil {
// 		return err
// 	}
// 	for _, folder := range _subFolders {
// 		_, err := pc.Get(folder)
// 		if err != nil {
// 			if err = pc.Set(PathSetup{Path: folder}); err != nil {
// 				log.Printf("Failed to add folder %s - %v", folder, err)
// 			}
// 		}
// 	}
// 	log.Printf("%d folders readen", len(_subFolders))
// 	return nil
// }
