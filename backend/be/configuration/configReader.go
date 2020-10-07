package config

import (
	"fmt"
	"io/ioutil"
	"os"
)

type EnvironmentKey string

const (
	Dev  EnvironmentKey = "Dev"
	Test EnvironmentKey = "Test"
	Prod EnvironmentKey = "Prod"
)

func ReadDbConfig(env EnvironmentKey) []byte {
	config, err := os.Open("./configuration/configs/conf.json")

	if err != nil {
		fmt.Println(err)
	}

	defer config.Close()

	byteValue, _ := ioutil.ReadAll(config)

	return byteValue
}
