package main

import (
	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		response := Message{Text: "GoLang Test"}
		return c.JSON(response)
	})

	app.Listen(":3000")

}
