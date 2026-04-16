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

	xRegion := "enter the region"
	xTier := "enter the tier"

	resp, err := client.Ko.V1.AccessKeys.Create(ctx, xRegion, xTier)
	if err != nil {
		fmt.Println("API Error:", err)
		return
	}

	out, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(out))
}
