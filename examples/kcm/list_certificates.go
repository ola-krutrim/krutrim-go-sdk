package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/ola-krutrim/krutrim-go-sdk"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
	"github.com/ola-krutrim/krutrim-go-sdk/packages/param"
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

	// Optional: Filter by VPC ID to list certificates in a specific VPC
	vpcID := "krn:vpc:In-Hyderabad-1:7344783839:2fbf2544-0c1d-481c-b745-4ef99e564b41:vpc:5891272d-60f5-4b8d-a11a-26a3af34c588"

	// Set up the parameters to list certificates
	params := krutrim.KcmV1CertListParams{
		VpcID: param.NewOpt(vpcID),
	}

	var bodyBytes []byte
	var resp *http.Response

	// Call the List method to retrieve all certificates
	_, err := client.Kcm.V1.Certs.List(ctx, params,
		option.WithResponseInto(&resp),
		option.WithResponseBodyInto(&bodyBytes),
	)
	if err != nil {
		log.Fatalf("Error listing certificates: %v", err)
	}

	// Check if response is empty
	if len(bodyBytes) == 0 {
		fmt.Println("No certificates found in this VPC")
		return
	}

	// Parse the response to see the structure
	var certResponse map[string]interface{}
	err = json.Unmarshal(bodyBytes, &certResponse)
	if err != nil {
		log.Fatalf("Error parsing response JSON: %v", err)
	}

	// Marshal to pretty JSON
	certsJSON, err := json.MarshalIndent(certResponse, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling certificates to JSON: %v", err)
	}

	// Print the certificates list
	fmt.Printf("Certificates retrieved successfully:\n%s\n", string(certsJSON))

	// Additional helpful output - try to parse certificate data
	if certs, ok := certResponse["certificates"].([]interface{}); ok && len(certs) > 0 {
		fmt.Printf("\n✓ Total certificates found: %d\n", len(certs))
		fmt.Println("\nCertificate Summary:")
		for i, certData := range certs {
			if cert, ok := certData.(map[string]interface{}); ok {
				commonName := ""
				krn := ""
				expiration := ""

				if cn, ok := cert["commonName"].(string); ok {
					commonName = cn
				}
				if k, ok := cert["krn"].(string); ok {
					krn = k
				}
				if exp, ok := cert["expiration"].(string); ok {
					expiration = exp
				}

				fmt.Printf("%d. %s\n", i+1, commonName)
				fmt.Printf("   KRN: %s\n", krn)
				fmt.Printf("   Expiration: %s\n", expiration)
			}
		}
	} else if data, ok := certResponse["data"].([]interface{}); ok && len(data) > 0 {
		fmt.Printf("\n✓ Total certificates found: %d\n", len(data))
		fmt.Println("\nCertificate Summary:")
		for i, certData := range data {
			if cert, ok := certData.(map[string]interface{}); ok {
				commonName := ""
				krn := ""
				expiration := ""

				if cn, ok := cert["commonName"].(string); ok {
					commonName = cn
				}
				if k, ok := cert["krn"].(string); ok {
					krn = k
				}
				if exp, ok := cert["expiration"].(string); ok {
					expiration = exp
				}

				fmt.Printf("%d. %s\n", i+1, commonName)
				fmt.Printf("   KRN: %s\n", krn)
				fmt.Printf("   Expiration: %s\n", expiration)
			}
		}
	} else {
		fmt.Println("\n⚠ No certificates array found in the expected fields.")
		fmt.Println("Check the JSON output above to see the actual structure.")
	}
}
