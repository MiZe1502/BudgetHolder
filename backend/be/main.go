package main

import (
	db "./db"

	"context"
	"fmt"
	"os"
)

func main() {
	conn := db.Connect("Dev")

	defer conn.Close(context.Background())

	var greeting string
	err := conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
