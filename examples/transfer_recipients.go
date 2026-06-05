//go:build ignore

// Transfer recipients examples
//
// Demonstrates how to manage payout recipients for bank accounts and mobile money.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bilasdk/go"
	"github.com/bilasdk/go/option"
)

const recipientID = "68f11209-451f-4a15-bfcd-d916eb8b09f4"

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

	// Retrieve transfer recipient
	recipient, err := client.TransferRecipients.Get(ctx, recipientID)
	if err != nil {
		panic(err)
	}
	fmt.Println("retrieve:", recipient)

	// List transfer recipients
	listParams := bila.TransferRecipientListParams{
		Page:    bila.Float(1),
		PerPage: bila.Float(50),
		Type:    bila.TransferRecipientListParamsTypeBankAccount,
	}
	recipients, err := client.TransferRecipients.List(ctx, listParams)
	if err != nil {
		panic(err)
	}
	fmt.Println("list:", recipients)

	// Create bank account recipient
	bankParams := bila.TransferRecipientNewBankAccountParams{
		AccountNumber: "1234567890",
		BankID:        "bank-001",
		AccountName:   bila.String("John Doe"),
		Country:       bila.TransferRecipientNewBankAccountParamsCountryZm,
	}
	bankRecipient, err := client.TransferRecipients.NewBankAccount(ctx, bankParams)
	if err != nil {
		panic(err)
	}
	fmt.Println("createBankAccount:", bankRecipient)

	// Create mobile money recipient
	mobileParams := bila.TransferRecipientNewMobileMoneyParams{
		Country:     bila.TransferRecipientNewMobileMoneyParamsCountryZm,
		Operator:    bila.TransferRecipientNewMobileMoneyParamsOperatorAirtel,
		Phone:       "0977433571",
		AccountName: bila.String("John Doe"),
	}
	mobileRecipient, err := client.TransferRecipients.NewMobileMoney(ctx, mobileParams)
	if err != nil {
		panic(err)
	}
	fmt.Println("createMobileMoney:", mobileRecipient)
}
