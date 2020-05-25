package glogger

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/lestrrat-go/strftime"
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

func loadConfigFile(path string) *FileLoggerConfig {
	var conf FileLoggerConfigFile
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		//TODO: panic in lib is not good
		log.Panicf("decode config file failed! Error: %s", err)
	}
	return conf.convertConfigFile()
}

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

func (cfg *FileLoggerConfig) validate() {
	if _, err := strftime.New(cfg.FilePath); err != nil {
		//TODO: panic in lib is not good
		panic("invalid path string")
	}
	if _, err := strftime.New(cfg.FilePrefix); err != nil {
		//TODO: panic in lib is not good
		panic("invalid prefix string")
	}
	if _, err := strftime.New(cfg.FileSuffix); err != nil {
		//TODO: panic in lib is not good
		panic("invalid suffix string")
	}
	// TODO: validate permision

}

func (cfg *FileLoggerConfig) GetConfig() IConfig {
	return cfg
}

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

func (cfg *FileLevelLoggerConfig) validate() {
	if _, err := strftime.New(cfg.filename); err != nil {
		//TODO: panic in lib is not good
		panic("invalid filename string")
	}
	//TODO: valid linked  file
}

func (cfg *FileLevelLoggerConfig) GetConfig() IConfig {
	return cfg
}

type BaseFileLogConfig struct {
	LoggerConfig
	rotationTime time.Duration
	maxAge       time.Duration
	maxCount     uint
}

var _ IConfig = &BaseFileLogConfig{}

func (cfg *BaseFileLogConfig) validate() {
	if cfg.rotationTime < 0 {
		//TODO: panic in lib is not good
		panic("invalid rotation time")
	}
	if cfg.maxAge > 0 && cfg.maxCount > 0 {
		//TODO: panic in lib is not good
		panic("only one of maxAge  and maxCount can set")
	}

	if cfg.maxAge <= 0 {
		cfg.maxAge = 0
	}
	if cfg.maxCount <= 0 {
		cfg.maxCount = 0
	}
}

func (cfg *BaseFileLogConfig) GetConfig() IConfig {
	return cfg
}

type LoggerConfig struct {
	Flags int
	Level string
}

var _ IConfig = &LoggerConfig{}

func (cfg *LoggerConfig) GetConfig() IConfig {
	return cfg
}

//optimize: less call
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
