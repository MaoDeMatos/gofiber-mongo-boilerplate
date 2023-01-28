package config

import (
	"os"
)

type CustomConfig = struct {
	Port string
}

var Current = CustomConfig{
	Port: ":8080",
}

func init() {
	if port := os.Getenv("PORT"); len(port) > 0 {
		Current.Port = ":" + port
	}
}
