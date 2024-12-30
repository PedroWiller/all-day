package main

import "github.com/PedroWiller/all-day/configs"

func main() {
	var err error

	if err = configs.LoadEnv(); err != nil {
		panic(err)
	}
}
