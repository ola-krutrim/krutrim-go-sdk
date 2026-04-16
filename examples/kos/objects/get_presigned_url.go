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

	bucketKRN := "enter the bucket krn"
	objectKey := "enter the object key"

	xRegion := "enter the region"
	xSessionToken := "enter the session token"

	resp, err := client.Ko.V1.Objects.GetPreSignedDownloadURL(
		ctx,
		bucketKRN,
		objectKey,
		xRegion,
		xSessionToken,
	)
	if err != nil {
		fmt.Println("API Error:", err)
		return
	}

	out, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(out))
}