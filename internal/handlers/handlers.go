package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gobip/internal/utils"
)

func HandleIndex(c *fiber.Ctx) error {
	return c.SendFile("./views/src/index.html")
}

func HandleWrite(c *fiber.Ctx) (err error) {
	phrase := string(c.Body())
	return utils.AppendText("./phrases.txt", phrase+"\n")
}
