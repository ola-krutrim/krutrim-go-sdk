package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ola-silicon/krutrim-go-sdk"
)

func main() {

	client := krutrim.NewClient()
	ctx := context.Background()

	email := "enter the email"
	password := "enter the password"

	resp, err := client.IAM.SignInAsRootUser(
		ctx,
		krutrim.IAMSignInParams{
			Email:    email,
			Password: password,
		},
	)
	if err != nil {
		log.Fatal("Root login failed:", err)
	}

	token := resp.AccessToken

	content := fmt.Sprintf(`#!/usr/bin/env bash
export KRUTRIM_CLIENT_API_KEY="%s"
`, token)

	err = os.WriteFile("env.sh", []byte(content), 0600)
	if err != nil {
		log.Fatal("Failed to write env.sh:", err)
	}

	fmt.Println("s Token saved to env.sh")

	// Print response
	data, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Printf("Root User Login Successful:\n%s\n", string(data))
}
