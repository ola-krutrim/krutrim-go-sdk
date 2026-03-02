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

	params := krutrim.CreateSubnetNewParams{

		VpcID: "enter the vpc_krn",

		SubnetData: krutrim.SubnetData{
			Name:        "enter the name",
			Description: "",
			CIDR:        "enter the cidr", // possible values 192.168.1.0/24
			GatewayIP:   "192.168.1.1",    //192.168.1.1
			NetworkID:   "enter the network_krn",
			IPVersion:   "4", //enter as per requirement
			Ingress:     true,
			Egress:      true,
		},
	}

	var resp *http.Response

	err := client.CreateSubnet.New(
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
