package glogger

import (
	"log"
	"strings"

	"github.com/BurntSushi/toml"
)

type IConfig interface {
	GetConfig() IConfig
	validate()
}

type FileLoggerConfig struct {
	LoggerConfig

	Path            string
	LogFilePrefix   string
	LogFileSuffix   string
	SplitLevelFiles bool
	RotationMinutes int
	MaxAgeHours     int
}

var _ IConfig = &FileLoggerConfig{}

//TODO: file logger validate

func (cfg *FileLoggerConfig) GetConfig() IConfig {
	return cfg
}

func loadConfigFile(path string) *FileLoggerConfig {
	var conf FileLoggerConfig
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		log.Panicf("decode config file failed! Error: %s", err)
	}
	return &conf
}

type LoggerConfig struct {
	Flags int
	Level string
}

var _ IConfig = &LoggerConfig{}

func (cfg *LoggerConfig) GetConfig() IConfig {
	return cfg
}

func (cfg *LoggerConfig) validate() {
	if cfg.Flags == 0 {
		cfg.Flags = defaultLogFlags
	}

	level := strings.ToUpper(cfg.Level)
	if _, ok := lLevel[level]; !ok {
		level = defaultLogLevel
	}

	cfg.Level = level
}
