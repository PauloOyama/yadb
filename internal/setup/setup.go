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

	v0 := app.Group("/v0")

	v0.Get("/", handler.NoContent)
	v0.Post("/", handler.Ping)

	return app
}
