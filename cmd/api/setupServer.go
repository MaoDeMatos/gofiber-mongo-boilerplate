package main

import (
	"github.com/maodematos/hi-gofiber/config"
	"github.com/maodematos/hi-gofiber/internal/book"
	"github.com/maodematos/hi-gofiber/internal/delayed"
	"github.com/maodematos/hi-gofiber/internal/middleware"
	"github.com/maodematos/hi-gofiber/pkg/util"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
)

func setupServer() *fiber.App {
	// Create the app
	app := fiber.New(fiber.Config{
		Prefork:      config.Current.MULTITHREADING,
		ErrorHandler: middleware.GeneralErrorHandler,
	})

	// Add middlewares. Order matters !
	app.Use(favicon.New())
	app.Use(logger.New())

	app.Use(func(ctx *fiber.Ctx) (r error) {
		ctx.Accepts("application/json")
		return ctx.Next()
	})
	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{ExposeHeaders: "X-Total-Count"}))

	app.Use(compress.New())
	// app.Use(cache.New())

	// Prevents the server from panicking
	app.Use(recover.New())

	// Routes
	app.Get("/status", func(ctx *fiber.Ctx) error {
		return ctx.JSON(util.ApiResponse{"response": "OK"})
	})

	// Test route
	app.Get("/error", func(ctx *fiber.Ctx) error {
		return fiber.ErrTeapot
	})

	v1 := app.Group("/v1", func(ctx *fiber.Ctx) error {
		ctx.Set("Version", "v1")
		return ctx.Next()
	})
	delayed.CreateRoutes(v1)
	book.CreateRoutes(v1)

	return app
}
