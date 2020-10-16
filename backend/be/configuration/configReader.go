package config

import (
	"fmt"
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
	dbConfigPath = "./configuration/configs/db.json"
	loggerConfigPath = "./configuration/configs/logger.json"
)

// ReadDbConfig reads config for pg db as byte array
func ReadDbConfig(env EnvironmentKey) []byte {
	config, err := os.Open(dbConfigPath)

	if err != nil {
		fmt.Println(err)
	}

	defer config.Close()

	byteValue, _ := ioutil.ReadAll(config)

	return byteValue
}

// ReadLoggerConfig reads config for logger as byte array
func ReadLoggerConfig(env EnvironmentKey) []byte {
	config, err := os.Open(loggerConfigPath)

	if err != nil {
		fmt.Println(err)
	}

	defer config.Close()

	byteValue, _ := ioutil.ReadAll(config)

	return byteValue
}
