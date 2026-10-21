//go:build ignore

// Run with: go run ./examples/kks/create_cluster.go
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
	vpcKRN := "enter the vpckrn"
	subnetKRN := "enter the subnetkrn"
	client := krutrim.NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	region := "In-Bangalore-1" // Replace with your region.

	res, err := client.KKS.V1.Cluster.Create(ctx, krutrim.KKSV1ClusterCreateParams{
		Name:    "test",
		Version: "1.36",
		XRegion: region,
		ResourcesVPCConfig: krutrim.KKSResourcesVPCConfig{
			VPCKRN:                vpcKRN,
			SubnetKRNs:            subnetKRN,
			EndpointPrivateAccess: false,
		},
		KubernetesNetworkConfig: krutrim.KKSKubernetesNetworkConfig{
			PodIPv4CIDR:     "10.244.0.0/16",
			ServiceIPv4CIDR: "10.96.0.0/12",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(res); err != nil {
		log.Fatal(err)
	}
}
