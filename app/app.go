package app

import (
	"hi-gofiber/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
)

func setup() *fiber.App {
	app := fiber.New()

	// Order of the middlewares is important !
	app.Use(favicon.New())
	app.Use(logger.New())
	app.Use(func(c *fiber.Ctx) (r error) {
		c.Accepts("application/json")
		return c.Next()
	})

	app.Use(cache.New())

	app.Use(helmet.New())
	app.Use(cors.New())

	app.Use(compress.New())

	return app
}

func Start() {
	app := setup()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"response": "Pong !"})
	})

	app.Listen(config.Current.Port)
}
