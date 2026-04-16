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

	params := krutrim.DNSV1RecordAddParams{

		Krnid: "Enter Zone KRN",

		Rname: "Enter Record name",
		Type:  "Enter record type", // Record type: AAAA, CNAME, TXT, etc.
		Ttl:   300,                 // TTL in seconds

		Records: []krutrim.DNSV1RecordAddParamsRecord{
			{
				Value: param.Opt[string]{Value: "Enter IP value"},
			},
			// you can add more values if needed
			// {
			// 	Value: param.Opt[string]{Value: "10.1.1.2"},
			// },
		},
	}

	var resp *http.Response

	// Call API
	err := client.DNS.V1.Record.Add(ctx, params, option.WithResponseInto(&resp))
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
	fmt.Printf("Record Attached successfully:\n%s\n", string(responseJSON))
}
