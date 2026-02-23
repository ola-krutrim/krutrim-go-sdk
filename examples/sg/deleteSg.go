package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	krutrim "github.com/ola-silicon/krutrim-go-sdk"
)

func main() {

	// Load .env file
	_ = godotenv.Load()

	if os.Getenv("KRUTRIM_CLIENT_API_KEY") == "" {
		fmt.Println("API key missing")
		return
	}

	// Create SDK client (auth picked from env internally)
	client := krutrim.NewClient()

	ctx := context.Background()

	params := krutrim.SecurityGroupV1DeleteParams{
		XRegion: "Enter the Region",
	}

	securityGroupKrn := "Enter the SG KRN Id"

	// Call API
	err := client.SecurityGroup.V1.Delete(ctx, securityGroupKrn, params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("SG deleted successfully\n")
}
