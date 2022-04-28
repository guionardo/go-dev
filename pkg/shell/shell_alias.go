package shell

import (
	"fmt"
	"strings"

	"github.com/guionardo/go-dev/pkg/io"
)

func SetAlias(aliasFile string, aliasName string, command string) error {
	lines, err := io.ReadFileToLines(aliasFile)
	if err != nil {
		return err
	}
	index := indexOfAlias(aliasName, lines)
	aliasLine := fmt.Sprintf("%s=%s", aliasName, command)
	if index >= 0 {
		lines[index] = aliasLine
	} else {
		lines = append(lines, aliasLine)
	}
	return io.SaveLinesToFile(aliasFile, lines)

}

func GetAlias(aliasFile string, aliasName string) (command string, err error) {
	lines, err := io.ReadFileToLines(aliasFile)
	command = ""
	if err != nil {
		return "", err
	}
	index := indexOfAlias(aliasName, lines)
	if index < 0 {
		return "", fmt.Errorf("there are no alias '%s' in '%s'", aliasName, aliasFile)
	}
	words := strings.SplitN(lines[index], "=", 2)
	return words[1], nil
}

func RemoveAlias(aliasFile string, aliasName string) error {
	lines, err := io.ReadFileToLines(aliasFile)
	if err != nil {
		return err
	}
	index := indexOfAlias(aliasName, lines)
	if index < 0 {
		return nil
	}
	lines = remove(lines, index)
	return io.SaveLinesToFile(aliasFile, lines)
}

func indexOfAlias(aliasName string, lines []string) int {
	for index, alias := range lines {
		if strings.HasPrefix(alias, aliasName+"=") {
			return index
		}
	}
	return -1
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
