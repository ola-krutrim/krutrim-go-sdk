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

	// Set up the volume ID you want to retrieve
	volumeID := "Enter the Volume Krn Id"

	// Set up the query parameters with required k-tenant-id header
	Params := krutrim.KBV1VolumeDeleteParams{
		KTenantID: "Enter the K tenant Id",
	}

	// Call the Delete method
	err := client.KBV1.Volumes.Delete(ctx, volumeID, Params)
	if err != nil {
		log.Fatalf("Error deleting volume: %v", err)
	}

	// Print success message
	fmt.Println("Volume deleted successfully")
}
