package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gobip/internal/config"
	"gobip/internal/handlers"
	"gobip/internal/utils"
)

func main() {
	config.Init("config.json")
	app := fiber.New()
	app.Use(cors.New())
	app.Static("/", "./views/src")
	app.Get("/", handlers.HandleIndex)
	app.Post("/write", handlers.HandleWrite)
	err := utils.ProcessIndexJS("./views/templates/index.js.tmpl")
	if err != nil {
		panic(err)
	}

	err = app.Listen(config.Cfg.Hostname)
	if err != nil {
		panic(err)
	}
}
