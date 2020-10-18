package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
    "runtime/debug" 

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
	Output        OutputConfig `json:"output"`
	LogLevel      logrus.Level `json:"logLevel"`
	LogMethodName bool         `json:"logMethodName"`
}

// OutputConfig represents struct for logger output configuration
type OutputConfig struct {
	Type            LoggerOutputType `json:"type"`
	Path            string           `json:"path"`
	DefaultFileName string           `json:"defaultFileName"`
	UseDefaultFile  bool             `json:"useDefaultFile"`
	FileDateFormat  string           `json:"fileDateFormat"`
	FileExt         string           `json:"fileExt"`
}

type Logger struct {
	instance *logrus.Logger
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

func formFileName(parsedConfig Config) string {
	fileName := ""
	if !parsedConfig.Output.UseDefaultFile {
		fileName = time.Now().UTC().Format(parsedConfig.Output.FileDateFormat) + parsedConfig.Output.FileExt
	} else {
		fileName = parsedConfig.Output.DefaultFileName + parsedConfig.Output.FileExt
	}

	return fileName
}

// InitLogger creates and configures logger instance
func InitLogger(env conf.EnvironmentKey) *logrus.Logger {
	conf := conf.ReadLoggerConfig(env)

	parsedConfig := ParseLoggerConfig(conf)

	var log = logrus.New()

	if parsedConfig.Output.Type == File {
		fileName := formFileName(parsedConfig)
		fullPath := filepath.Join(parsedConfig.Output.Path, fileName)

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

	log.SetReportCaller(parsedConfig.LogMethodName)

	return log
}


func writeLog(level logrus.Level) {
	switch level {
    case logrus.TraceLevel:
        fmt.Println("one")
    case logrus.InfoLevel:
        fmt.Println("two")
    case logrus.ErrorLevel:
		fmt.Println("three")
	case logrus.PanicLevel:
        fmt.Println("three")
    }
}