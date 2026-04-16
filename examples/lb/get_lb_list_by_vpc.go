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


	vpcKrn := "enter the vpckrn"
	region := "enter the region"

	page  := int64(1)  
	limit := int64(10)


	resp, err := client.HighlvlLoadBalancer.GetLoadBalancerList(
		ctx,
		vpcKrn,
		page,
		limit,
		region,
	)
	if err != nil {
		fmt.Println("API Error:", err)
		return
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}