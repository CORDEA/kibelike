package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
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

	subcommands.Register(&SearchCmd{client: client}, "")
	subcommands.Register(&ListCmd{client: client}, "")
	subcommands.Register(&LikeCmd{client: client}, "")
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
