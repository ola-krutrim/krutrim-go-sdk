package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	krutrim "github.com/ola-krutrim/krutrim-go-sdk"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
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

	// ─────────────────────────────
	// REQUEST PARAMS
	// ─────────────────────────────
	params := krutrim.SshkeySearchParams{
		// REQUIRED FIELDS
		KeyName: "Enter the KEY name or prefix", //eg: testSS, testSSH
	}

	var resp *http.Response

	// Call API
	err := client.Sshkey.Search(
		ctx,
		params,
		option.WithResponseInto(&resp),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if len(bodyBytes) == 0 {
		fmt.Println("No SSH found for this VPC")
		return
	}

	var pretty bytes.Buffer
	_ = json.Indent(&pretty, bodyBytes, "", "  ")

	fmt.Printf("List of fetched SSH:\n%s\n", pretty.String())

}
