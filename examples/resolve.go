//go:build ignore

// Resolve examples
//
// Demonstrates how to verify bank account and mobile money account details.
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

	// Resolve bank account
	resolveBankParams := bila.ResolveBankAccountParams{
		AccountNumber: "1234567890",
		BankID:        "bank-001",
		Country:       bila.ResolveBankAccountParamsCountryZm,
	}
	bankAccount, err := client.Resolve.BankAccount(ctx, resolveBankParams)
	if err != nil {
		panic(err)
	}
	fmt.Println("bankAccount:", bankAccount)

	// Resolve mobile money
	resolveMobileParams := bila.ResolveMobileMoneyParams{
		Country:  bila.ResolveMobileMoneyParamsCountryZm,
		Operator: bila.ResolveMobileMoneyParamsOperatorAirtel,
		Phone:    "0977433571",
	}
	mobileMoney, err := client.Resolve.MobileMoney(ctx, resolveMobileParams)
	if err != nil {
		panic(err)
	}
	fmt.Println("mobileMoney:", mobileMoney)
}
