package db

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"context"
	"os"

	conf "../configuration"

	"github.com/jackc/pgx"
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
func ParseDbConfig(config []byte) Config {
	var parsedConfig Config

	err := json.Unmarshal(config, &parsedConfig)

	if err != nil {
		fmt.Println(err)
	}

	return parsedConfig
}

// Connect creates global connection to pg db
func Connect(env conf.EnvironmentKey) *pgx.Conn {
	config := conf.ReadDbConfig(env)

	parsedConfig := ParseDbConfig(config)

	conString := FormConnectionString(parsedConfig)

	conn, err := pgx.Connect(context.Background(), conString)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}
