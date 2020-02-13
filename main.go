package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	/**

	/api --> assigns all headers for routes under api route

	GET
	events --> all calendar events
	events/:type --> return only a subgroup of all events
	event/:id --> returns specific event

	POST
	events --> create multiple new events
	event --> create a single new event

	PATCH
	events --> update the data of all events at once
	events/:type --> update all the events of a single type
	event/:id --> update a single event

	DELETE
	events --> completely delete all events (for testing only)[should not be in production]
	events/:type --> deletes all events under a specific type
	event/:id --> delete a specific event

	* --> handles all unknown routes

	*/

	// Match all routes starting with /api
	app.Use("/api", func(c *fiber.Ctx) {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Headers", "X-Requested-With")
		c.Set("Content-Type", "application/json")
		c.Next()
	})

	// Optional param
	// Test in your terminal with ==> curl -X POST http://localhost:8080/api/register -d "username=john&password=doe"
	app.Post("/api/register", func(c *fiber.Ctx) {
		username := c.Body("username")
		password := c.Body("password")
		c.Status(200).Send("username: " + username + ", password: " + password)
		// ..
	})

	app.Listen(8080)
}
