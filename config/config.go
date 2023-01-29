package config

import (
	"os"
	"strings"
)

type CustomConfig struct {
	PORT           string
	MULTITHREADING bool
}

var Current = CustomConfig{
	PORT:           "8080",
	MULTITHREADING: false,
}

func Init() {
	if port := os.Getenv("PORT"); len(port) > 0 {
		Current.PORT = port
	}

	if useMultipleThreads := os.Getenv("USE_MULTIPLE_THREADS"); len(useMultipleThreads) > 0 {
		if strings.ToLower(useMultipleThreads) == "true" {
			Current.MULTITHREADING = true
		}
	}
}
