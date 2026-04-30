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

	params := krutrim.GetUserParams{
		UserKRN: "enter the user_krn",
	}

	resp, err := iamClient.GetUser(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	printJSON(resp)
}

func printJSON(resp any) {
	data, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(data))
}