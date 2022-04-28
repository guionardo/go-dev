package shell

import (
	"os"
	"strings"
)

func GetSearchPaths() []string {
	searchPath := os.Getenv("PATH")
	searchPaths := strings.Split(searchPath, string(os.PathListSeparator))
	return searchPaths
}

func AddSearchPath(rcFile string, searchPath string) error {

}
