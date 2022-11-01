package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	gochoice "github.com/TwiN/go-choice"
	"github.com/guionardo/go-dev/pkg/logger"
)

func PathExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func CreatePath(path string) error {
	if PathExists(path) {
		return nil
	}
	return os.Mkdir(path, 0666)
}

func FileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	return err == nil && !info.IsDir()
}
func FolderChoice(pastas []string, offset int, choiceType string) string {
	switch len(pastas) {
	case 0:
		log.Println("Folder not found")
		return ""
	case 1:
		return pastas[0]
	}
	switch choiceType {
	case "browse":
		return browseChoice(pastas, offset)
	case "index":
		return indexChoice(pastas, offset)
	}
	logger.Error("Invalid choice type: %s - expected 'browse' or 'index'", choiceType)
	return ""
}

func Filter(vs []string, f func(string) bool) []string {
	filtered := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func ReplaceAll(text string, replaces map[string]string) string {
	for _, key := range replaces {
		text = strings.ReplaceAll(text, key, replaces[key])
	}
	return text
}

func browseChoice(pastas []string, offset int) string {
	choices := make([]string, len(pastas))
	for i, s := range pastas {
		choices[i] = s[offset:]
	}
	_, index, err := gochoice.Pick(
		"Pick your folder\n",
		choices,
		gochoice.OptionBackgroundColor(gochoice.Black),
		gochoice.OptionTextColor(gochoice.White),
		gochoice.OptionSelectedTextColor(gochoice.Red),
		gochoice.OptionSelectedTextBold(),
	)
	if err != nil || index < 0 {
		return ""
	}
	return pastas[index]
}

func indexChoice(pastas []string, offset int) string {
	var choice = -1
	for choice < 0 {
		for i, s := range pastas {
			fmt.Printf("%d - %s\n", i, s[offset:])
		}
		fmt.Printf("Choice any folder: ")
		reader := bufio.NewReader(os.Stdin)
		n, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		choice, err = strconv.Atoi(strings.Split(n, "\n")[0])
		if err != nil {
			log.Fatal("Invalid choice")
		}
		if choice >= 0 && choice < len(pastas) {
			return pastas[choice]
		}

		choice = -1
	}

	return ""
}
