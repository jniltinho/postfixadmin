package util

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/fatih/color"
)

func LOG(format string, a ...any) {
	color.Set(color.FgGreen)
	slog.Info(fmt.Sprintf(format, a...))
	color.Unset()
}

func LOG2(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	r := len(msg) + 26
	color.Red(strings.Repeat("-", r))
	color.Set(color.FgGreen)
	slog.Info(msg)
	color.Unset()
	color.Red(strings.Repeat("-", r))
}
