//go:build ignore

// Accounts examples
//
// Demonstrates how to retrieve accounts, list accounts, and check balances.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bilasdk/go"
	"github.com/bilasdk/go/option"
)

const accountID = "68f11209-451f-4a15-bfcd-d916eb8b09f4"

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

	// Retrieve account
	account, err := client.Accounts.Get(ctx, accountID)
	if err != nil {
		panic(err)
	}
	fmt.Println("retrieve:", account)

	// List accounts
	listParams := bila.AccountListParams{
		Page:    bila.Float(1),
		PerPage: bila.Float(50),
	}
	accounts, err := client.Accounts.List(ctx, listParams)
	if err != nil {
		panic(err)
	}
	fmt.Println("list:", accounts)

	// Get account balance
	balance, err := client.Accounts.GetBalance(ctx, accountID)
	if err != nil {
		panic(err)
	}
	fmt.Println("getBalance:", balance)
}
