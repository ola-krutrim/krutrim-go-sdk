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

	lbKrn := "enter the lbkrn"
	xRegion := "enter region"
	createPort := true

	params := krutrim.UpdateLoadBalancerParams{
		XRegion: xRegion,
		LBKrn:   lbKrn,

		SecurityGroups: []string{
			"",
		},

		LoadBalancerData: &krutrim.UpdateLoadBalancerData{
			LBName:      "enter the name",
			Description: "",
			CreatePort:  &createPort,

			VpcID:       "enter the vpc id",
			NetworkID:   "enter the network krn",
			VipSubnetID: "enter the subnetid",
		},

		Listeners: []krutrim.UpdateListener{
			{
				ListenerData: krutrim.UpdateListenerData{
					ListenerID:    "",
					ListenerIndex: 1,
					ListenerName:  "fgb",
					Protocol:      "TCP",
					ProtocolPort:  80,
				},

				PoolData: []krutrim.UpdatePoolData{
					{
						PoolID:       "",
						PoolName:     "sdfxsdfgb",
						Protocol:     "TCP",
						LBAlgorithm:  "ROUND_ROBIN",
						AdminStateUp: true,

						TargetGroupName: "xc",

						HealthMonitor: &krutrim.UpdateHealthMonitor{
							HealthMonitorID: "",
							HType:           "HTTP", // ⚠️ matches curl (even though TCP LB)
							Delay:           30,
							Timeout:         5,
							MaxRetries:      3,
							URLPath:         "/",
							Name:            "enter the name",
						},

						MemberData: []krutrim.UpdateMember{
							{
								KRN:         "",
								MemberID:    "",
								MemberIndex: 1,
								Status:      "success",
							},
						},
					},
				},

			},
		},
	}



	resp, err := client.HighlvlLoadBalancer.UpdateLoadBalancer(
		ctx,
		lbKrn,
		params,
	)
	if err != nil {
		fmt.Println("API Error:", err)
		return
	}

	resJSON, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println("Response:\n", string(resJSON))
}