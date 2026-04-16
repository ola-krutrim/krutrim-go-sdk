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

	params := krutrim.ObjectRenameParams{
		XRegion:       "enter the region",
		XSessionToken:"enter the session token",

		BucketKRN: "enter the bucket krn",
		OldKey:    "enter the object key",    //test.img
		NewKey:    "enter the newkey",
	}
	
	res, err := client.Ko.V1.Objects.Rename(ctx, params)
	if err != nil {
		fmt.Println("API Error:", err)
		return
	}
	
	b, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(b))
}