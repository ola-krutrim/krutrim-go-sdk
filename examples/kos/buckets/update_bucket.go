package main

import (
	"context"
	"fmt"
	"os"
	"encoding/json"

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

	params := krutrim.BucketUpdateParams{
		XRegion: "enter the region",
		XTier:   "enter the xtier",

		BucketKRN: "enter the bucketkrn",

		Versioning:      false,
		AnonymousAccess: true,
		Tags: map[string]string{
			"environment": "dev",
			"project":     "storage-test",          //change according to your requirement
		},
	}

	res, err := client.Ko.V1.Buckets.Update(ctx, params)
	if err != nil {
		fmt.Println("API Error:", err)
		return
	}

	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println("JSON Marshal Error:", err)
		return
	}

	fmt.Println(string(b))
}