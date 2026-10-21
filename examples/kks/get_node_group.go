//go:build ignore

// Run with: go run ./examples/kks/get_node_group.go
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
		log.Fatal("set KRUTRIM_CLIENT_API_KEY to your bearer token, without the Bearer prefix")
	}
	clusterKRN := "enter the clusterkrn"
	nodeGroupKRN := "enter the nodegroupkrn"
	client := krutrim.NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	region := "In-Bangalore-1" // Replace with your region.

	res, err := client.KKS.V1.Cluster.NodeGroup.Get(ctx, clusterKRN, nodeGroupKRN, krutrim.KKSV1ClusterNodeGroupGetParams{XRegion: region})
	if err != nil {
		log.Fatal(err)
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(res); err != nil {
		log.Fatal(err)
	}
}
