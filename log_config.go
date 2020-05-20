package glogger

import (
	"log"

	"github.com/BurntSushi/toml"
)

type FileLoggerConfig struct {
	LoggerConfig

	Path            string
	LogFilePrefix   string
	LogFileSuffix   string
	SplitLevelFiles bool
	RotationMinutes int
	MaxAgeHours     int
}

// var config *LoggerConfig

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
