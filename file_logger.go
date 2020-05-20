package glogger

type FileLogger struct {
	*baseLogger
	fileConfig FileLoggerConfig
}

func NewFileLoggerWithConfigFile(path string) *FileLogger {
	config := loadConfigFile(path)
	return NewFileLoggerWithConfig(*config)
}

func NewFileLoggerWithConfig(cfg FileLoggerConfig) *FileLogger {
	logger := FileLogger{
		fileConfig: cfg,
	}
	return &logger
}

func (l *FileLogger) GetConfig() FileLoggerConfig {
	return l.fileConfig
}
