package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx"
)

type Config struct {
	User     string `json:user`
	Host     string `json:host`
	Database string `json:database`
	Password string `json:password`
	Port     int    `json:port`
}

func main() {
	config, err := os.Open("./configs/conf.json")

	if err != nil {
		fmt.Println(err)
	}

	defer config.Close()

	byteValue, _ := ioutil.ReadAll(config)

	var parsedConfig Config

	json.Unmarshal(byteValue, &parsedConfig)

	fmt.Println(parsedConfig)

	dbURL := strings.Join([]string{"host=", parsedConfig.Host, "port=", strconv.Itoa(parsedConfig.Port), "user=", parsedConfig.User, "password=", parsedConfig.Password, "database=", parsedConfig.Database}, " ")

	fmt.Println(dbURL)

	conn, err := pgx.Connect(context.Background(), dbURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())
}
