package commands

import (
	"github.com/gofiber/fiber/v2"
	"context"
)

func CreateDrive(app *fiber.App) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	app.Post("/-server/CreateDrive", func(c *fiber.Ctx) error {
		
	})

}