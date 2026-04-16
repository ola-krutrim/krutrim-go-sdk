package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	krutrim "github.com/ola-krutrim/krutrim-go-sdk"
)

func main() {
	_ = godotenv.Load()

	if os.Getenv("KRUTRIM_CLIENT_API_KEY") == "" {
		fmt.Println("API key missing")
		return
	}

	client := krutrim.NewClient()
	ctx := context.Background()

	bucketKRN := "enter bucket krn"
	objectKey := "enter the object key" //jammy-server-cloudimg-arm64.img
	localFile := "enter the local file path" // /home/ola/jammy-server-cloudimg-arm64.img

	xRegion := "enter the region"
	xSessionToken := "enter the session token"


	resp, err := client.Ko.V1.Objects.InitUpload(
		ctx,
		bucketKRN,
		objectKey,
		xRegion,
		xSessionToken,
	)
	if err != nil {
		fmt.Println("Init upload failed:", err)
		return
	}

	uploadURL := resp["uploadPreSignedUrl"]


	file, _ := os.Open(localFile)
	defer file.Close()

	req, _ := http.NewRequest(http.MethodPut, uploadURL, file)
	stat, _ := file.Stat()
	req.ContentLength = stat.Size()

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil || res.StatusCode >= 300 {
		body, _ := io.ReadAll(res.Body)
		fmt.Println("Upload failed:", string(body))
		return
	}

	fmt.Println("Object uploaded successfully")
}