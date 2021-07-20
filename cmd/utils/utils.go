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

func FileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	return err == nil && !info.IsDir()
}
func FolderChoice(pastas []string) string {
	switch len(pastas) {
	case 0:
		log.Println("Não foi encontrada pasta")
		return ""
	case 1:
		return pastas[0]
	}

	var choice = -1
	for choice < 0 {
		for i, s := range pastas {
			fmt.Printf("%d - %s\n", i, s)
		}
		fmt.Printf("Escolha uma pasta: ")
		reader := bufio.NewReader(os.Stdin)
		n, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		choice, err = strconv.Atoi(strings.Split(n, "\n")[0])
		if err != nil {
			log.Fatal("Opção inválida")
		}
		if choice >= 0 && choice < len(pastas) {
			return pastas[choice]
		}

		choice = -1
	}

	return ""
}
