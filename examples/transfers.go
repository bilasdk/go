//go:build ignore

// Transfers examples
//
// Demonstrates how to send payouts via bank transfer and mobile money.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bilasdk/go"
	"github.com/bilasdk/go/option"
)

const (
	transferID          = "68f11209-451f-4a15-bfcd-d916eb8b09f4"
	accountID           = "68f11209-451f-4a15-bfcd-d916eb8b09f4"
	transferRecipientID = "68f11209-451f-4a15-bfcd-d916eb8b09f4"
	walletID            = "68f11209-451f-4a15-bfcd-d916eb8b09f4"
	bankReference       = "transfer-001"
	mobileReference     = "mobile-transfer-001"
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

	// Retrieve transfer
	transfer, err := client.Transfers.Get(ctx, transferID)
	if err != nil {
		panic(err)
	}
	fmt.Println("retrieve:", transfer)

	// List transfers
	listParams := bila.TransferListParams{
		AccountID: bila.String(accountID),
		StartDate: bila.String("2024-01-01T00:00:00Z"),
		EndDate:   bila.String("2024-12-31T23:59:59Z"),
		Page:      bila.Float(1),
		PerPage:   bila.Float(50),
		Status:    bila.TransferListParamsStatusPending,
		Type:      bila.TransferListParamsTypeBankAccount,
	}
	transfers, err := client.Transfers.List(ctx, listParams)
	if err != nil {
		panic(err)
	}
	fmt.Println("list:", transfers)

	// Get transfer status by reference
	status, err := client.Transfers.GetStatusByReference(ctx, bankReference)
	if err != nil {
		panic(err)
	}
	fmt.Println("getStatusByReference:", status)

	// Initiate bank transfer
	bankParams := bila.TransferInitiateBankTransferParams{
		AccountID:           accountID,
		Amount:              1000,
		Reference:           bankReference,
		AccountNumber:       bila.String("1234567890"),
		BankID:              bila.String("bank-001"),
		Country:             bila.TransferInitiateBankTransferParamsCountryZm,
		Narration:           bila.String("Payment for services"),
		RecipientName:       bila.String("Jane Doe"),
		TransferRecipientID: bila.String(transferRecipientID),
		WalletID:            bila.String(walletID),
	}
	bankTransfer, err := client.Transfers.InitiateBankTransfer(ctx, bankParams)
	if err != nil {
		panic(err)
	}
	fmt.Println("initiateBankTransfer:", bankTransfer)

	// Initiate mobile money transfer
	mobileParams := bila.TransferInitiateMobileMoneyTransferParams{
		Amount:        250,
		Country:       bila.TransferInitiateMobileMoneyTransferParamsCountryZm,
		Operator:      bila.TransferInitiateMobileMoneyTransferParamsOperatorAirtel,
		Phone:         "0977433571",
		Reference:     mobileReference,
		Narration:     bila.String("Mobile money payout"),
		RecipientName: bila.String("Jane Doe"),
		WalletID:      bila.String(walletID),
	}
	mobileTransfer, err := client.Transfers.InitiateMobileMoneyTransfer(ctx, mobileParams)
	if err != nil {
		panic(err)
	}
	fmt.Println("initiateMobileMoneyTransfer:", mobileTransfer)
}
