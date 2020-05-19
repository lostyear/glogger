package glogger

import (
	"github.com/BurntSushi/toml"
)

type LoggerConfig struct {
	Path            string
	LogFilePrefix   string
	LogFileSuffix   string
	logLevel        string
	SplitLevelFiles bool
	RotationMinutes int
	MaxAgeHours     int
}

// var config *LoggerConfig

func loadConfigFile(path string) *LoggerConfig {
	var conf LoggerConfig
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		panic("decode config file failed!")
	}
	return &conf
}
