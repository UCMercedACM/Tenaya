package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	// USE: /api --> assigns all headers for routes under api route
	app.Use("/api", func(c *fiber.Ctx) {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Headers", "X-Requested-With")
		c.Set("Content-Type", "application/json")
		c.Next()
	})

	// GET: /events --> all calendar events
	app.Get("/api/events", func(c *fiber.Ctx) {
		c.Status(200).Send("All Event Data")
	})

	// GET: /events/:type --> return only a subgroup of all events
	app.Get("/api/events/:type", func(c *fiber.Ctx) {
		c.Status(200).Send("Event Data for: " + c.Params("type"))
	})

	// GET: /event/:id --> returns specific event
	app.Get("/api/event/:id", func(c *fiber.Ctx) {
		c.Status(200).Send("Displaying Event Data for: " + c.Params("id"))
	})

	// POST: /events --> create multiple new events
	app.Post("/events", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully created new events")
	})

	// POST: /event --> create a single new event
	app.Post("/event", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully created a new event")
	})

	// PATCH: /events --> update the data of all events at once
	app.Patch("/events", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully updated all events")
	})

	// PATCH: /events/:type --> update all the events of a single type
	app.Patch("/events/:type", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully updated a all events with type: " + c.Params("type"))
	})

	// PATCH: /event/:id --> update a single event
	app.Patch("/event/:id", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully updated event: " + c.Params("id"))
	})

	// DELETE: /events --> completely delete all events (for testing only)[should not be in production]
	app.Delete("/events", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully deleted all events")
	})

	// DELETE: /events/:type --> deletes all events under a specific type
	app.Delete("/events/:type", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully deleted all events of type: " + c.Params(("type")))
	})

	// DELETE: /event/:id --> delete a specific event
	app.Delete("/event/:id", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully deleted event with id: " + c.Params("id"))
	})

	// * --> handles all unknown routes
	app.Get("*", func(c *fiber.Ctx) {
		c.Status(200).Send("Unknown Request")
	})

	app.Listen(8080)
}
