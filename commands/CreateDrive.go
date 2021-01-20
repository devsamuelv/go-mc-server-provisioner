package commands

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

// CreateDrive Create Drives Post Request
func CreateDrive(app *fiber.App) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	app.Post("/-server/CreateDrive", func(c *fiber.Ctx) error {
		return ctx.Err()
	})
}
