package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ola-silicon/krutrim-go-sdk"
)

func main() {

	client := krutrim.NewClient()
	ctx := context.Background()


	accountID := "enter the account id"
	email := "enter the email"
	password := "enter the password"



	resp, err := client.IAM.SignInAsIAMUser(
		ctx,
		krutrim.IAMUserSignInParams{
			Email:     email,
			Password:  password,
			AccountID: accountID,
		},
	)
	if err != nil {
		log.Fatal("IAM user login failed:", err)
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Printf("IAM User Login Successful:\n%s\n", string(data))
}
