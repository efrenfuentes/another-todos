package main

import (
	"github.com/efrenfuentes/todo-backend-golang-fiber/database"
	"github.com/efrenfuentes/todo-backend-golang-fiber/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	// api group
	api := app.Group("/v1")

	// send todos route group to TodoRoutes of routes package
	router.HealthCheckRoute(api.Group("/healthcheck"))
	router.TodosRoute(api.Group("/todos"))
}

func main() {
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	app.Use(logger.New())
	app.Use(cors.New(cors.ConfigDefault))

	// setup routes
	setupRoutes(app)

	err := app.Listen(":4000")
	if err != nil {
		panic(err)
	}
}
