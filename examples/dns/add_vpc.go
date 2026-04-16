package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	krutrim "github.com/ola-krutrim/krutrim-go-sdk"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
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

	params := krutrim.DNSV1ZoneVpcAddParams{
		// REQUIRED FIELDS
		Vpcinfo: []string{
			"Enter VPC KRN",
			// you can add more VPC KRNs if needed
			// "krn:vpc:in-bangalore-1:0987654321",
		},
	}

	zoneId := "Enter Zone KRN ID"
	var resp *http.Response

	// Call API
	err := client.DNS.V1.Zone.Vpc.Add(ctx, zoneId, params, option.WithResponseInto(&resp))
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	// Parse and print JSON response
	var response map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		log.Fatalf("Error parsing response: %v", err)
	}

	responseJSON, _ := json.MarshalIndent(response, "", "  ")
	fmt.Printf("Vpc added successfully:\n%s\n", string(responseJSON))
}
