package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Status(500).Send("Welcome to my broken app v4.0.0!!")
	})
	app.Get("/healthz", func(c *fiber.Ctx) {
		c.Status(500)
	})

	app.Listen(3004)
}
