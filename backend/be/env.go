package main

import (
	"log"

	"github.com/jackc/pgx/pgxpool"
)

//Env stores current global app variables
type Env struct {
	db     *pgxpool.Pool
	logger *log.Logger
}
