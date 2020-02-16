package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-pg/pg/v9"
	"github.com/gofiber/fiber"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/joho/godotenv"
	// "github.com/go-pg/pg/v9/orm" // Incase we want to use the orm
)

// Event --> Struct defining the global data structure for all events
type Event struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Date        string `json:"date"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	CreatedAt   string `json:"createdAt"`
}

// Workshop --> Struct defining the global data structure for all events
type Workshop struct {
	Key int64 `json:"key"`
}

// GetAllEvents --> Returns array of all events and errors if they exist
func GetAllEvents(db *pg.DB) ([]Event, error) {
	var events []Event
	_, err := db.Query(&events, `select * from events`)
	return events, err
}

// GetEventsByType --> Returns array of all events based off the type of events requested and errors if they exist
func GetEventsByType(db *pg.DB, eventType string) ([]Event, error) {
	var events []Event
	_, err := db.Query(&events, `SELECT * FROM events WHERE type = ?`, eventType)
	return events, err
}

// GetEventByID --> Returns specific event based off the id number
func GetEventByID(db *pg.DB, id int64) (*Event, error) {
	var event Event
	_, err := db.QueryOne(&event, `SELECT * FROM events WHERE id = ?`, id)
	return &event, err
}

// CreateEvent --> Returns specific event based off the id number
func CreateEvent(db *pg.DB, event Event) error {
	fmt.Println(event)
	_, err := db.QueryOne(event, `
		INSERT INTO events (name, description, type, date, start_time, end_time) VALUES (?name, ?description, ?type, ?date, ?start_time, ?end_time)
		RETURNING id
	`, event)
	return err
}

// INSERT INTO events (name, description, type, date, starttime, endtime) VALUES ("test event", "This is a test event", "test", "February 14, 2020", "12:00)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Loads environment variables and handles errors
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Database Credentials
	dbAddress := os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	// Open connection to database
	db := pg.Connect(&pg.Options{
		Addr:            dbAddress,
		User:            dbUser,
		Password:        dbPassword,
		Database:        dbDatabase,
		ApplicationName: "Onama",
	})

	// Check if event table has already been created and create if it has not
	db.Exec(`CREATE TABLE IF NOT EXISTS EVENTS (
		ID serial PRIMARY KEY NOT NULL, 
		name varchar(255) NOT NULL, 
		description varchar(255), 
		type varchar(255), 
		date varchar(255), 
		start_time varchar(255), 
		end_time varchar(255), 
		created_at TIMESTAMPTZ DEFAULT NOW()
	);`)

	// Create new fiber server
	app := fiber.New()

	// USE: /api --> assigns all headers for routes under api route
	app.Use("/api", func(c *fiber.Ctx) {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Headers", "X-Requested-With")
		c.Set("Content-Type", "application/json")
		c.Next()
	})

	app.Static("templates/index.html")
	app.Static("/static", "static")

	/**
	 * @api {GET} /api/events
	 * @apiDescription Returns all calendar events
	 * @apiVersion 1.1.0
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
		events, err := GetAllEvents(db)
		panicIf(err)

		c.Status(200).Send("All Event Data")
		c.JSON(events)
	})

	/**
	 * @api {GET} /api/events/:type
	 * @apiDescription Returns only a subgroup of all events
	 * @apiVersion 1.1.0
	 * @apiName Get Events Based on Type
	 * @apiGroup Event(s)
	 * @apiPermission public
	 *
	 * @param type Defines the type of event
	 *
	 * @apiSuccess {Object[]} Returns an object array of all event information.
	 *
	 * @apiError (Unauthorized 401)  Unauthorized  Only authenticated users can access the data
	 * @apiError (Forbidden 403)     Forbidden     Only admins can access the data
	 */
	app.Get("/api/events/:type", func(c *fiber.Ctx) {
		events, err := GetEventsByType(db, c.Params("type"))
		panicIf(err)

		c.Status(200).Send("Event Data for: " + c.Params("type"))
		c.JSON(events)
	})

	/**
	 * @api {GET} /api/event/:id
	 * @apiDescription Returns a single event
	 * @apiVersion 1.1.0
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
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		event, err := GetEventByID(db, id)
		panicIf(err)

		c.Status(200).Send("Displaying Event Data for: " + c.Params("id"))
		c.JSON(event)
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
	app.Post("/api/events", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully created new events")
	})

	/**
	 * @api {POST} /api/event
	 * @apiDescription Create a new event
	 * @apiVersion 1.1.0
	 * @apiName Create Event
	 * @apiGroup Event(s)
	 * @apiPermission public
	 *
	 * @apiSuccess {string} Returns a success message logging that the new event were created successfully.
	 *
	 * @apiError (Unauthorized 401)  Unauthorized  Only authenticated users can access the data
	 * @apiError (Forbidden 403)     Forbidden     Only admins can access the data
	 */
	app.Post("/api/event", func(c *fiber.Ctx) {
		var event Event

		if err := c.BodyParser(&event); err != nil {
			panicIf(err)
		}

		err := CreateEvent(db, event)
		panicIf(err)

		c.Status(200).Send("Successfully created a new event")
		c.JSON(c.BodyParser(event))
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
	app.Patch("/api/events", func(c *fiber.Ctx) {
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
	app.Patch("/api/events/:type", func(c *fiber.Ctx) {
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
	app.Patch("/api/event/:id", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully updated event: " + c.Params("id"))
	})

	/**
	 * @api {DELETE} /api/events
	 * @apiDescription Deletes all events in the database
	 * @apiVersion 1.0.0
	 * @apiName Delete All Events
	 * @apiGroup Event(s)
	 * @apiPermission admin
	 *
	 * @apiSuccess {string} Returns a success message logging that the events were deleted successfully.
	 *
	 * @apiError (Unauthorized 401)  Unauthorized  Only authenticated users can access the data
	 * @apiError (Forbidden 403)     Forbidden     Only admins can access the data
	 */
	app.Delete("/api/events", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully deleted all events")
	})

	/**
	 * @api {DELETE} /api/events/:type
	 * @apiDescription Deletes all events within a group type in the database
	 * @apiVersion 1.0.0
	 * @apiName Delete All Events Under a Group
	 * @apiGroup Event(s)
	 * @apiPermission admin
	 *
	 * @apiSuccess {string} Returns a success message logging that the events were deleted successfully.
	 *
	 * @apiError (Unauthorized 401)  Unauthorized  Only authenticated users can access the data
	 * @apiError (Forbidden 403)     Forbidden     Only admins can access the data
	 */
	app.Delete("/api/events/:type", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully deleted all events of type: " + c.Params(("type")))
	})

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
	app.Delete("/api/event/:id", func(c *fiber.Ctx) {
		c.Status(200).Send("Successfully deleted event with id: " + c.Params("id"))
	})

	// * --> handles all unknown routes
	app.Get("*", func(c *fiber.Ctx) {
		c.Status(404).Send("Unknown Request")
	})

	app.Recover(func(c *fiber.Ctx) {
		c.Status(500).Send(c.Error())
	})

	app.Listen(port)
}
