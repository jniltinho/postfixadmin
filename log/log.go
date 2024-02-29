package log

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/sirupsen/logrus"
)

func LOG(format string, a ...any) {
	logg := getLog()
	c := color.New(color.FgYellow).SprintFunc()
	msg := c(fmt.Sprintf(format, a...))
	logg.Warn(msg)
}

func DEBUG(format string, a ...any) {
	logg := getLog()
	c := color.New(color.FgYellow).SprintFunc()
	msg := c(fmt.Sprintf(format, a...))
	logg.Warn(msg)
}

func INFO(format string, a ...any) {
	logg := getLog()
	c := color.New(color.FgGreen).SprintFunc()
	msg := c(fmt.Sprintf(format, a...))
	logg.Info(msg)
}

func ERROR(format string, a ...any) {
	logg := getLog()
	c := color.New(color.FgRed).SprintFunc()
	msg := c(fmt.Sprintf(format, a...))
	logg.Error(msg)
}

func getLog() *logrus.Logger {
	var logger = logrus.New()
	logger.Formatter = new(logrus.TextFormatter)
	logger.Formatter.(*logrus.TextFormatter).TimestampFormat = "2006-01-02 15:04:05"
	logger.Formatter.(*logrus.TextFormatter).FullTimestamp = true
	logger.Formatter.(*logrus.TextFormatter).ForceColors = true
	logger.Formatter.(*logrus.TextFormatter).DisableLevelTruncation = true
	return logger
}
