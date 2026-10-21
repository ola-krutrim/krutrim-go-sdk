//go:build ignore

// Run with: go run ./examples/kks/delete_addon.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	krutrim "github.com/ola-krutrim/krutrim-go-sdk"
)

func main() {
	_ = godotenv.Load()
	if os.Getenv("KRUTRIM_CLIENT_API_KEY") == "" {
		log.Fatal("set KRUTRIM_CLIENT_API_KEY to your bearer token, without the Bearer prefix")
	}
	clusterKRN := "enter the clusterkrn"
	addonKRN := "enter the addonkrn"
	client := krutrim.NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	region := "In-Bangalore-1" // Replace with your region.

	err := client.KKS.V1.Cluster.Addon.Delete(ctx, clusterKRN, addonKRN, krutrim.KKSV1ClusterAddonDeleteParams{XRegion: region})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Add-on deletion request accepted.")
}
