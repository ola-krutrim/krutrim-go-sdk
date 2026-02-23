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
	params := krutrim.HighlvlvpcSearchInstancesParams{
		VpcID: "Enter the VPC KRN ID",
		Limit: 10, //Enter the Limit of response on page
		Page:  1,  //Enter number of pages
	}

	// Call API
	resp, err := client.Highlvlvpc.SearchInstances(
		ctx,
		params)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Instance fetched successfully: %+v\n", resp)
	responseJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling response to JSON: %v", err)
	}
	fmt.Printf("Instance fetched successfully:\n%s\n", string(responseJSON))
}
