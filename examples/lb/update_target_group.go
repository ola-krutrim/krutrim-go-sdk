package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	krutrim "github.com/ola-krutrim/krutrim-go-sdk"
)

func main() {
	client := krutrim.NewClient()
	ctx := context.Background()

	params := krutrim.UpdateTargetGroupParams{
		XRegion:         "enter the region",
		VpcID:           "enter the vpcid",
		TargetGroupName: "enter the target group name",
		Members: []krutrim.TargetGroupMember{
			{
				Name:         "enter the name",
				Address:      "enter the address",                //10.0.0.20
				ProtocolPort: 90,
				Weight:       1,
			},
		},
		HealthMonitor: krutrim.HealthMonitor{
			Name:       "enter the name",
			HType:      "enter the htype",
			Delay:      30,
			Timeout:    5,
			MaxRetries: 3,
			URLPath:    "/",
		},
	}

	resp, err := client.HighlvlLoadBalancer.UpdateTargetGroup(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
