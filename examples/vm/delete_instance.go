package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	krutrim "github.com/ola-krutrim/krutrim-go-sdk"
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
	InstanceKrn := "Enter the Instance KRN"
	DeleteVolume := true //Select for deleting Volume or not

	// Call API
	resp, err := client.Highlvlvpc.DeleteInstance(
		ctx,
		InstanceKrn, DeleteVolume,
	)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("Instance created successfully:\n%s\n", string(instanceJSON))
	responseJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling response to JSON: %v", err)
	}
	fmt.Printf("Instance deleted successfully:\n%s\n", string(responseJSON))
}
