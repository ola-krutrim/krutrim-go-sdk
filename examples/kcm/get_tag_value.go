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

	client := krutrim.NewClient()

	ctx := context.Background()

	certID := "enter the certid"

	tagName := "enter the tagname"

	params := krutrim.KcmV1CertTagGetByNameParams{
		CertID:  certID,
		TagName: tagName,
	}

	var resp *http.Response

	err := client.Kcm.V1.Certs.Tags.GetByName(ctx, params, option.WithResponseInto(&resp))
	if err != nil {
		log.Fatalf("Error has occurred: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		log.Fatalf("Error parsing response: %v", err)
	}

	responseJSON, _ := json.MarshalIndent(response, "", "  ")
	fmt.Printf("%s\n", string(responseJSON))
}
