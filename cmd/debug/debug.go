package debug

import (
	"fmt"
	"log"
)

var (
	Enabled  = false
	debugMsg = make([]string, 0)
)

func Debug(message string) {
	if len(message)>0 {
		debugMsg = append(debugMsg, fmt.Sprintf("[DEBUG] %s", message))
	}
	if !Enabled {
		return
	}

	for _, msg := range debugMsg {
		log.Print(msg)
	}
	debugMsg = make([]string, 0)
}
