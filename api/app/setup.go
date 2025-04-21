package app

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gramkow-dev/split-it/router"
)

func SetupAndRunApp() error {

	// load env

	// start database

	// create app
	app := fiber.New()

	// attach middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	// setup routes
	router.SetupRoutes(app)

	// get port and start
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	app.Listen(":" + port)

	return nil
}
