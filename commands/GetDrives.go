package commands

import (
	"main/util"

	"github.com/gofiber/fiber/v2"
)

// GetDrives Get Drives Post Request
func GetDrives(app *fiber.App) {
	app.Post(util.Prefix("GetDrives"), func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"status": 404,
		})
	})
}
