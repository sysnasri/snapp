package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello World!")
	})
	app.Get("/healthz", func(c *fiber.Ctx) {
		c.Status(200)
	})

	app.Listen(3001)
}
