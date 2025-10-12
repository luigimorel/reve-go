package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	revego "github.com/luigimorel/reve-go"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := revego.NewClient(
		os.Getenv("API_KEY"))
	if err != nil {
		panic(err)
	}

	req := revego.CreateImageRequest{
		Prompt:      "Create a cart",
		AspectRatio: "16:9",
	}

	resp, err := client.CreateImage(context.Background(), req)
	if err != nil {
		panic(err)
	}

	slog.Debug("image generated", "resp", resp)
}
