package shell

import (
	"fmt"
	"os"
	"strings"
)

func FindByLine(filename string, finder func(line string) bool) (lineNumber int, lineContent string, err error) {
	var content []byte
	content, err = os.ReadFile(filename)
	if err != nil {
		return
	}
	lines := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")
	for i, line := range lines {
		if finder(line) {
			lineNumber = i
			lineContent = line
			return
		}
	}
	err = fmt.Errorf("Line not found in file '%s'", filename)
	return
}
