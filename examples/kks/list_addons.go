//go:build ignore

// Run with: go run ./examples/kks/list_addons.go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	krutrim "github.com/ola-krutrim/krutrim-go-sdk"
)

func main() {
	_ = godotenv.Load()
	if os.Getenv("KRUTRIM_CLIENT_API_KEY") == "" {
		log.Fatal("set KRUTRIM_CLIENT_API_KEY")
	}

	client := krutrim.NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	region := "enter the region"

	res, err := client.KKS.V1.Addon.List(ctx, krutrim.KKSV1AddonListParams{XRegion: region})
	if err != nil {
		log.Fatal(err)
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(res); err != nil {
		log.Fatal(err)
	}
}
