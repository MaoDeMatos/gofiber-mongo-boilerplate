package main

import (
	"github.com/maodematos/hi-gofiber/app"
	"github.com/maodematos/hi-gofiber/config"

	// Autoload `.env`
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config.Init()
	app.Start()
}
