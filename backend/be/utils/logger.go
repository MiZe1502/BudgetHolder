package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"runtime"

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
	WriteStackTrace bool         `json:"writeStackTrace"`
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

// Logger is the custom logger realization
type Logger struct {
	instance *logrus.Logger
	config Config
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
func InitLogger(env conf.EnvironmentKey) *Logger {
	conf := conf.ReadLoggerConfig(env)
	parsedConfig := ParseLoggerConfig(conf)

	log := logrus.New()
	logger := Logger{}

	logger.setInstance(log)
	logger.setConfig(parsedConfig)
	logger.setLoggerOutput()
	logger.setLoggingSettings()

	return &logger
}

func (l *Logger) setInstance(logger *logrus.Logger) {
	l.instance = logger
}

func (l *Logger) setConfig(parsedConfig Config) {
	l.config = parsedConfig
}

func (l *Logger) setLoggerOutput() {
	if l.instance == nil {
		return
	}

	if l.config.Output.Type == File {
		fileName := formFileName(l.config)
		fullPath := filepath.Join(l.config.Output.Path, fileName)

		file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			l.instance.SetOutput(file)
		} else {
			l.instance.Info("Failed to log to file, using default stdout")
		}
	} else if l.config.Output.Type == StdOut {
		l.instance.SetOutput(os.Stdout)
	}

}

func (l *Logger) setLoggingSettings() {
	if l.instance == nil {
		return
	}

	l.instance.SetLevel(l.config.LogLevel)
	l.instance.SetReportCaller(l.config.LogMethodName)
}

// GetInstance returns inner logrus logger from custom logger
func (l *Logger) GetInstance() *logrus.Logger {
	return l.instance
}

// Info writes log with info level
func (l *Logger) Info(message string) {
	if l.instance == nil {
		return
	}

	l.writeLog(logrus.InfoLevel, message)
}

// Trace writes log with trace level
func (l *Logger) Trace(message string) {
	if l.instance == nil {
		return
	}

	l.writeLog(logrus.TraceLevel, message)
}

// Error writes log with error level
func (l *Logger) Error(message string) {
	if l.instance == nil {
		return
	}

	l.writeLog(logrus.ErrorLevel, message)
}

// Panic writes log with panic level
func (l *Logger) Panic(message string) {
	if l.instance == nil {
		return
	}

	l.writeLog(logrus.PanicLevel, message)
}

func (l *Logger) addStackTraceToLog() *logrus.Entry {
	b := make([]byte, 2048)
	traceBytes := runtime.Stack(b, false)
	traceString := string(b[:traceBytes])

	return l.instance.WithFields(logrus.Fields{"stackTrace": traceString})
}

func (l *Logger) writeLog(level logrus.Level, message string) {
	if (l.config.WriteStackTrace) {
		entry := l.addStackTraceToLog()
		switch level {
		case logrus.TraceLevel:
			entry.Trace(message)
		case logrus.InfoLevel:
			entry.Info(message)
		case logrus.ErrorLevel:
			entry.Error(message)
		case logrus.PanicLevel:
			entry.Panic(message)
		}
	} else {
		switch level {
		case logrus.TraceLevel:
			l.instance.Trace(message)
		case logrus.InfoLevel:
			l.instance.Info(message)
		case logrus.ErrorLevel:
			l.instance.Error(message)
		case logrus.PanicLevel:
			l.instance.Panic(message)
		}	
	}
}
