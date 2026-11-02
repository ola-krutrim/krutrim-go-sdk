//go:build ignore

// Run with: go run ./examples/kks/delete_node_group.go
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
		log.Fatal("set KRUTRIM_CLIENT_API_KEY")
	}
	clusterKRN := "enter the clusterkrn"
	nodeGroupKRN := "enter the nodegroupkrn"
	client := krutrim.NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	region := "enter the region"

	err := client.KKS.V1.Cluster.NodeGroup.Delete(ctx, clusterKRN, nodeGroupKRN, krutrim.KKSV1ClusterNodeGroupDeleteParams{XRegion: region})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Node group deletion request accepted.")
}
