package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	krutrim "github.com/ola-krutrim/krutrim-go-sdk"
)

func main() {
	_ = godotenv.Load()

	if os.Getenv("KRUTRIM_CLIENT_API_KEY") == "" {
		log.Fatal("API key missing")
	}

	client := krutrim.NewClient()
	ctx := context.Background()
	//enter as per requirement
	params := krutrim.CreateTargetGroupParams{
		XRegion: "enter the region",

		VpcID:           "enter the vpckrn",
		TargetGroupName: "enter the target group name",
		Members: []krutrim.TargetGroupMember{
			{
				Name:         "enter the name",
				Address:      "enter the address", //10.0.38.52
				ProtocolPort: 80,
				Weight:       1,
			},
		},
		//enter as per requirement
		HealthMonitor: krutrim.HealthMonitor{
			Name:       "enter the name",
			HType:      "enter the htype", //HTTP
			Delay:      30,
			Timeout:    5,
			MaxRetries: 3,
			URLPath:    "/",
		},
	}

	resp, err := client.HighlvlLoadBalancer.CreateTargetGroup(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status:", resp.Status)

}
