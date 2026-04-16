package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	krutrim "github.com/ola-krutrim/krutrim-go-sdk"
)

func main() {
	client := krutrim.NewClient()
	ctx := context.Background()

	params := krutrim.GetTargetGroupListParams{
		XRegion: "enter the region",
		VpcID:   "enter the vpcid",
	}

	resp, err := client.HighlvlLoadBalancer.GetTargetGroupList(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

