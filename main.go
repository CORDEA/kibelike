package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Failed to load .env")
	}
	org := os.Getenv("ORG")
	token := os.Getenv("TOKEN")
	client := NewClient("https://"+org+".kibe.la/api/v1", token)
}
