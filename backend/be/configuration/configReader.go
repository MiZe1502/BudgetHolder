package config

import (
	"io/ioutil"
	"os"
)

// EnvironmentKey is a type for env keys enum
type EnvironmentKey string

const (
	Dev  EnvironmentKey = "Dev"
	Test EnvironmentKey = "Test"
	Prod EnvironmentKey = "Prod"
)

const (
	dbConfigPath     = "./configuration/configs/db.json"
	loggerConfigPath = "./configuration/configs/logger.json"
	tokenConfigPath = "./configuration/configs/token.json"
)

// ReadDbConfig reads config for pg db as byte array
func ReadDbConfig(env EnvironmentKey) ([]byte, error) {
	config, err := os.Open(dbConfigPath)

	if err != nil {
		return nil, err
	}

	defer config.Close()

	byteValue, err := ioutil.ReadAll(config)

	return byteValue, err
}

// ReadLoggerConfig reads config for logger as byte array
func ReadLoggerConfig(env EnvironmentKey) ([]byte, error) {
	config, err := os.Open(loggerConfigPath)

	if err != nil {
		return nil, err
	}

	defer config.Close()

	byteValue, err := ioutil.ReadAll(config)

	return byteValue, err
}

// ReadTokenConfig reads config for token generation as byte array
func ReadTokenConfig(env EnvironmentKey) ([]byte, error) {
	config, err := os.Open(tokenConfigPath)

	if err != nil {
		return nil, err
	}

	defer config.Close()

	byteValue, err := ioutil.ReadAll(config)

	return byteValue, err
}