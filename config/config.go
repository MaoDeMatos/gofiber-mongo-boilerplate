package config

import (
	"os"
	"strings"
)

type CustomConfig struct {
	PORT           string
	MULTITHREADING bool
	MONGODB_URI    string
}

var Current = CustomConfig{
	PORT:           "8080",
	MULTITHREADING: false,
	MONGODB_URI:    "mongodb://localhost:27017",
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

	if mongoUri := os.Getenv("MONGO_URI"); len(mongoUri) > 0 {
		Current.MONGODB_URI = mongoUri
	}

}
