package db

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	conf "../configuration"
)

type DbConfig struct {
	User     string `json:"user"`
	Host     string `json:"host"`
	Database string `json:"database"`
	Password string `json:"password"`
	Port     int    `json:"port"`
}

func FormConnectionString(config DbConfig) string {
	return strings.Join([]string{"host=", config.Host, "port=", strconv.Itoa(config.Port), "user=", config.User, "password=", config.Password, "database=", config.Database}, " ")
}

func ParseDbConfig(config []byte) DbConfig {
	var parsedConfig DbConfig

	err := json.Unmarshal(config, &parsedConfig)

	if err != nil {
		fmt.Println(err)
	}

	return parsedConfig
}

func Connect(env conf.EnvironmentKey) string {
	config := conf.ReadDbConfig(env)

	parsedConfig := ParseDbConfig(config)

	conString := FormConnectionString(parsedConfig)

	return conString
}
