package configuration

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/guionardo/go-dev/cmd/debug"
)

//go:embed metadata.txt
var metadata string

var MetaData MetadataType

const (
	AuthorName  = "Guionardo Furlan"
	AuthorEmail = "guionardo@gmail.com"
)

type MetadataType struct {
	AppName     string
	Version     string
	BuildDate   string
	BuilderInfo string
	CompileTime time.Time
	AuthorName  string
	AuthorEmail string
}

func (metadata *MetadataType) ToString() string {
	return fmt.Sprintf("%s v%s %s", metadata.AppName, metadata.Version, metadata.BuildDate)
}

func init() {
	LoadMetaData()
}

func LoadMetaData() MetadataType {
	ct, _ := time.Parse("2006-01-02T15:04:05", getValue("build_date"))
	MetaData = MetadataType{
		AppName:     getValue("name"),
		BuildDate:   getValue("build_date"),
		Version:     getValue("version"),
		BuilderInfo: getValue("builder_info"),
		CompileTime: ct,
		AuthorName:  AuthorName,
		AuthorEmail: AuthorEmail,
	}
	debug.Debug(fmt.Sprintf("Metadata: %v", MetaData))
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
