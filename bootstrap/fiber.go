package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func newFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: getApplicationName(),
	})

	app.Use(logger.New())
	app.Use(recover.New())

	app.Use(cors.Config{
		AllowOrigins: getAllowedOrigins(),
		AllowMethods: getAllowedMethods(),
	})

	return app
}
