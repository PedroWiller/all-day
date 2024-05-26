package helpers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func RecoverMiddleware(app *fiber.App) {
	app.Use(func(c *fiber.Ctx, next fiber.Handler) error {
		if err := next(c); err != nil {
			log.Println("Error:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Something went wrong",
			})
		}
		return nil
	})
}
