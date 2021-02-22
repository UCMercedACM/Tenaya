package main

import (
	"context"
	"log"
	"os"
    "time"

    "github.com/UCMercedACM/Tenaya/pkg/entities"
    "github.com/UCMercedACM/Tenaya/pkg/event"
    "github.com/UCMercedACM/Tenaya/api/routes"

	"github.com/go-pg/pg/v10"
    "github.com/go-pg/pg/v10/orm"
    "github.com/go-pg/pg/extra/pgdebug"
    "github.com/go-pg/pg/extra/pgotel"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/recover"
    "github.com/gofiber/fiber/v2/middleware/pprof"
    "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet"
    "github.com/joho/godotenv"
)

// Database instance
var db *pg.DB

// Database settings
var (
	dbHost     = "localhost"
	dbPort     = "5432" // Default port
	dbUser     = "postgres"
	dbPassword = "password"
	dbDatabase = "fiber_demo"
)

func createSchema(db *pg.DB) error {
	createTableOpts := &orm.CreateTableOptions{
        Varchar: 255,
		IfNotExists: true,
	}

	if err := db.Model(*entities.EventModel).CreateTable(createTableOpts); err != nil {
        log.Fatal("Could not create table: $s", err)
		return err
	}

	return nil
}

// DatabaseConnection -->  should establish connection to database otherwise error out
func DatabaseConnection() (*pg.DB, error) {
	// Database Credentials
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbDatabase = os.Getenv("DB_DATABASE")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")

	// Open connection to database
	db := pg.Connect(&pg.Options{
		Addr:            dbHost + ":" + dbPort,
		User:            dbUser,
		Password:        dbPassword,
		Database:        dbDatabase,
		ApplicationName: "Tenaya",
    })

    if debug {
        db.AddQueryHook(pgdebug.DebugHook{
            Verbose: true,
        })
    }

    db.AddQueryHook(pgotel.TracingHook{})

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := db.Ping(ctx); err != nil {
		log.Fatal("Was unable to ping the database: $s", err)
	}

	if err := createSchema(db); err != nil {
		log.Fatal("Could not create schema: $s", err)
    }

    return db, nil
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	if port := os.Getenv("PORT"); port == "" {
		log.Fatal("$PORT must be set")
	}

	if db, err := DatabaseConnection(); err != nil {
		log.Fatal("Database Connection Error $s", err)
    }

    fmt.Println("Database connection success!")

	// Create a Fiber app
	app := fiber.New(fiber.Config{
        // Override default error handler
        ErrorHandler: func(ctx *fiber.Ctx, err error) error {
            // Statuscode defaults to 500
            code := fiber.StatusInternalServerError

            // Retreive the custom statuscode if it's an fiber.*Error
            if e, ok := err.(*fiber.Error); ok {
                code = e.Code
            }

            // Send custom error page
            if err = ctx.Status(code); err != nil {
                // In case the SendFile fails
                return ctx.Status(500).SendString("Internal Server Error")
            }

            // Return from handler
            return nil
        },
    })

	// Helmet to protect cross-site scripting (XSS) attack
	app.Use(helmet.New())

	// Gives web servers cross-domain access controls, which enable secure cross-domain data transfers
    app.Use(cors.New())

    app.Use(pprof.New())

    app.Use(requestid.New())

    ​app​.​Use​(​logger​.​New​(logger.​Config​{
        // For more options, see the Config section
        Format​: "[${time}] ${pid} ${locals:requestid} ${status} - ${latency} ${method} ${path}​\n​"​,
        TimeFormat:   "15:04:05",
        TimeZone:     "Local",
        TimeInterval: 500 * time.Millisecond,
        Output:       os.Stderr,
    }))

	// Recover from panic errors within any route
	app.Use(recover.New(recover.Config{
		Handler: func(c *fiber.Ctx, err error) {
			c.SendString(err.Error())
			c.SendStatus(500)
		},
	}))

	// USE: /api --> assigns all headers for routes under api route
	app.Use("/", func(c *fiber.Ctx) {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Headers", "X-Requested-With")
		c.Set("Content-Type", "application/json")
		c.Next()
    })

    api := app.Group("/api")
    routes.EventRouter(api, eventService)

    app.Get("*", func(c  *fiber.Ctx) {
        c.Status(404).Send("Unknown Request")
    })

    data, _ := json.MarshalIndent(app.Stack(), "", "  ")
    fmt.Println(string(data))

    _ = app.Listen(":" + port)
}
