package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load environment variables
	env := os.Getenv("ENV")

	if env == "" || env == "development" {
		// Load .env only for dev environment
		if err := godotenv.Load(); err != nil {
			fmt.Println("⚠️ Warning: no .env file found, relying on environment variables")
		} else {
			fmt.Println("✅ .env loaded successfully")
		}
	} else {
		fmt.Printf("✅ Skipping .env load for environment: %s", env)
	}
}
