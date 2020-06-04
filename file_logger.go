package glogger

import (
	"fmt"

	rlogs "github.com/lestrrat-go/file-rotatelogs"
)

type FileLogger interface {
	Logger

	//todo: file logger funcs
}

type fileLogger struct {
	*baseLogger
	fileConfig FileLoggerConfig
}

var _ Logger = &fileLogger{}

func NewFileLoggerWithConfigFile(path string) (FileLogger, error) {
	if config, err := loadConfigFile(path); err != nil {
		return nil, err
	} else {
		return NewFileLoggerWithConfig(*config)
	}
}

func NewFileLoggerWithConfig(cfg FileLoggerConfig) (FileLogger, error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	l := fileLogger{
		fileConfig: cfg,
		baseLogger: newBaseLoggerWithConfig(cfg.LoggerConfig),
	}

	for level, num := range lLevel {
		if num < lLevel[cfg.Level] {
			l.lLoggers[level] = defaultEmptyLogger
		} else {
			if log, err := newFileLevelLoggerWithConfig(cfg.newFileLevelLoggerConfig(level)); err != nil {
				return nil, err
			} else {
				l.lLoggers[level] = log
			}
		}
	}
	return &l, nil
}

func (l *fileLogger) GetConfig() IConfig {
	return &l.fileConfig
}

type fileLevelLogger struct {
	*baseLevelLogger

	fileWriter *rlogs.RotateLogs
}

func newFileLevelLoggerWithConfig(cfg FileLevelLoggerConfig) (*fileLevelLogger, error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	w, err := rlogs.New(
		cfg.filename,
		rlogs.WithLinkName(cfg.linkedFilename),
		rlogs.WithRotationTime(cfg.rotationTime),
		rlogs.WithMaxAge(cfg.maxAge),
		rlogs.WithRotationCount(cfg.maxCount),
	)
	if err != nil {
		return nil, fmt.Errorf("open log file error! ", err)
	}

	logger := newBaseLogger(w, cfg.Level, cfg.Flags)
	return &fileLevelLogger{
		baseLevelLogger: logger,
	}, nil
}
