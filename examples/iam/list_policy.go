package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	krutrim "github.com/ola-silicon/krutrim-go-sdk"
)

func main() {
	_ = godotenv.Load()



	iamClient := krutrim.NewIAMClient()
	ctx := context.Background()

	params := krutrim.ListPoliciesParams{
		Limit:          999,
		Offset:         0,
		KrutrimManaged: "all",
	}

	resp, err := iamClient.ListPolicies(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	printJSON(resp)
}

func printJSON(resp any) {
	fmt.Printf("%+v\n", resp)
}