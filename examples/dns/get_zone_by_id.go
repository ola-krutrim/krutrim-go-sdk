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

	zoneId := "Enter Zone KRN ID"

	var resp *http.Response

	err := client.DNS.V1.Zone.Get(ctx, zoneId, option.WithResponseInto(&resp))
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("SG Rule Attached successfully: %+v\n", resp)
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
	fmt.Printf("Zone fetched successfully:\n%s\n", string(responseJSON))
}
