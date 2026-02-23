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
	krutrim "github.com/ola-silicon/krutrim-go-sdk"
	"github.com/ola-silicon/krutrim-go-sdk/option"
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
	params := krutrim.SshkeyNewParams{
		// REQUIRED FIELDS
		KeyName:   "Enter the SSH key name",
		PublicKey: "Enter the rsa Public SSH key",
		XRegion:   "Enter the Region",
	}

	var resp *http.Response

	// Call API
	err := client.Sshkey.New(
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

	var pretty bytes.Buffer
	_ = json.Indent(&pretty, bodyBytes, "", "  ")

	fmt.Printf("ssh created successfully: %+v\n", pretty.String())
}
