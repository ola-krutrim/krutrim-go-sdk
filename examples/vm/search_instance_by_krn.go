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
	params := krutrim.HighlvlvpcGetInstanceParams{
		Krn:     "Enter the Instance KRN ID",
		XRegion: "Enter the Region",
	}

	// Call API
	resp, err := client.Highlvlvpc.GetInstance(
		ctx,
		params)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("Instance fetched successfully: %+v\n", resp)
	responseJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling response to JSON: %v", err)
	}
	fmt.Printf("Instance fetched successfully:\n%s\n", string(responseJSON))
}
