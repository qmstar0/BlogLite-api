package logging

import (
	"github.com/charmbracelet/log"
	"os"
)

var Logger = log.NewWithOptions(os.Stdout, log.Options{
	TimeFormat:      "2006/01/02 15:04:05",
	ReportTimestamp: true,
})

func WithPrefix(prefix string) *log.Logger {
	return Logger.WithPrefix(prefix)
}
