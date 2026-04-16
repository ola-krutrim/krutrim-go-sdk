package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ola-krutrim/krutrim-go-sdk"
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

	// Set up the certificate ID you want to delete
	certID := "krn:kcm:In-Hyderabad-1:7344783839:2fbf2544-0c1d-481c-b745-4ef99e564b41:certs:a7005001-e9c8-40b7-a7d8-c079c135621c"

	// Call the Delete method
	err := client.Kcm.V1.Certs.Delete(ctx, certID)
	if err != nil {
		log.Fatalf("Error deleting certificate: %v", err)
	}

	// Print success message (Delete returns no response body, just error)
	fmt.Printf("Certificate %s deleted successfully\n", certID)
}
