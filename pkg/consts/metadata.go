package consts

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"time"
)

const (
	AuthorName  = "Guionardo Furlan"
	AuthorEmail = "guionardo@gmail.com"
	AppName     = "go-dev"
)

type MetadataStruct struct {
	BuildTime   time.Time `json:"build_time"`
	Version     string    `json:"version"`
	BuildRunner string    `json:"build_runner"`
}

//go:embed metadata.json
var MetadataJson []byte

var (
	Metadata MetadataStruct
)

func init() {
	if err := json.Unmarshal(MetadataJson, &Metadata); err != nil {
		panic(fmt.Errorf("Error reading metadata.json: %v", err))
	}
}
