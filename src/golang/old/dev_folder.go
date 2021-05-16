package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var devFolder string
var subFolders []string
var devFolderConfig string

const MaximumAge = 86400

func Setup() {
	var home, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	devFolder = path.Join(home, "dev")
	devFolderConfig = path.Join(home, ".dev_folders_go.json")
}

func FindSubFolders() []string {
	devFolderLevel := len(strings.Split(devFolder, "/"))
	var _subFolders []string

	err := filepath.Walk(path.Join(devFolder, "."),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				return nil
			}
			folderLevel := len(strings.Split(path, "/"))
			if folderLevel-devFolderLevel < 3 {
				_subFolders = append(_subFolders, path)
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)

	}
	return _subFolders
}

func SaveConfig() {
	folderJson, _ := json.Marshal(subFolders)
	file, err := os.OpenFile(devFolderConfig, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0655)
	if err != nil {
		log.Fatal(err)
	}
	file.Write(folderJson)
	file.Close()
}

func LoadConfig() {
	file, err := os.Open(devFolderConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	stats, statsErr := file.Stat()
	if statsErr != nil {
		log.Fatal(err)
	}
	var size int64 = stats.Size()
	bytes := make([]byte, size)
	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(bytes, &subFolders); err != nil {
		log.Fatal(err)
	}
}

func NeedUpdateFile(force bool) bool {
	if force {
		return true
	}
	info, err := os.Stat(devFolderConfig)
	if err != nil {
		return true
	}
	return time.Now().Unix()-info.ModTime().Unix() > MaximumAge
}

func FindFolder(words []string) []string {
	//  expression = '('+''.join([f'(.*{w}*?)' for w in where])+')'

	var expression = "("
	for _,s := range words{
		expression=expression+"(.*"+s+"*?)"
	}
	expression+=")"

	var searchPattern = regexp.MustCompile(expression)
	var matches []string
	for _,s := range subFolders{
		if searchPattern.Match([]byte(s)){
			matches=append(matches,s)
		}
	}

	return matches


}

func EscolhePasta(pastas []string) string{
	switch len(pastas) {
	case 0:
		log.Println("Não foi encontrada pasta")
		return ""
	case 1:
		return pastas[0]
	}

	escolha:=-1
	for escolha<0 {
		for i,s :=range(pastas) {
			fmt.Printf("%d - %s\n", i, s)
		}
		fmt.Printf("Escolha uma pasta: ")
		reader :=bufio.NewReader(os.Stdin)
		n,err:=reader.ReadString('\n')
		if err!=nil{
			log.Fatal(err)
		}
		escolha, err=strconv.Atoi(strings.Split(n,"\n")[0])
		if err!=nil{
			log.Fatal(err)
		}
		if escolha>=0 && escolha<len(pastas){
			return pastas[escolha]
		}

		escolha=-1
	}



	return ""
}

func GetFolders() string{
	return strings.Join(subFolders,"\n")
}