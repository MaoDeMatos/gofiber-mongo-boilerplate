package main

import (
	"hi-gofiber/app"
	_ "hi-gofiber/config"

	// Autoload `.env`
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app.Start()
}
