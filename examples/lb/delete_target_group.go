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

	vpcID := "enter the vpcID"
	targetGroupName := "enter the target group name"
	xRegion := "enter the region"

	resp, err := client.HighlvlLoadBalancer.DeleteTargetGroup(
		ctx,
		vpcID,
		targetGroupName,
		xRegion,
	)
	if err != nil {
		fmt.Println("API Error:", err)
		return
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}