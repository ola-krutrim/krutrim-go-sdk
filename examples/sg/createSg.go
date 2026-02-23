package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	krutrim "github.com/ola-silicon/krutrim-go-sdk"
	"github.com/ola-silicon/krutrim-go-sdk/packages/param"
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

	params := krutrim.SecurityGroupV1NewParams{
		// REQUIRED FIELDS
		Vpcid:       "Enter the VPC KRN Id",
		Name:        "Enter the SG name",
		XRegion:     "Enter the Region",
		Description: param.NewOpt("Enter the Description of SG"),
	}

	// Call API
	resp, err := client.SecurityGroup.V1.New(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("SG created successfully: %+v\n", resp)
}
