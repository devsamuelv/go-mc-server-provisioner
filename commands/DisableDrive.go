package commands

import (
	"main/util"

	"github.com/gofiber/fiber/v2"
)

// DisableDrive Disable Drive Command For Post Request
func DisableDrive(app *fiber.App) {
	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()

	app.Post(util.Prefix("DisableDrive"), func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"message": "Not Finished",
			"code":    404,
		})
	})
}
