package glogger

type fileLogger struct {
	*baseLogger
	fileConfig FileLoggerConfig
}

var _ ILogger = &fileLogger{}

func NewFileLoggerWithConfigFile(path string) ILogger {
	config := loadConfigFile(path)
	return NewFileLoggerWithConfig(*config)
}

func NewFileLoggerWithConfig(cfg FileLoggerConfig) ILogger {
	cfg.ValidateConfig()
	logger := fileLogger{
		fileConfig: cfg,
	}
	return &logger
}

func (l *fileLogger) GetConfig() IConfig {
	return &l.fileConfig
}
