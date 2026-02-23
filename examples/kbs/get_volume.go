package main

import (
	"context"
	"encoding/json"
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
	volumeID := "Enter the Volume ID"

	// Set up the query parameters with required k-tenant-id header
	params := krutrim.KBV1VolumeGetParams{
		KTenantID: "Enter the K tenant ID", // Replace with actual tenant ID
	}

	// Call the Get method - returns already decoded response
	volume, err := client.KBV1.Volumes.Get(ctx, volumeID, params)
	if err != nil {
		log.Fatalf("Error getting volume: %v", err)
	}

	// Marshal the already decoded response to pretty JSON
	volumeJSON, err := json.MarshalIndent(volume, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling volume to JSON: %v", err)
	}

	// Print the volume details
	fmt.Printf("Volume retrieved successfully:\n%s\n", string(volumeJSON))
}
