package pkg

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	pwd, _ := os.Getwd()

	err := godotenv.Load(pwd, ".env")

	if err != nil {
		panic("Error loading .env file")
	}

	SendInfoLog("Environment variables loaded successfully")

}
