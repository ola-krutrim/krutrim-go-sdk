package main

import (
	"context"
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

	// Certificate ID to update
	certID := "krn:kcm:In-Hyderabad-1:7344783839:2fbf2544-0c1d-481c-b745-4ef99e564b41:certs:610c1fc5-c24e-4cd9-bad5-d6cc24d83e97"

	// Path to your new/updated certificate file
	certFilePath := "/Users/divya.sonthalia/Downloads/bundle.p12"

	// Open the certificate file
	certFile, err := os.Open(certFilePath)
	if err != nil {
		log.Fatalf("Error opening certificate file: %v", err)
	}
	defer certFile.Close()

	// VPC ID (required in header)
	vpcID := "krn:vpc:In-Hyderabad-1:7344783839:2fbf2544-0c1d-481c-b745-4ef99e564b41:vpc:5891272d-60f5-4b8d-a11a-26a3af34c588"

	// Set up the parameters to update the certificate
	params := krutrim.KcmV1CertUpdateParams{
		XVpcID:   vpcID,                  // Required: VPC ID in header
		CertFile: certFile,               // Required: New certificate file
		Flag:     param.NewOpt[int64](0), // Optional: Flag value
	}

	// Call the Update method
	err = client.Kcm.V1.Certs.Update(ctx, certID, params)
	if err != nil {
		log.Fatalf("Error updating certificate: %v", err)
	}

	// Print success message (Update returns no response body, just error)
	fmt.Printf("Certificate %s updated successfully\n", certID)
}
