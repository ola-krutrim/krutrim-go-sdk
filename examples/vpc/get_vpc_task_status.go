package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	krutrim "github.com/ola-krutrim/krutrim-go-sdk"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
)

func main() {

	_ = godotenv.Load()

	if os.Getenv("KRUTRIM_CLIENT_API_KEY") == "" {
		fmt.Println("API key missing")
		return
	}

	client := krutrim.NewClient()

	ctx := context.Background()

	params := krutrim.GetVpcTaskStatusNewParams{
		TaskID:  "Enter the Task ID: {eg :dee31451-d66e-4a41-984d-16e4aa80c503}",
		XRegion: "Enter the Region : {eg: colo-1}",
	}

	var resp *http.Response

	err := client.GetVpcTaskStatus.New(
		ctx,
		params,
		option.WithResponseInto(&resp),
	)

	if err != nil {
		fmt.Println("API Error:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Error:", err)
		return
	}

	fmt.Println(string(body))
}
