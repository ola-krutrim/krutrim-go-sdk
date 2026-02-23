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

	params := krutrim.CreateVpcAsyncNewParams{

		XRegion: "enter the region",

		Vpc: krutrim.Vpc{
			Name:        "enter the name",
			Description: "enter the description",
			Enabled:     true,
		},

		Network: krutrim.Network{
			Name:         "enter the network",
			AdminStateUp: true,
		},

		Subnet: krutrim.Subnet{
			Name:        "enter the subnetname",
			Description: "enter the description",
			CIDR:        "enter cidr",      //eg.192.168.2.0/24
			GatewayIP:   "enter gatewayip", //192.168.2.1
			IPVersion:   "4",               //enter as per requirement
			Ingress:     true,
			Egress:      true,
		},
	}

	var resp *http.Response

	err := client.CreateVpcAsync.New(
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
