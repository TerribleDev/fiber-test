package main

import (
	"os"

	"github.com/TerribleDev/fiber-test/people"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		person := people.Person{
			Name: "John",
			Age:  30,
		}
		return c.JSON(person)
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}

	app.Listen(port)
}
