package main

import (
	"locket-clone/backend/pkg/api/rest"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load .env")
	}

	rest.Init()
}
