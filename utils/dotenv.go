package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GodotEnv(key string) string {
	env := make(chan string, 1)

	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			return ""
		}
		env <- os.Getenv(key)
	} else {
		env <- os.Getenv(key)
	}

	return <-env
}
