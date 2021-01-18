package main

import (
	"log"
	"password-manager/cmd"

	"github.com/joho/godotenv"
)

func init() {
	// load env variable from .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Fatal(cmd.Execute())
}
