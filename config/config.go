package config

import (
	"github.com/joho/godotenv"
	"log"
)

//LoadEnv is load file .env
func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
