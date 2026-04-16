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
	"github.com/ola-krutrim/krutrim-go-sdk"
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

	// Certificate ID
	certID := "krn:kcm:In-Hyderabad-1:7344783839:2fbf2544-0c1d-481c-b745-4ef99e564b41:certs:610c1fc5-c24e-4cd9-bad5-d6cc24d83e97"

	// Tags to add
	tags := map[string]string{
		"xyz": "aab",
	}

	// Set up parameters
	params := krutrim.KcmV1CertTagAddParams{
		Body: tags,
	}

	var resp *http.Response

	// Call API
	err := client.Kcm.V1.Certs.Tags.Add(ctx, certID, params, option.WithResponseInto(&resp))
	if err != nil {
		log.Fatalf("Error has occurred: %v", err)
	}
	defer resp.Body.Close()

	// Read response
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
	fmt.Printf("%s\n", string(responseJSON))
}
