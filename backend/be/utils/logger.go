package utils

import (
	"github.com/sirupsen/logrus"
)

type Config struct {

	Output OutputConfig `json:"output"`
	LogLevel string `json:"logLevel"`
}

type OutputConfig struct {
	Type     string `json:"type"`
	Path     string `json:"path"`
}