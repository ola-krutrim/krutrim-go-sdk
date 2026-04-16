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

	params := krutrim.CreateLoadBalancerParams{

		XRegion: "enter the region",

		LoadBalancerData: krutrim.LoadBalancerData{
			LBName:      "enter the name",
			CreatePort:  true,
			FloatingIP:  false,
			VpcID:       "enter the vpc id",
			NetworkID:   "enter the network id",
			VipSubnetID: "enter the subneid",
			LBType:      "ALB",
			Flavor:      "standard",
		},

		SecurityGroups: []string{
			"enter the sg",
		},

		ListenerCount: 1,

		Listeners: []krutrim.Listener{
			{
				ListenerData: krutrim.ListenerData{
					Name:         "enter the name",
					ListenerName: "ente the listener name",
					Protocol:     "enter the protocol",
					ProtocolPort: 443,
					DefaultPool:  true,
					//remove the optinal field for http protocol
					DefaultTLSContainerRef: "enter the cert krn",
					SNIContainerRefs: []string{
						"enter the cert krn",
					},
				},

				PoolData: []krutrim.PoolData{
					{
						PoolName:        "enter the pool name",
						Protocol:        "enter the protocol",
						LBAlgorithm:     "ROUND_ROBIN",
						AdminStateUp:    true,
						TargetGroupName: "xc",
						F5:              false,
					},
				},

				PolicyData: []krutrim.PolicyData{
					{
						PolicyName:       "enter the policy name",
						Action:           "enter the action",
						RedirectPoolName: "enter the pool name",
						RuleData: []krutrim.RuleData{
							{
								CompareType: "EQUAL_TO",
								Type:        "PATH",
								Value:       "/test",
							},
						},
					},
				},
			},
		},
	}

	res, err := client.HighlvlLoadBalancer.CreateLoadBalancer(ctx, params)
	if err != nil {
		fmt.Println("API Error:", err)
		return
	}

	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println("JSON Marshal Error:", err)
		return
	}

	fmt.Println("RESPONSE:")
	fmt.Println(string(b))
}
