package router

import (
	"github.com/efrenfuentes/todo-backend-golang-fiber/internals/handlers"
	"github.com/gofiber/fiber/v2"
)

func HealthCheckRoute(route fiber.Router) {
	route.Get("", handlers.HealthCheck)
}
