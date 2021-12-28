package main

import (
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

	app.Listen(":3000")
}
