package glogger

import (
	"log"
	"time"

	rlogs "github.com/lestrrat-go/file-rotatelogs"
)

type FileLogger interface {
	ILogger
}

type fileLogger struct {
	*baseLogger
	fileConfig FileLoggerConfig
}

var _ ILogger = &fileLogger{}

func NewFileLoggerWithConfigFile(path string) FileLogger {
	config := loadConfigFile(path)
	return NewFileLoggerWithConfig(*config)
}

func NewFileLoggerWithConfig(cfg FileLoggerConfig) FileLogger {
	cfg.validate()

	l := fileLogger{
		fileConfig: cfg,
		baseLogger: newBaseLoggerWithConfig(cfg.LoggerConfig),
	}

	for level, num := range lLevel {
		if num < lLevel[cfg.Level] {
			l.lLoggers[level] = defaultEmptyLogger
		} else {
			l.lLoggers[level] = newFileLevelLoggerWithConfig(FileLoggerConfig{})
		}
	}
	return &l
}

func (l *fileLogger) GetConfig() IConfig {
	return &l.fileConfig
}

type fileLevelLogger struct {
	*baseLevelLogger
}

func newFileLevelLoggerWithConfig(cfg FileLoggerConfig) *fileLevelLogger {
	cfg.validate()

	w, err := rlogs.New(
		"filename",
		rlogs.WithRotationTime(time.Second),
		rlogs.WithMaxAge(time.Hour),
	)
	if err != nil {
		log.Panic("open log file error! ", err)
	}

	logger := newBaseLogger(w, cfg.Level, cfg.Flags)
	return &fileLevelLogger{
		baseLevelLogger: logger,
	}
}
