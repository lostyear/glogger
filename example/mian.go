package main

import (
	"fmt"

	"github.com/lostyear/glogger"
)

func main() {
	dl := glogger.GetDefaultLogger()
	fmt.Println(dl.GetConfig())
	glogger.Debugf("This is a debug log")
	glogger.Infof("This is an info log")
	glogger.Warnf("This is a watn log")
	glogger.Errorf("This is an error log")
	// glogger.Fatalf("This is a fatal log")
	logger := glogger.NewFileLoggerWithConfigFile("config.toml")
	fmt.Println(logger.GetConfig())
	logger.Debugf("This is a debug log")
	logger.Infof("This is an info log")
	// logger.Warnf("This is a watn log")
	// logger.Errorf("This is an error log")
}
