// Package handler contains handlers passed to the *fiber.App application
package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// ping represents a PING payload as documented at:
// https://discord.com/developers/docs/interactions/slash-commands#receiving-an-interaction
type ping struct {
	Type int `json:"type"`
}

// NotFound returns a JSON message indicating that the given endpoint was not found
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).JSON(map[string]string{"message": "Not found"})
}

// NoContent returns a 204 status with an empty body. It may be used to check connectivity.
func NoContent(c *fiber.Ctx) error {
	c.Status(204)
	return nil
}

// Ping handles a PING payload according according to Discord's API documentation
// https://discord.com/developers/docs/interactions/slash-commands#receiving-an-interaction
func Ping(c *fiber.Ctx) error {
	payload := ping{}
	if err := c.BodyParser(&payload); err != nil {
		fmt.Println("err", err)
		return err
	}

	if payload.Type != 1 {
		c.Status(400)
		return nil
	}

	return c.Status(200).JSON(ping{Type: 1})
}
