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
	SshKeyID := "Enter the SSH key ID not KRN" //eg: 2fc91107-b557-4d8a-984a-f8f8fb4888b7

	var resp *http.Response

	// Call API
	err := client.Sshkey.Delete(
		ctx,
		SshKeyID,
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

	fmt.Printf("%+v\n", pretty.String())
}
