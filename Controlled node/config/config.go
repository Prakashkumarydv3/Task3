package config

import "os"

type Config struct { // Config struct to hold application settings
	Port string
}

func Load() Config { // Load configuration from environment variables
	port := os.Getenv("PORT") // Get port from environment variable
	if port == "" {           // Default to 8080 if not set
		port = "8080"
	}

	return Config{ // Return Config with port
		Port: port,
	}
}
