package config

import "github.com/joho/godotenv"

func LoadEnvs() {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load .env file")
	}
}
