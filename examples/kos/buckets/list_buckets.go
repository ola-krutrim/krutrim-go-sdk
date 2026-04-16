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

	params := &krutrim.BucketListParams{
		Tier: "tier-1",
	}

	res, err := client.Ko.V1.Buckets.List(ctx, params)
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