package main

import (
	"context"
	"encoding/json"
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

	iamClient := krutrim.NewIAMClient()
	ctx := context.Background()

	params := krutrim.CreateUserParams{}

	params.User.UserName = "enter the user_name"
	params.User.Email = "enter the mail"
	params.User.Password = "enter the password"
	params.User.ConsoleAccess = true

	resp, err := iamClient.CreateUser(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	printJSON(resp)
}

func printJSON(resp any) {
	data, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(data))
}