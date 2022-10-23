package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
func FolderChoice(pastas []string, offset int) string {
	switch len(pastas) {
	case 0:
		log.Println("Folder not found")
		return ""
	case 1:
		return pastas[0]
	}

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
