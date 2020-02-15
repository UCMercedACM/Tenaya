package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

func main() {
	// Create new fiber server
	app := fiber.New()

	// Loads environment variables and handles errors
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database Credentials
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	//! Debug only
	fmt.Printf("%s: %s, %s\n", dbName, dbUsername, dbPassword)

	// USE: /api --> assigns all headers for routes under api route
	app.Use("/api", func(c *fiber.Ctx) {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Headers", "X-Requested-With")
		c.Set("Content-Type", "application/json")
		c.Next()
	})

	/**
	 * @api {GET} /api/events
	 * @apiDescription Returns all calendar events
	 * @apiVersion 1.0.0
	 * @apiName Get Events
	 * @apiGroup Event(s)
	 * @apiPermission public
	 *
	 * @apiSuccess {Object[]} Returns an object array of all event information.
	 *
	 * @apiError (Unauthorized 401)  Unauthorized  Only authenticated users can access the data
	 * @apiError (Forbidden 403)     Forbidden     Only admins can access the data
	 */
	app.Get("/api/events", func(c *fiber.Ctx) {
		c.Status(200).Send("All Event Data")
	})

	/**
	 * @api {GET} /api/events/:type
	 * @apiDescription Returns only a subgroup of all events
	 * @apiVersion 1.0.0
	 * @apiName Get Events Based on Type
	 * @apiGroup Event(s)
	 * @apiPermission public
	 *
	 * @apiSuccess {Object[]} Returns an object array of all event information.
	 *
	 * @apiError (Unauthorized 401)  Unauthorized  Only authenticated users can access the data
	 * @apiError (Forbidden 403)     Forbidden     Only admins can access the data
	 */
	app.Get("/api/events/:type", func(c *fiber.Ctx) {
		c.Status(200).Send("Event Data for: " + c.Params("type"))
	})

	/**
	 * @api {GET} /api/event/:id
	 * @apiDescription Returns a single event
	 * @apiVersion 1.0.0
	 * @apiName Get Events Based on ID
	 * @apiGroup Event(s)
	 * @apiPermission public
	 *
	 * @apiSuccess {Object[]} Returns an object array of only one event.
	 *
	 * @apiError (Unauthorized 401)  Unauthorized  Only authenticated users can access the data
	 * @apiError (Forbidden 403)     Forbidden     Only admins can access the data
	 */
	app.Get("/api/event/:id", func(c *fiber.Ctx) {
		c.Status(200).Send("Displaying Event Data for: " + c.Params("id"))
	})

	/**
	 * @api {POST} /api/events
	 * @apiDescription Creates multiple new events
	 * @apiVersion 1.0.0
	 * @apiName Creates Events
	 * @apiGroup Event(s)
	 * @apiPermission public
	 *
	 * @apiSuccess {string} Returns a success message logging that the new events were created successfully.
	 *
	 * @apiError (Unauthorized 401)  Unauthorized  Only authenticated users can access the data
	 * @apiError (Forbidden 403)     Forbidden     Only admins can access the data
	 */
	app.Post("/events", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully created new events")
	})

	/**
	 * @api {POST} /api/event
	 * @apiDescription Create a new event
	 * @apiVersion 1.0.0
	 * @apiName Create Event
	 * @apiGroup Event(s)
	 * @apiPermission public
	 *
	 * @apiSuccess {string} Returns a success message logging that the new event were created successfully.
	 *
	 * @apiError (Unauthorized 401)  Unauthorized  Only authenticated users can access the data
	 * @apiError (Forbidden 403)     Forbidden     Only admins can access the data
	 */
	app.Post("/event", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully created a new event")
	})

	/**
	 * @api {PATCH} /api/events/:type
	 * @apiDescription Updates all events in the database
	 * @apiVersion 1.0.0
	 * @apiName Update Events
	 * @apiGroup Event(s)
	 * @apiPermission public
	 *
	 * @apiSuccess {string} Returns a success message logging that the events were updated successfully.
	 *
	 * @apiError (Unauthorized 401)  Unauthorized  Only authenticated users can access the data
	 * @apiError (Forbidden 403)     Forbidden     Only admins can access the data
	 */
	app.Patch("/events", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully updated all events")
	})

	/**
	 * @api {PATCH} /api/events/:type
	 * @apiDescription Updates all events of a specific type in the database
	 * @apiVersion 1.0.0
	 * @apiName Update Events By Type
	 * @apiGroup Event(s)
	 * @apiPermission public
	 *
	 * @apiSuccess {string} Returns a success message logging that the events were updated successfully.
	 *
	 * @apiError (Unauthorized 401)  Unauthorized  Only authenticated users can access the data
	 * @apiError (Forbidden 403)     Forbidden     Only admins can access the data
	 */
	app.Patch("/events/:type", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully updated a all events with type: " + c.Params("type"))
	})

	/**
	 * @api {PATCH} /api/event/:id
	 * @apiDescription Updates an Event
	 * @apiVersion 1.0.0
	 * @apiName Update Event
	 * @apiGroup Event(s)
	 * @apiPermission public
	 *
	 * @apiSuccess {string} Returns a success message logging that the event were updated successfully.
	 *
	 * @apiError (Unauthorized 401)  Unauthorized  Only authenticated users can access the data
	 * @apiError (Forbidden 403)     Forbidden     Only admins can access the data
	 */
	app.Patch("/event/:id", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully updated event: " + c.Params("id"))
	})

	// DELETE: /events --> completely delete all events (for testing only)[should not be in production]
	/**
	 * @api {DELETE} /api/events
	 * @apiDescription Returns only a subgroup of all events
	 * @apiVersion 1.0.0
	 * @apiName Get Events Based on Type
	 * @apiGroup Event(s)
	 * @apiPermission public
	 *
	 * @apiSuccess {Object[]} Returns an object array of all event information.
	 *
	 * @apiError (Unauthorized 401)  Unauthorized  Only authenticated users can access the data
	 * @apiError (Forbidden 403)     Forbidden     Only admins can access the data
	 */
	app.Delete("/events", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully deleted all events")
	})

	// DELETE: /events/:type --> deletes all events under a specific type
	/**
	 * @api {DELETE} /api/events/:type
	 * @apiDescription Returns only a subgroup of all events
	 * @apiVersion 1.0.0
	 * @apiName Get Events Based on Type
	 * @apiGroup Event(s)
	 * @apiPermission public
	 *
	 * @apiSuccess {Object[]} Returns an object array of all event information.
	 *
	 * @apiError (Unauthorized 401)  Unauthorized  Only authenticated users can access the data
	 * @apiError (Forbidden 403)     Forbidden     Only admins can access the data
	 */
	app.Delete("/events/:type", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully deleted all events of type: " + c.Params(("type")))
	})

	// DELETE: /event/:id --> delete a specific event
	/**
	 * @api {DELETE} /api/event/:id
	 * @apiDescription Deletes all information on a single event
	 * @apiVersion 1.0.0
	 * @apiName Deletes an Event
	 * @apiGroup Event(s)
	 * @apiPermission public
	 *
	 * @apiSuccess {string} Returns a success message logging that the event was deleted successfully.
	 *
	 * @apiError (Unauthorized 401)  Unauthorized  Only authenticated users can access the data
	 * @apiError (Forbidden 403)     Forbidden     Only admins can access the data
	 */
	app.Delete("/event/:id", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully deleted event with id: " + c.Params("id"))
	})

	// * --> handles all unknown routes
	app.Get("*", func(c *fiber.Ctx) {
		c.Status(404).Send("Unknown Request")
	})

	app.Listen(8080)
}
