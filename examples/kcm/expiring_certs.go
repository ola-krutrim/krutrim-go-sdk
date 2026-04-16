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

	// VPC ID
	vpcID := "krn:vpc:In-Hyderabad-1:7344783839:2fbf2544-0c1d-481c-b745-4ef99e564b41:vpc:5891272d-60f5-4b8d-a11a-26a3af34c588"

	// Date in simple format (YYYY-MM-DD)
	dateStr := "2026-12-31"

	var resp *http.Response

	// Call API with custom query string
	err := client.Kcm.V1.Certs.GetExpiring(
		ctx,
		krutrim.KcmV1CertGetExpiringParams{},
		option.WithResponseInto(&resp),
		option.WithQuery("date", dateStr),
		option.WithQuery("vpcId", vpcID),
	)
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
