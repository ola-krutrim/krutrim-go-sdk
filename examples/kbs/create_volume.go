package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ola-krutrim/krutrim-go-sdk"
	"github.com/ola-krutrim/krutrim-go-sdk/packages/param"
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

	// Set up the parameters to create a new volume
	params := krutrim.KBV1VolumeNewParams{ // Replace with your availability zone
		Name:       "Enter the Volume Name",
		Size:       20,     //Enter the Volume size
		Volumetype: "HNSS", //Enter the Volume Type
		KTenantID:  "Enter the Tenant ID",

		// Optional fields
		Description: param.NewOpt("Test volume created via SDK"),
		Multiattach: param.NewOpt(false), //True or False based on requirement of multiattach
		Metadata: map[string]string{
			"environment": "",
			"team":        "",
		},

		// Source configuration (optional)
		Source: krutrim.KBV1VolumeNewParamsSource{
			ID:   param.NewOpt(""), // Optional: snapshot or volume ID
			Type: param.NewOpt(""), // Optional: "snapshot" or "volume"
		},
	}

	volume, err := client.KBV1.Volumes.New(ctx, params)
	if err != nil {
		log.Fatalf("Error getting volume: %v", err)
	}
	volumeJSON, err := json.MarshalIndent(volume, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling volume to JSON: %v", err)
	}

	// Print the volume details
	fmt.Printf("Volume retrieved successfully:\n%s\n", string(volumeJSON))
}
