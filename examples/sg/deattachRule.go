package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	krutrim "github.com/ola-silicon/krutrim-go-sdk"
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

	params := krutrim.SecurityGroupV1DetachRuleParams{
		// REQUIRED FIELDS
		Vpcid:      "Enter the VPC KRN",
		Securityid: "Enter the SG KRN",
		Ruleid:     "Enter the SGR KRN",
		XRegion:    "Enter the Region",
	}

	// Call API
	resp, err := client.SecurityGroup.V1.DetachRule(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("SG Rule deattached successfully: %+v\n", resp)
	responseJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling response to JSON: %v", err)
	}
	fmt.Printf("SG Rule deattached successfully:\n%s\n", string(responseJSON))
}
