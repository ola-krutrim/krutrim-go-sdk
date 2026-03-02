package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	krutrim "github.com/ola-krutrim/krutrim-go-sdk"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
)

func main() {
	_ = godotenv.Load()

	if os.Getenv("KRUTRIM_CLIENT_API_KEY") == "" {
		fmt.Println("API key missing")
		return
	}
	client := krutrim.NewClient()
	ctx := context.Background()

	params := krutrim.DescribeVpcGetParams{
		VpcID:   "Enter the VPC KRN ID",
		VpcName: "Enter the VPC name",
		XRegion: "Enter the Region : colo-1/colo-2",
	}

	var resp *http.Response

	err := client.DescribeVpc.Get(
		ctx,
		params,
		option.WithResponseInto(&resp),
	)
	if err != nil {
		fmt.Println("API Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Println(string(body))
}
