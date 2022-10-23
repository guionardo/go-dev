package folders

import (
	"fmt"

	"github.com/guionardo/go-dev/pkg/consts"
)

var AllowedCommandsFunctions = map[string]func(string) string{
	"vscode":           func(p string) string { return fmt.Sprintf("code \"%s\"", p) },
	consts.FlagDisable: func(p string) string { return "" },
	"explorer":         func(p string) string { return fmt.Sprintf("xdg-open \"%s\"", p) },
}

var AllowedCommands = func() []string {
	keys := make([]string, 0, len(AllowedCommandsFunctions))
	for k := range AllowedCommandsFunctions {
		keys = append(keys, k)
	}
	return keys
}()
