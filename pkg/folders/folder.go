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

func (f *Folder) String() string {
	ignored := ""
	if f.Ignore {
		ignored = " [IGNORED]"		
	}
	command := ""
	if f.Command != "" {
		command = fmt.Sprintf(" [CMD: %s]", f.Command)
	}
	return fmt.Sprintf("Path: %s%s%s", f.Path, ignored, command)
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

func (f *Folder) SetCommand(cmd string) {
	if command, ok := AllowedCommandsFunctions[cmd]; ok {
		f.Command = command(f.Path)
		return
	}
	f.Command = cmd
}
