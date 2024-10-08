package logging

import (
	"github.com/charmbracelet/log"
	"os"
)

var logger *log.Logger

func Init(debug bool) {
	var level log.Level
	if debug {
		level = log.DebugLevel
	} else {
		level = log.WarnLevel
	}
	logger = log.NewWithOptions(os.Stdout, log.Options{
		TimeFormat:      "2006/01/02 15:04:05",
		ReportTimestamp: true,
		Level:           level,
	})
}

func Logger() *log.Logger {
	return logger
}

func WithPrefix(prefix string) *log.Logger {
	return logger.WithPrefix(prefix)
}
