//go:build ignore

// Transactions examples
//
// Demonstrates how to retrieve and list transaction history.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bilasdk/go"
	"github.com/bilasdk/go/option"
)

const (
	transactionID = "68f11209-451f-4a15-bfcd-d916eb8b09f4"
	accountID     = "68f11209-451f-4a15-bfcd-d916eb8b09f4"
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

	// Retrieve transaction
	transaction, err := client.Transactions.Get(ctx, transactionID)
	if err != nil {
		panic(err)
	}
	fmt.Println("retrieve:", transaction)

	// List transactions
	listParams := bila.TransactionListParams{
		AccountID: bila.String(accountID),
		StartDate: bila.String("2024-01-01T00:00:00Z"),
		EndDate:   bila.String("2024-12-31T23:59:59Z"),
		Page:      bila.Float(1),
		PerPage:   bila.Float(50),
		Type:      bila.TransactionListParamsTypeCredit,
	}
	transactions, err := client.Transactions.List(ctx, listParams)
	if err != nil {
		panic(err)
	}
	fmt.Println("list:", transactions)
}
