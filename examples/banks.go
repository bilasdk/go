//go:build ignore

// Banks examples
//
// Demonstrates how to list supported banks and financial institutions.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bilasdk/go"
	"github.com/bilasdk/go/option"
)

func main() {
	apiKey := os.Getenv("BILA_API_KEY")
	if apiKey == "" {
		apiKey = "sk_test_your_api_key_here"
	}

	client := bila.NewClient(
		option.WithAPIKey(apiKey),
		option.WithEnvironmentSandbox(),
	)
	ctx := context.Background()

	listParams := bila.BankListParams{
		Country: bila.String("zm"),
	}
	banks, err := client.Banks.List(ctx, listParams)
	if err != nil {
		panic(err)
	}
	fmt.Println("list:", banks)
}
