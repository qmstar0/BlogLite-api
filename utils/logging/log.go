package logging

import (
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"blog/infra/config"
)

type Log struct {
	*logrus.Entry
}

var (
	Logger *logrus.Logger
)

func init() {
	Logger = logrus.New()
	// 设置日志格式
	{
		Logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006/01/02 15:04:05",
		})
	}
	//设置日志LEVEL
	{
		if config.Conf.Logger.Level == "" {
			panic("Log output level is not configured; see **config.go**; from **logging/log.go**")
		}
		l, err := logrus.ParseLevel(config.Conf.Logger.Level)
		if err != nil {
			Logger.SetLevel(4)
		} else {
			Logger.SetLevel(l)
		}
	}
	// 设置日志输出
	{
		if config.Conf.Logger.OutputPath == "" {
			panic("The log output directory is not configured; see **config.go**; from **logging/log.go**")
		}
		logDir := config.Conf.Logger.OutputPath
		if err := os.MkdirAll(logDir, os.FileMode(0666)); err != nil {
			panic("An httpError occurred while creating the log folder; from **logging/log.go**")
		}
		logPath := path.Join(logDir, "log.log")
		l := genRotateOutput(logPath)
		Logger.SetOutput(l)
	}
}

func New(name string) *Log {
	return &Log{Logger.WithField("logger", name)}
}

func genRotateOutput(logPath string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    10, // 10 MB
		MaxBackups: 10,
		MaxAge:     30, // 30 days
		Compress:   true,
		LocalTime:  true,
	}
}
