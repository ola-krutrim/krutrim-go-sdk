package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	krutrim "github.com/ola-krutrim/krutrim-go-sdk"
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

	params := krutrim.SecurityGroupV1ListParams{
		// REQUIRED FIELDS
		Limit:     param.NewOpt(int64(10)),
		Page:      param.NewOpt(int64(1)),
		SortBy:    param.NewOpt("createdAt"),
		SortOrder: krutrim.SecurityGroupV1ListParamsSortOrderAsc,
		XRegion:   "Enter the Region",
	}

	vpcKrn := "Enter the VPC KRN Id"

	// Call API
	resp, err := client.SecurityGroup.V1.List(ctx, vpcKrn, params)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("List of SG fetched successfully: %+v\n", resp)
	responseJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling response to JSON: %v", err)
	}
	fmt.Printf("List of SG fetched successfully:\n%s\n", string(responseJSON))
}
