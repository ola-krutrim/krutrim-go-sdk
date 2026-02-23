package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	krutrim "github.com/ola-silicon/krutrim-go-sdk"
	"github.com/ola-silicon/krutrim-go-sdk/option"
)

func main() {

	_ = godotenv.Load()

	if os.Getenv("KRUTRIM_CLIENT_API_KEY") == "" {
		fmt.Println("API key missing")
		return
	}

	client := krutrim.NewClient()

	ctx := context.Background()

	params := krutrim.SearchSubnetListParams{

		XRegion: "Enter the Region : {eg: colo-1}",

		VpcID: "Enter the VPC KRN ID",

		Page: 1,
		Size: 10,
	}

	var resp *http.Response

	err := client.SearchSubnet.List(
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
		fmt.Println("Read Error:", err)
		return
	}

	fmt.Println(string(body))
}
