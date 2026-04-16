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

	params := krutrim.GetTargetGroupNamesParams{
		XRegion: "enter the region",
		VpcID:   "enter the vpcid",
	}

	resp, err := client.HighlvlLoadBalancer.GetTargetGroupNames(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
