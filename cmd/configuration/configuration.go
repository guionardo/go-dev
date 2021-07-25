package configuration

import (
	"errors"
	"fmt"
	"github.com/guionardo/go-dev/cmd/utils"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type PathSetup struct {
	Ignore  bool   `json:"ignore"`
	Command string `json:"cmd"`
	Path    string `json:"path"`
}

type Paths map[string]PathSetup

func (path PathSetup) ToString() string {
	if path.Ignore {
		return fmt.Sprintf("Path: %s [IGNORED]", path.Path)
	}
	return fmt.Sprintf("Path: %s [Command=%s]", path.Path, path.Command)
}

func (pc *Paths) Set(path PathSetup) error {
	if !utils.PathExists(path.Path) {
		return errors.New("Path not found: " + path.Path)
	}
	var index = ""
	for i, p := range *pc {
		if p.Path == path.Path {
			index = i
			break
		}
	}
	if len(index) == 0 {
		index = strconv.Itoa(len(*pc) + 1)
	}

	(*pc)[index] = path
	return nil
}

func (pc *Paths) Get(path string) (PathSetup, error) {
	for key, p := range *pc {
		if p.Path == path {
			if !utils.PathExists(p.Path) {
				delete(*pc, key)
				break
			}
			return p, nil
		}
	}
	return PathSetup{}, errors.New("Dev path not found: " + path)
}

func (pc *Paths) FolderList() []string {
	var list []string
	for _, p := range *pc {
		if !p.Ignore {
			list = append(list, p.Path)
		}
	}
	sort.Strings(list)
	return list
}
func matchPath(path string, words []string) bool {
	lastIndex := -1
	for _, s := range words {
		if len(s) == 0 {
			continue
		}
		i := strings.Index(path, s)
		if i <= lastIndex {
			return false
		}
		lastIndex = i
	}
	return true
}
func (pc *Paths) FindFolder(words []string) []PathSetup {
	fmt.Printf("Finding %s\n", words)

	var matches []PathSetup
	for _, s := range *pc {
		p := s.Path[len(DevFolder):]
		if !s.Ignore && matchPath(p, words) {
			matches = append(matches, s)
		}
	}

	return matches
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

func (pc *Paths) ReadFolders(devFolder string, maxSubLevel int) error {
	devFolder, err := filepath.Abs(devFolder)
	if err != nil {
		return err
	}
	log.Printf("Reading folders: %s\n", devFolder)
	devFolderLevel := len(strings.Split(devFolder, string(os.PathSeparator)))
	var _subFolders []string
	err = filepath.WalkDir(path.Join(devFolder, "."),
		func(path string, info os.DirEntry, err error) error {
			if err == nil && info.IsDir() && !DirectoryHasHiddenFolder(path) {
				folderLevel := len(strings.Split(path, "/"))
				if folderLevel-devFolderLevel < maxSubLevel {
					_subFolders = append(_subFolders, path)
				}
			}
			return nil
		})
	if err != nil {
		return err
	}
	for _, folder := range _subFolders {
		_, err := pc.Get(folder)
		if err != nil {
			if err = pc.Set(PathSetup{Path: folder}); err != nil {
				log.Printf("Failed to add folder %s - %v", folder, err)
			}
		}
	}
	log.Printf("%d folders readen", len(_subFolders))
	return nil
}
