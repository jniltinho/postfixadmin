package log

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/fatih/color"
)

func LOG(format string, a ...any) {
	green := color.New(color.FgGreen).SprintFunc()
	msg := green(fmt.Sprintf(format, a...))
	slog.Info(msg)
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
