package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello World."})
	})

	log.Fatal(app.Listen(":4000"))
}
