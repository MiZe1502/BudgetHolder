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

// ReadDbConfig reads config for pg db as byte array
func ReadDbConfig(env EnvironmentKey) []byte {
	config, err := os.Open("./configuration/configs/conf.json")

	if err != nil {
		fmt.Println(err)
	}

	defer config.Close()

	byteValue, _ := ioutil.ReadAll(config)

	return byteValue
}
