package log

import (
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
)

func LOG(format string, a ...any) {
	c := color.New(color.FgYellow).SprintFunc()
	msg := c(fmt.Sprintf(format, a...))
	level := c("LOG")
	log.Println(level, msg)
}

func LOG2(format string, a ...any) {
	c := color.New(color.FgYellow).SprintFunc()
	msg := c(fmt.Sprintf(format, a...))
	level := c("LOG2")
	r := len(msg) + 26
	color.Red(strings.Repeat("-", r))
	log.Println(level, msg)
	color.Red(strings.Repeat("-", r))
}

func DEBUG(format string, a ...any) {
	c := color.New(color.FgGreen).SprintFunc()
	level := c("DEBUG")
	msg := c(fmt.Sprintf(format, a...))
	log.Println(level, msg)
}

func INFO(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	level := "INFO"
	log.Println(level, msg)
}

func ERROR(format string, a ...any) {
	c := color.New(color.FgRed).SprintFunc()
	msg := c(fmt.Sprintf(format, a...))
	level := "ERROR"
	log.Println(level, msg)
}
