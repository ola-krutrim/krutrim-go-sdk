package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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


	bucketKrn := "Enter the bucket KRN" // e.g., "krn:kos:bucket:region:account-id:bucket-name"

	// Call API
	resp, err := client.Ko.V1.Buckets.Get(ctx, bucketKrn)
	if err != nil {
		log.Fatal(err)
	}

	responseJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling response: %v", err)
	}

	fmt.Printf("Bucket details:\n%s\n", string(responseJSON))
}
