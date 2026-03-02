package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ola-krutrim/krutrim-go-sdk"
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


	isGPU := false
	floatingIP := true
	deleteOnTermination := true
	volumeSize := int64(20)					//enter as per requirement

	imageKrn := "enter the image krn"
	sshKey := "enter the ssh key name"

	userData := "enter the user data"


	params := krutrim.HighlvlvpcNewInstanceParams{
		// REQUIRED
		InstanceName: "enter the instance name",
		InstanceType: "enter the instancetype",                      //CPU-2x-8GB
		Region:       "enter the region",                              //In-Bangalore-1

		VpcID:      "enter the vpc krn",
		NetworkID:  "enter the network krn",
		SubnetID:   "enter the subnet krn",
		Volumetype: "enter the volumetype",										//HNSS
		VolumeName: "enter the volume name",

		SecurityGroups: []string{
			"enter the security group krn",
		},

		IsGPU:               &isGPU,
		FloatingIP:          &floatingIP,
		DeleteOnTermination: &deleteOnTermination,
		VolumeSize:          &volumeSize,
		ImageKrn:            &imageKrn,
		SshkeyName:          &sshKey,
		UserData:            &userData,

		Tags:    []krutrim.HighlvlvpcNewInstanceParamsTag{},
		Volumes: []any{},
	}


	resp, err := client.Highlvlvpc.NewInstance(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	responseJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling response: %v", err)
	}

	fmt.Printf("Instance created successfully:\n%s\n", string(responseJSON))
}
