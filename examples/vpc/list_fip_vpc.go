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

	params := krutrim.FloatingIPListListParams{
		XRegion: "Enter the Region : {eg: colo-1}",
		VpcID:   "Enter the VPC KRN ID",
	}

	var resp *http.Response

	err := client.FloatingIPList.List(
		ctx,
		params,
		option.WithResponseInto(&resp),
	)

	if err != nil {
		fmt.Println("API Error:", err)
		return
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}
