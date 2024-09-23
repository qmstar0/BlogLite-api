package gopubsub

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/charmbracelet/log"
	"github.com/qmstar0/BlogLite-api/pkg/logging"
)

type Logger struct {
	l *log.Logger
}

func NewLogger() Logger {
	return Logger{l: logging.Logger()}
}

func (l Logger) Error(msg string, err error, fields watermill.LogFields) {
	l.fields(fields).Error(msg, "err", err)
}

func (l Logger) Info(msg string, fields watermill.LogFields) {
	l.fields(fields).Info(msg)
}

func (l Logger) Debug(msg string, fields watermill.LogFields) {
	l.fields(fields).Debug(msg)
}

func (l Logger) Trace(msg string, fields watermill.LogFields) {
	l.fields(fields).Info(msg)
}

func (l Logger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	return &Logger{l: l.fields(fields)}
}

func (l Logger) fields(f watermill.LogFields) *log.Logger {
	var fields = make([]any, 0)
	for s, i := range f {
		fields = append(fields, s, i)
	}
	return l.l.With(fields...)
}
