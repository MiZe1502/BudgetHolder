package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	conf "../configuration"
	logrus "github.com/sirupsen/logrus"
)

// LoggerOutputType is used to make enum of logger output types
type LoggerOutputType string

// File is used to log into file
// StdOut is default
const (
	File   LoggerOutputType = "file"
	StdOut LoggerOutputType = "stdout"
)

// Config stores logger configuration
type Config struct {
	Output   OutputConfig `json:"output"`
	LogLevel logrus.Level `json:"logLevel"`
}

// OutputConfig represents struct for logger output configuration
type OutputConfig struct {
	Type            LoggerOutputType `json:"type"`
	Path            string           `json:"path"`
	DefaultFileName string           `json:"defaultFileName"`
	UseDefaultFile bool `json:"useDefaultFile"`
}

// ParseLoggerConfig parses json config for logger
func ParseLoggerConfig(config []byte) Config {
	var parsedConfig Config

	err := json.Unmarshal(config, &parsedConfig)

	if err != nil {
		fmt.Println(err)
	}

	return parsedConfig
}

// InitLogger creates and configures logger instance
func InitLogger(env conf.EnvironmentKey) *logrus.Logger {
	conf := conf.ReadLoggerConfig(env)

	parsedConfig := ParseLoggerConfig(conf)

	var log = logrus.New()

	fileName := parsedConfig.Output.DefaultFileName
	if !parsedConfig.Output.UseDefaultFile {
		fileName = time.Now().UTC().Format("2006-01-02 MST") + ".log"
	}
	if parsedConfig.Output.Type == File {
		fullPath := filepath.Join(parsedConfig.Output.Path, fileName)

		fmt.Println(fullPath)

		file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(file)
		} else {
			log.Info("Failed to log to file, using default stdout")
		}
	} else if parsedConfig.Output.Type == StdOut {
		log.SetOutput(os.Stdout)
	}

	log.SetLevel(parsedConfig.LogLevel)

	return log
}
