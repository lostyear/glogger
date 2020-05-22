package glogger

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

type IConfig interface {
	GetConfig() IConfig
	validate()
}

type FileLoggerConfigFile struct {
	LoggerConfig

	Path            string
	FilePrefix      string
	FileSuffix      string
	SplitLevelFiles bool
	RotationMinutes int
	MaxAgeHours     int
	MaxCount        uint
}

var _ IConfig = &FileLoggerConfigFile{}

func (fcfg *FileLoggerConfigFile) convertConfigFile() *FileLoggerConfig {
	cfg := FileLoggerConfig{
		BaseFileLogConfig: BaseFileLogConfig{
			LoggerConfig: fcfg.LoggerConfig,
			rotationTime: time.Duration(fcfg.RotationMinutes) * time.Minute,
			maxAge:       time.Duration(fcfg.MaxAgeHours) * time.Hour,
			maxCount:     fcfg.MaxCount,
		},

		SplitLevelFiles: fcfg.SplitLevelFiles,
		FilePath:        fcfg.Path,
		FilePrefix:      fcfg.FilePrefix,
		FileSuffix:      fcfg.FileSuffix,
	}

	return &cfg
}

type FileLoggerConfig struct {
	BaseFileLogConfig

	FilePath        string
	FilePrefix      string
	FileSuffix      string
	SplitLevelFiles bool
}

var _ IConfig = &FileLoggerConfig{}

//TODO: file logger validate
func (fcfg *FileLoggerConfig) newFileLevelLoggerConfig(level string) FileLevelLoggerConfig {
	var filename string
	if fcfg.SplitLevelFiles {
		filename = fmt.Sprintf("%s%s%s%s", fcfg.FilePath, fcfg.FilePrefix, level, fcfg.FileSuffix)
	} else {
		filename = fmt.Sprintf("%s%s%s%s", fcfg.FilePath, fcfg.FilePrefix, "", fcfg.FileSuffix)
	}

	cfg := FileLevelLoggerConfig{
		BaseFileLogConfig: fcfg.BaseFileLogConfig,
		filename:          filename,
		//TODO: setup linked filename
		linkedFilename: fmt.Sprintf(""),
	}

	return cfg
}

type FileLevelLoggerConfig struct {
	BaseFileLogConfig

	filename       string
	linkedFilename string
}

var _ IConfig = &FileLevelLoggerConfig{}

type BaseFileLogConfig struct {
	LoggerConfig
	rotationTime time.Duration
	maxAge       time.Duration
	maxCount     uint
}

var _ IConfig = &BaseFileLogConfig{}

func (cfg *FileLoggerConfig) GetConfig() IConfig {
	return cfg
}

func loadConfigFile(path string) *FileLoggerConfig {
	var conf FileLoggerConfigFile
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		log.Panicf("decode config file failed! Error: %s", err)
	}
	return conf.convertConfigFile()
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
