package io

import (
	"os"

	"github.com/mattn/go-isatty"
)

func IsRunningFromInteractiveTerminal() bool {
	return isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
}
