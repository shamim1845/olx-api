package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds the configuration for the application
type Config struct {
	Port   string
	Env    string
	DB_Url string
}

// MustLoad loads the configuration from the environment variables.
// It attempts to load a .env file for local development but does not
// panic if the file is absent (e.g., in production on Render).
func MustLoad() Config {
	// load .env file – ignore error so production works without a file
	_ = godotenv.Load()

	// Get port from environment variable
	Port := os.Getenv("PORT")
	if Port == "" {
		panic("PORT environment variable is not set")
	}

	// Get env from environment variable
	Env := os.Getenv("ENV")
	if Env == "" {
		panic("ENV environment variable is not set")
	}

	// Get DB_URL from environment variable
	DB_Url := os.Getenv("DB_URL")
	if DB_Url == "" {
		panic("DB_URL environment variable is not set")
	}

	return Config{
		Port:   Port,
		Env:    Env,
		DB_Url: DB_Url,
	}
}
