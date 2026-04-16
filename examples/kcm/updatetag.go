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

	// Create SDK client
	client := krutrim.NewClient()

	ctx := context.Background()

	// Certificate ID
	certID := "enter the certid"

	// Tags to add
	body := krutrim.KcmV1CertTagAddParams{
		Body: map[string]string{
			"env":     "prod", //enter as per requirement
			"project": "demo",
		},
	}

	var resp *http.Response

	// Call API
	err := client.Kcm.V1.CertTag.Add(
		ctx,
		certID,
		body,
		option.WithResponseInto(&resp),
	)

	if err != nil {
		log.Fatalf("Error has occurred: %v", err)
	}

	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	// Convert response to JSON
	var response interface{}
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		log.Fatalf("Error parsing response: %v", err)
	}

	responseJSON, _ := json.MarshalIndent(response, "", "  ")
	fmt.Printf("%s\n", string(responseJSON))
}
