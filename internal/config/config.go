package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	PORT       int64
	ConnString string
)

func LoadEnv() error {
	var err error

	PORT, err = strconv.ParseInt(os.Getenv("PORT"), 10, 32)
	if err != nil {
		PORT = 8080
	}

	ConnString = os.Getenv("DB_CONNSTRING")
	if ConnString == "" {
		return fmt.Errorf("DB_CONNSTRING not found!")
	}

	return nil
}
