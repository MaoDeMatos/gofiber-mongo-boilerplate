package main

import (
	"github.com/maodematos/hi-gofiber/config"
	"github.com/maodematos/hi-gofiber/pkg/database"
	"github.com/maodematos/hi-gofiber/pkg/shutdown"

	// Autoload `.env`
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.Init()

	cleanup, err := run()
	if err != nil {
		panic(err)
	}

	defer cleanup()
	shutdown.Gracefully()
}

func run() (func(), error) {
	app := setupServer()

	// Start the server
	go func() {
		// Connect to database
		database.StartMongoDB()
		app.Listen(":" + config.Current.PORT)
	}()

	// Return a function to close the server and database
	return func() {
		database.CloseMongoDB()
		app.Shutdown()
	}, nil
}
