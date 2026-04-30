package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	krutrim "github.com/ola-silicon/krutrim-go-sdk"
)

func main() {
	_ = godotenv.Load()

	iamClient := krutrim.NewIAMClient()
	ctx := context.Background()

	params := krutrim.ProgrammaticSigninParams{
		AccountID: "enter the accountid",
		AccessKey: "enter the accesskey",
		SecretKey: "enter the secretkey",
	}

	resp, err := iamClient.SigninProgrammaticUser(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	printJSON(resp)
}

func printJSON(resp any) {
	data, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(data))
}