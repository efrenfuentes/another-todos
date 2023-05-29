package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// HealthCheck is a simple health check endpoint
func HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "available",
		"message": "Simple TODO API with Golang and Fiber",
	})
}
