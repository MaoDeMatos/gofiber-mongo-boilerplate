package main

import (
	"hi-gofiber/app"
	"hi-gofiber/config"

	// Autoload `.env`
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config.Init()
	app.Start()
}
