// Package handler contains handlers passed to the *fiber.App application
package handler

import "github.com/gofiber/fiber/v2"

// NotFound returns a JSON message indicating that the given endpoint was not found
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).JSON(map[string]string{"message": "Not found"})
}
