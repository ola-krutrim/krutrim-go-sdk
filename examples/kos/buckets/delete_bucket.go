package main

import (
	"context"
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

	bucketKRN := "enter the bucket krn"
	xRegion := "enter the region"
	xTier := "enter the xtier"

	err := client.Ko.V1.Buckets.Delete(ctx, bucketKRN, xRegion, xTier)
	if err != nil {
		fmt.Println("API Error:", err)
		return
	}

	fmt.Println("Bucket deleted successfully")
}