package glogger

import "log"

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
		}
	}
	return &l
}

func (l *fileLogger) GetConfig() IConfig {
	return &l.fileConfig
}

type fileLevelLogger struct {
	*log.Logger
}
