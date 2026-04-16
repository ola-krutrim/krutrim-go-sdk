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

	taskID := "enter the taskid"
	xRegion := "enter the region"

	parsedResp, err := client.HighlvlLoadBalancer.GetTaskStatus(
		ctx,
		taskID,
		xRegion,
	)
	if err != nil {
		fmt.Println("API Error:", err)
		return
	}

	b, err := json.MarshalIndent(parsedResp, "", "  ")
	if err != nil {
		fmt.Println("JSON Marshal Error:", err)
		return
	}

	fmt.Println(string(b))
}
