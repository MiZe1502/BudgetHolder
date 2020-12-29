package utils

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"
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
	Output          OutputConfig `json:"output"`
	LogLevel        logrus.Level `json:"logLevel"`
	LogMethodName   bool         `json:"logMethodName"`
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
	instance  *logrus.Logger
	config    Config
	traceUUID string
	curDay    int
}

// ParseLoggerConfig parses json config for logger
func ParseLoggerConfig(config []byte) (Config, error) {
	var parsedConfig Config

	err := json.Unmarshal(config, &parsedConfig)

	return parsedConfig, err
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
func InitLogger(env conf.EnvironmentKey) (*Logger, error) {
	conf, err := conf.ReadLoggerConfig(env)
	if err != nil {
		return nil, err
	}

	parsedConfig, err := ParseLoggerConfig(conf)
	if err != nil {
		return nil, err
	}

	log := logrus.New()
	logger := Logger{}

	logger.setInstance(log)
	logger.setConfig(parsedConfig)
	logger.updateCurLoggerDay()

	err = logger.setLoggerOutput()
	if err != nil {
		return nil, err
	}

	err = logger.setLoggingSettings()

	return &logger, err
}

func (l *Logger) updateCurLoggerDay() {
	l.curDay = time.Now().Day()
}

func (l *Logger) setInstance(logger *logrus.Logger) {
	l.instance = logger
}

func (l *Logger) setConfig(parsedConfig Config) {
	l.config = parsedConfig
}

func (l *Logger) setLoggerOutput() error {

	if l.instance == nil {
		return errors.New("no logger instance")
	}

	if l.config.Output.Type == File {
		fileName := formFileName(l.config)
		fullPath := filepath.Join(l.config.Output.Path, fileName)

		file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			l.instance.SetOutput(file)
		} else {
			return errors.New("Failed to log to file, using default stdout")
		}
	} else if l.config.Output.Type == StdOut {
		l.instance.SetOutput(os.Stdout)
	}

	return nil
}

func (l *Logger) isDayChanged() bool {
	curTime := time.Now()
	curDay := curTime.Day()

	return curDay != l.curDay
}

func (l *Logger) updateLogFile() {
	isNewDay := l.isDayChanged()

	if isNewDay {
		l.updateCurLoggerDay()
		l.setLoggerOutput() // initialize new log file
	}
}

// ProcessLogFileUpdate check current date and creates new log file for new day
func (l *Logger) ProcessLogFileUpdate() {
	for {
		l.updateLogFile()
	}
}

func (l *Logger) setLoggingSettings() error {
	if l.instance == nil {
		return errors.New("no logger instance")
	}

	l.instance.SetLevel(l.config.LogLevel)
	l.instance.SetReportCaller(l.config.LogMethodName)
	return nil
}

// GetInstance returns inner logrus logger from custom logger
func (l *Logger) GetInstance() *logrus.Logger {
	return l.instance
}

// SetTraceUUID injects traceUUID to trace all requests route
func (l *Logger) SetTraceUUID(traceUUID string) {
	l.traceUUID = traceUUID
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
	if &l.instance == nil {
		return
	}

	l.writeLog(logrus.ErrorLevel, message)
}

// Fatal writes log with fatal level
func (l *Logger) Fatal(message string) {
	if l.instance == nil {
		return
	}

	l.writeLog(logrus.FatalLevel, message)
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

	return l.instance.WithFields(logrus.Fields{"stackTrace": traceString, "traceUUID": l.traceUUID})
}

//TODO: too dirty, find the way how to append this field not in two different places
func (l *Logger) addTraceUUIDToLog() *logrus.Entry {
	return l.instance.WithFields(logrus.Fields{"traceUUID": l.traceUUID})
}

func (l *Logger) writeLog(level logrus.Level, message string) {
	if l.config.WriteStackTrace {
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
		entry := l.addTraceUUIDToLog()
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
	}
}
