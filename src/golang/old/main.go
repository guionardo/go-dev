package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)
func parseArgs() (string, []string){
	args:=os.Args[1:]

	var params []string

	if len(args)==0 {
		return "help", params
	}
	return args[0],args[1:]
}

func ShowList(){
	fmt.Printf("Pastas:\n%s",GetFolders())
}

func ShowHelp(){
log.Println("go [termos] [de] [pesquisa]")
log.Println("list")
log.Println("update")
}

func GoFolder(args []string){
	var match = FindFolder(args)
	var pasta = EscolhePasta(match)
	if len(pasta) == 0 {
		fmt.Print("#")
	} else {
		fmt.Print(pasta)
	}
}

func UpdateFolders(){
	subFolders = FindSubFolders()
	SaveConfig()
	log.Printf("%s\n\nPastas atualizadas",GetFolders())
}

func main() {
	Setup()
	if NeedUpdateFile(false) {
		subFolders = FindSubFolders()
		SaveConfig()
	} else {
		LoadConfig()
	}
	comando, args := parseArgs()
	switch strings.ToLower(comando) {

	case "list":
		ShowList()
	case "update":
		UpdateFolders()
	case "go":
		GoFolder(args)
	default:
		ShowHelp()
	}




}
