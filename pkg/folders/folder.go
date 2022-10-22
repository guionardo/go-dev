package folders

import (
	"fmt"
	"strings"
)

type Folder struct {
	Path             string `yaml:"-"`
	Ignore           bool   `yaml:"ignore"`
	IgnoreSubFolders bool   `yaml:"ignore_sub"`
	Command          string `yaml:"cmd"`
}

func (f *Folder) ToString() string {
	if f.Ignore {
		return fmt.Sprintf("Path: %s [IGNORED]", f.Path)
	}
	return fmt.Sprintf("Path: %s [Command=%s]", f.Path, f.Command)
}

func (f *Folder) Match(words []string) bool {
	if len(f.Path) == 0 {
		return false
	}
	path := f.Path
	for i := len(words) - 1; i >= 0; i-- {
		word := words[i]
		p := strings.Index(path, word)
		if p == -1 {
			return false
		}
		path = path[:p]
	}
	return true
}
