package source

import (
	"fmt"
	"os"
)

const (
	DBHost     = "DB_HOST"
	DBPort     = "DB_PORT"
	DBUser     = "DB_USER"
	DBPass     = "DB_PASS"
	ServerPort = "SERVER_PORT"
)

func CheckEnvVars() {
	envVars := []string{
		DBHost,
		DBPort,
		DBUser,
		DBPass,
		ServerPort,
	}

	for _, v := range envVars {
		if os.Getenv(v) == "" {
			panic(fmt.Sprintf("env variable %s must be defined", v))
		}
	}
}
