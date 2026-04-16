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

	params := krutrim.DNSV1RecordUpdateParams{
		Rname:   param.Opt[string]{Value: "Enter new record name"},  //(optional)
		Type:    param.Opt[string]{Value: "Enter new record Type"},  // (optional)
		Routing: param.Opt[string]{Value: "Enter new routing name"}, // (optional)

		Records: []krutrim.DNSV1RecordUpdateParamsRecord{
			{
				Value: param.Opt[string]{Value: "Enter new address value"},
			},
			// You can add more values if needed
		},
	}
	recordId := "Enter Record Id"

	var resp *http.Response

	// Call API
	err := client.DNS.V1.Record.Update(ctx, recordId, params, option.WithResponseBodyInto(&resp))
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
	fmt.Printf("Record updated successfully:\n%s\n", string(responseJSON))
}
