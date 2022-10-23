package consts

import (
	"time"
)

const (
	AuthorName  = "Guionardo Furlan"
	AuthorEmail = "guionardo@gmail.com"
	AppName     = "go-dev"
)

var (
	CompileTimeString = "1970-01-01T00:00:00"
	CompileTime       time.Time
	Version           = "0.0.0"
	BuildRunner       = "unknown"
)

func SetupBuildInfo(build_date string, build_runner string, release string) {
	CompileTimeString = build_date
	CompileTime, _ = time.Parse("2006-01-02T15:04:05", CompileTimeString)
	Version = release
	BuildRunner = build_runner
}
