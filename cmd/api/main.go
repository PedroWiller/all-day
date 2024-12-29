package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/PedroWiller/all-day/cmd/api/routers"
	"github.com/PedroWiller/all-day/configs"
	"github.com/PedroWiller/all-day/pkg/helpers"
)

func main() {
	var err error

	fmt.Println("Started... ;)")

	err = configs.LoadEnv()
	if err != nil {
		fmt.Printf("Error on start config, error %s", err.Error())
		panic(err)
	}

	db, err := configs.ConnectDB()
	if err != nil {
		fmt.Printf("Error on start config, error %s", err.Error())
		panic(err)
	}

	app := fiber.New()

	helpers.RecoverMiddleware(app)
	injectRoute(app, db)
}

func injectRoute(app *fiber.App, db *gorm.DB) {
	routers.UserRouter(app, db)
}
