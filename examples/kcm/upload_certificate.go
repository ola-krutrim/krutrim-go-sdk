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
	"github.com/ola-krutrim/krutrim-go-sdk"
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

	// Path to your certificate file
	certFilePath := "/Users/divya.sonthalia/Downloads/bundle.p12"

	// Open the certificate file
	certFile, err := os.Open(certFilePath)
	if err != nil {
		log.Fatalf("Error opening certificate file: %v", err)
	}
	defer certFile.Close()

	// Set up the parameters to import/upload a certificate
	params := krutrim.KcmV1CertImportParams{
		CertFile: certFile,
		Name:     "sdk-cert1345",
		XVpcID:   "krn:vpc:In-Hyderabad-1:7344783839:2fbf2544-0c1d-481c-b745-4ef99e564b41:vpc:5891272d-60f5-4b8d-a11a-26a3af34c588",
		// Remove Flag and Tags if they're causing issues
	}

	// Call the Import method to upload the certificate
	var resp *http.Response

	// Call the Import method to upload the certificate
	err = client.Kcm.V1.Certs.Import(ctx, params, option.WithResponseInto(&resp))
	if err != nil {
		log.Fatalf("Error uploading certificate: %v", err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var pretty bytes.Buffer
	_ = json.Indent(&pretty, bodyBytes, "", "  ")

	fmt.Printf("certi uploaded successfully: %+v\n", pretty.String())

	// fmt.Printf("Certificate '%s' uploaded successfully! %s\n", params.Name, resp.Body)
}
