package shell

import (
	"os"
	"strings"
)

func lineHasSearchPath(line string, searchPath string) bool {
	line = strings.TrimSpace(line)
	if !(strings.HasPrefix(line, "export PATH=") || strings.HasPrefix(line, "PATH=")) {
		return false
	}

	paths := strings.Split(strings.Split(line, "=")[1], string(os.PathListSeparator))
	for _, path := range paths {
		if path == searchPath {
			return true
		}
	}
	return false
}
