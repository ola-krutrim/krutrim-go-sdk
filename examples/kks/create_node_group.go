//go:build ignore

// Run with: go run ./examples/kks/create_node_group.go
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
	clusterKRN := "enter the clusterkrn"
	subnetKRN := "enter the subnetkrn"
	client := krutrim.NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	region := "enter the region"

	res, err := client.KKS.V1.Cluster.NodeGroup.Create(ctx, clusterKRN, krutrim.KKSV1ClusterNodeGroupCreateParams{
		Name:             "enter the nodegroup name",
		InstanceTypes:    "enter the instance type",
		DiskSize:         166,
		ScalingConfig:    krutrim.KKSScalingConfig{MinSize: 1, MaxSize: 5, DesiredSize: 2},
		SubnetsKRN:       subnetKRN,
		Labels:           map[string]string{},
		Taints:           []krutrim.KKSTaint{},
		RemoteAccess:     &krutrim.KKSRemoteAccess{SourceSecurityGroupsKRNs: []string{}},
		NodeRepairConfig: &krutrim.KKSNodeRepairConfig{Enabled: true},
		XRegion:          region,
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
