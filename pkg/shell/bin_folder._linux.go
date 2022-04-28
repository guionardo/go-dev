package shell

import (
	"fmt"
	"log"
	"os"
	"path"

	"strings"

	"github.com/guionardo/go-dev/cmd/debug"
)

func GetBinFolder() string {
	binFolder := path.Join(homePath, "bin")
	if stat, err := os.Stat(binFolder); err != nil || !stat.IsDir() {
		err = os.Mkdir(binFolder, 0776)
		if err != nil {
			log.Fatalf("Could not create folder %s - %v", binFolder, err)
		}
		debug.Debug(fmt.Sprintf("Folder created %s", binFolder))
	}
	return binFolder
}

func findWriteableBinFolders() (writeablePaths []string) {
	writeablePaths = make([]string, 0)
	for _, searchPath := range GetSearchPaths() {
		if !strings.HasSuffix(searchPath, "bin") {
			continue
		}
		tmp, err := os.CreateTemp(searchPath, "findFirstBin*")
		if err != nil {
			continue
		}
		os.Remove(tmp.Name())
		writeablePaths = append(writeablePaths, searchPath)
	}
	return
}

func binFolderCandidates() (binFolders []string) {
	binFolders = make([]string, 0)

}
