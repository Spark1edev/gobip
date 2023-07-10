package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	err := app.Listen("45.12.239.218:8080")
	if err != nil {
		panic(err)
	}
}
