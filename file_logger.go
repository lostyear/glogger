package glogger

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

	logger := fileLogger{
		fileConfig: cfg,
		baseLogger: newBaseLoggerWithConfig(cfg.LoggerConfig),
	}
	return &logger
}

func (l *fileLogger) GetConfig() IConfig {
	return &l.fileConfig
}
