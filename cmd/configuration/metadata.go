package configuration

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/guionardo/go-dev/pkg/logger"
)

//go:embed metadata.txt
var metadata string

var MetaData MetadataType

type MetadataType struct {
	BuilderInfo string
	CompileTime time.Time
}

func init() {
	LoadMetaData()
}

func LoadMetaData() MetadataType {
	ct, _ := time.Parse("2006-01-02T15:04:05", getValue("build_date"))
	MetaData = MetadataType{
		BuilderInfo: getValue("builder_info"),
		CompileTime: ct,
	}
	logger.Debug(fmt.Sprintf("Metadata: %v", MetaData))
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
