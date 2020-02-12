package main

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send(`{"message": "Hello World!"}`)
	})

	// GET /john
	app.Get("/:name", func(c *fiber.Ctx) {
		fmt.Printf("Hello %s!", c.Params("name"))
		// => Hello john!
	})

	// GET /john/18
	app.Get("/:name/:age?", func(c *fiber.Ctx) {
		fmt.Printf("Name: %s, Age: %s", c.Params("name"), c.Params("age"))
		// => Name: john, Age: 18
	})

	// GET /api/register
	app.Get("/api/*", func(c *fiber.Ctx) {
		fmt.Printf("/api/%s", c.Params("*"))
		// => /api/register
	})

	app.Listen(8080)
}
