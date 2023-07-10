package main

import (
	"github.com/gofiber/fiber/v2"
	"gobip/internal/config"
	"gobip/internal/handlers"
)

func main() {
	app := fiber.New()
	app.Static("/", "./views/src")
	app.Get("/", handlers.HandleIndex)
	app.Post("/write", handlers.HandleWrite)

	config.Init("config.json")

	err := app.Listen(config.Cfg.Hostname)
	if err != nil {
		panic(err)
	}
}
