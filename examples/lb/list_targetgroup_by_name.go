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

	params := krutrim.GetTargetGroupParams{
		XRegion:         "enter the region",
		VpcID:           "enter the vpcid",
		TargetGroupName: "enter the target group name",
	}

	resp, err := client.HighlvlLoadBalancer.GetTargetGroup(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
