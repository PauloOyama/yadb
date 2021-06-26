// Package setup contains functions that setup a simple configuration of a *fiber.App object
package setup

import (
	"github.com/agstrc/yadb/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func App() *fiber.App {
	app := fiber.New()
	app.Use(logger.New(), recover.New(), verify)
	defer app.Use(handler.NotFound) // call is deferred as the NotFound handler must be last in stack

	app.Get("/", noContent)

	return app
}

// noContent returns a 204 status with an empty body. It may be used to check connectivity.
func noContent(c *fiber.Ctx) error {
	c.Status(204)
	return nil
}
