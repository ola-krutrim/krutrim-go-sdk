package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	krutrim "github.com/ola-krutrim/krutrim-go-sdk"
)

func main() {
	_ = godotenv.Load()

	if os.Getenv("KRUTRIM_CLIENT_API_KEY") == "" {
		fmt.Println("API key missing")
		return
	}

	client := krutrim.NewClient()
	ctx := context.Background()

	params := krutrim.BucketCreateParams{
		XRegion: "enter the region",
		XTier:   "ente the tier",

		Name:        "enter the bucket name",
		Description: "enter the description",
		Tier:        "enter the tier",

		Versioning:      false,
		AnonymousAccess: false,
		Tags:            map[string]string{},
	}
	resp, err := client.Ko.V1.Buckets.Create(ctx, params)
	if err != nil {
		fmt.Println("API Error:", err)
		return
	}

	// Print EXACT API response (same fields as curl)
	out, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(out))
}