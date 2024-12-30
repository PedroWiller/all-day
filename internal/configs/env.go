package configs

import (
	"os"
	"strconv"
)

var (
	Port             int
	ConnectionString string
)

func LoadEnv() error {
	var err error
	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 8080
	}

	ConnectionString = os.Getenv("DB_CONNSTRING")

	return nil
}
