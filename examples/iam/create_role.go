package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	krutrim "github.com/ola-krutrim/krutrim-go-sdk"
)

func main() {
	_ = godotenv.Load()

	iamClient := krutrim.NewIAMClient()
	ctx := context.Background()

	params := krutrim.CreateRoleParams{
		PolicyIDs: []string{"policy_id_here"},
	}

	params.Role.Name = "enter the rolename"
	params.Role.Description = "enter the roledescription"

	resp, err := iamClient.CreateRoleWithPolicies(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	printJSON(resp)
}

func printJSON(resp any) {
	data, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(data))
}