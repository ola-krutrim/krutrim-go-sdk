package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ola-silicon/krutrim-go-sdk"
)

func main() {

	// Load .env file
	_ = godotenv.Load()

	if os.Getenv("KRUTRIM_CLIENT_API_KEY") == "" {
		fmt.Println("API key missing")
		return
	}

	// Create SDK client (auth picked from env internally)
	client := krutrim.NewClient()

	ctx := context.Background()

	region := "Enter the Region"

	resp, err := client.Vm.V1.Image.List(ctx, region)
	if err != nil {
		log.Fatalf("Error getting resp: %v", err)
	}

	fmt.Printf("Images fetched successfully:\n%s\n", resp)
}
