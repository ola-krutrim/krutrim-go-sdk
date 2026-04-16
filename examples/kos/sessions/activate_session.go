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

	xRegion := "enter the xregion"
	xTier   := "enter the region"

	params := krutrim.ActivateSessionParams{
		AccessKey: "enter the accesskey",
		SecretKey: "enter the secretkey",
	}

	resp, err := client.Ko.V1.Sessions.Activate(
		ctx,
		xRegion,
		xTier,
		params,
	)
	if err != nil {
		fmt.Println("API Error:", err)
		return
	}

	out, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println("Session activated successfully:")
	fmt.Println(string(out))
}