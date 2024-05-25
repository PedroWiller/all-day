package main

import (
	"fmt"

	"github.com/PedroWiller/all-day/internal/config"
)

func main() {
	var err error

	fmt.Println("Started... ;)")

	err = config.LoadEnv()
	if err != nil {
		fmt.Printf("Error on start config, error %s", err.Error())
		panic(err)
	}
}
