package initialisers

import (
	"log"

	"github.com/joho/godotenv"
)

func GetEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
