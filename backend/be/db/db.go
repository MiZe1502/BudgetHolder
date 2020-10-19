package db

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	conf "../configuration"

	"github.com/jackc/pgx/pgxpool"
)

// Config represents struct for db config
type Config struct {
	User     string `json:"user"`
	Host     string `json:"host"`
	Database string `json:"database"`
	Password string `json:"password"`
	Port     int    `json:"port"`
}

// FormConnectionString creates configuration string to
// connect to pg db
func FormConnectionString(config Config) string {
	return strings.Join([]string{"host=", config.Host, "port=", strconv.Itoa(config.Port), "user=", config.User, "password=", config.Password, "database=", config.Database}, " ")
}

// ParseDbConfig parses json config for pg db
func ParseDbConfig(config []byte) (Config, error) {
	var parsedConfig Config

	err := json.Unmarshal(config, &parsedConfig)

	return parsedConfig, err
}

// Connect creates global connection to pg db
func Connect(env conf.EnvironmentKey) (*pgxpool.Pool, error) {
	config, err := conf.ReadDbConfig(env)
	if err != nil {
		return nil, err
	}

	parsedConfig, err := ParseDbConfig(config)
	if err != nil {
		return nil, err
	}

	conString := FormConnectionString(parsedConfig)

	conn, err := pgxpool.Connect(context.Background(), conString)

	return conn, err
}
