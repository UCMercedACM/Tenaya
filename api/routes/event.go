package routes

import (
	"github.com/UCMercedACM/Tenaya/pkg/entities"
	"github.com/UCMercedACM/Tenaya/pkg/event"
	"github.com/gofiber/fiber/v2"
)

// EventRouter --> middleware to handle incoming requests
func EventRouter(app fiber.Route, service event.Service) {
	app.Get("/events", getEvents(service))
	app.Get("/events/:type", getEventsByType(service))
	app.Get("/event/:id", getEventByID(service))
	app.Post("/event", addEvent(service))
	app.Patch("/event", updateEvent(service))
	app.Delete("/event", removeEvent(service))
}

func getEvents(service event.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if fetched, err := service.FetchEvents(); err != nil {
			_ = c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}

		return c.JSON(&fiber.Map{
			"status": true,
			"events": fetched,
		})
	}
}

func getEventsByType(service event.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
        eventType := c.Params("type")

		if fetched, err := service.FetchEventsByType(eventType); err != nil {
			_ = c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}

		return c.JSON(&fiber.Map{
			"status": true,
			"events": fetched,
		})
	}
}

func getEventByID(service event.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
        id := c.Params("id")

		if fetched, err := service.FetchEventByID(id); err != nil {
			_ = c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}

		return c.JSON(&fiber.Map{
			"status": true,
			"events": fetched,
		})
	}
}

func addEvent(server event.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		var requestBody entities.Event

		if err := c.BodyParser(&requestBody); err != nil {
			_ = c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		if result, dberr := service.InsertEvent(&requestBody); dberr != nil {
			_ = c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}

		return c.JSON(&fiber.Map{
			"status": result,
			"error":  dberr,
		})
	}
}

func updateEvent(service event.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		var requestBody entities.Event

		if err := c.BodyParser(&requestBody); err != nil {
			_ = c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		if result, dberr := service.UpdateEvent(&requestBody); dberr != nil {
			_ = c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
        }

		return c.JSON(&fiber.Map{
			"status": result,
			"error":  dberr,
		})
	}
}

func removeEvent(service event.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		var requestBody entities.DeleteRequest

		if err := c.BodyParser(&requestBody); err != nil {
			_ = c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}

		if dberr := service.RemoveEvent(requestBody.ID); dberr != nil {
			_ = c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}

		return c.JSON(&fiber.Map{
			"status":  false,
			"message": "updated successfully",
		})
	}
}
