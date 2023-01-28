package config

import (
	"os"
	"strings"
)

type CustomConfig struct {
	Port           string
	Multithreading bool
}

var Current = CustomConfig{
	Port:           ":8080",
	Multithreading: false,
}

func Init() {
	if port := os.Getenv("PORT"); len(port) > 0 {
		Current.Port = ":" + port
	}

	if useMultipleThreads := os.Getenv("USE_MULTIPLE_THREADS"); len(useMultipleThreads) > 0 {
		if strings.ToLower(useMultipleThreads) == "true" {
			Current.Multithreading = true
		}
	}
}
