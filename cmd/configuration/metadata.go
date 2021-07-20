package configuration

import (
	_ "embed"
	"strings"
	"time"
)

//go:embed metadata.txt
var metadata string

var MetaData MetadataType

type MetadataType struct {
	AppName     string
	Version     string
	BuildDate   string
	CompileTime time.Time
}

func LoadMetaData() MetadataType {
	ct, _ := time.Parse("2006-01-02T15:04:05", getValue("build_date"))
	MetaData = MetadataType{
		AppName:     getValue("name"),
		BuildDate:   getValue("build_date"),
		Version:     getValue("version"),
		CompileTime: ct,
	}
	return MetaData
}

func getValue(key string) string {
	lines := strings.Split(metadata, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, key+"=") {
			words := strings.Split(strings.ReplaceAll(line, "\n", ""), "=")
			return words[1]
		}
	}
	return ""
}
