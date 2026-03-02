package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	krutrim "github.com/ola-krutrim/krutrim-go-sdk"
	"github.com/ola-krutrim/krutrim-go-sdk/packages/param"
)

func main() {

	// Load .env file
	_ = godotenv.Load()

	if os.Getenv("KRUTRIM_CLIENT_API_KEY") == "" {
		fmt.Println("API key missing")
		return
	}

	// Create SDK client (auth picked from env internally)
	client := krutrim.NewClient()

	ctx := context.Background()

	params := krutrim.SecurityGroupV1NewRuleParams{
		// REQUIRED FIELDS
		Vpcid:          "Enter the VPC KRN",
		Direction:      krutrim.SecurityGroupV1NewRuleParamsDirectionEgress, //Select from the two hardcoded consts
		Ethertypes:     krutrim.SecurityGroupV1NewRuleParamsEthertypesIPv4,  //Select from the hardcoded consts
		Protocol:       "Enter the Network protocol",
		XRegion:        "Enter the Region",
		PortMaxRange:   8800, //Enter the Maximum port range
		PortMinRange:   80,   //Enter the Minimum port range
		RemoteIPPrefix: "Enter the Remote Ip address in format: '10.20.12.7/15' ",
		Description:    param.NewOpt("Enter the Description Of SGR"),
	}

	// Call API
	resp, err := client.SecurityGroup.V1.NewRule(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("SG Rule created successfully: %+v\n", resp)
	responseJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling response to JSON: %v", err)
	}
	fmt.Printf("SG Rule created successfully:\n%s\n", string(responseJSON))
}
