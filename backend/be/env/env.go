package env

import (
	utils "../utils"
	"github.com/jackc/pgx/pgxpool"
)

//Env stores current global app variables
type Env struct {
	Db     *pgxpool.Pool
	Logger *utils.Logger
	Token  *utils.TokenGenerator
}

// EnvironmentKey is a type for env keys enum
type EnvironmentKey string

const (
	Dev  EnvironmentKey = "Dev"
	Test EnvironmentKey = "Test"
	Prod EnvironmentKey = "Prod"
)
