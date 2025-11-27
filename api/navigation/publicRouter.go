package router

import (
	"github.com/gofiber/fiber/v2"
)

func PublicRouter(app fiber.Router) {

	app.Get("/privacy", func(c *fiber.Ctx) error {
		return c.SendFile("./static/privacy.html")
	})

	app.Static("/api/v1/pdf", "./pdf")

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(map[string]any{
			"error":   "404 Not Found",
			"success": false,
		})
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

}
