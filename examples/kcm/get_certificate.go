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

	// Set up the certificate ID you want to retrieve
	certID := "krn:kcm:In-Hyderabad-1:7344783839:2fbf2544-0c1d-481c-b745-4ef99e564b41:certs:6f563b64-41e0-4758-9e4f-f6b7c179ad21" // Replace with actual certificate ID

	// Method 1: Use List with certID parameter (recommended)
	params := krutrim.KcmV1CertListParams{
		CertID: param.NewOpt(certID),
	}

	cert, err := client.Kcm.V1.Certs.List(ctx, params)
	if err != nil {
		log.Fatalf("Error getting certificate: %v", err)
	}

	// Marshal to pretty JSON
	certJSON, err := json.MarshalIndent(cert, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling certificate to JSON: %v", err)
	}

	// Print the certificate details
	fmt.Printf("Certificate retrieved successfully:\n%s\n", string(certJSON))

	// Alternative Method 2: Use Get directly (after fixing the method)
	// cert, err := client.Kcm.V1.Certs.Get(ctx, certID)
	// if err != nil {
	// 	log.Fatalf("Error getting certificate: %v", err)
	// }
	// certJSON, err := json.MarshalIndent(cert, "", "  ")
	// if err != nil {
	// 	log.Fatalf("Error marshaling certificate to JSON: %v", err)
	// }
	// fmt.Printf("Certificate retrieved successfully:\n%s\n", string(certJSON))
}
