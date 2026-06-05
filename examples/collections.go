//go:build ignore

// Collections examples
//
// Demonstrates how to collect payments via mobile money.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bilasdk/go"
	"github.com/bilasdk/go/option"
)

const (
	collectionID = "68f11209-451f-4a15-bfcd-d916eb8b09f4"
	walletID     = "68f11209-451f-4a15-bfcd-d916eb8b09f4"
	reference    = "collection-001"
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

	// Retrieve collection
	collection, err := client.Collections.Get(ctx, collectionID)
	if err != nil {
		panic(err)
	}
	fmt.Println("retrieve:", collection)

	// List collections
	listParams := bila.CollectionListParams{
		AccountID: bila.String(walletID),
		StartDate: bila.String("2024-01-01T00:00:00Z"),
		EndDate:   bila.String("2024-12-31T23:59:59Z"),
		Page:      bila.Float(1),
		PerPage:   bila.Float(50),
		Status:    bila.CollectionListParamsStatusPending,
	}
	collections, err := client.Collections.List(ctx, listParams)
	if err != nil {
		panic(err)
	}
	fmt.Println("list:", collections)

	// Get collection status by reference
	status, err := client.Collections.GetStatusByReference(ctx, reference)
	if err != nil {
		panic(err)
	}
	fmt.Println("getStatusByReference:", status)

	// Initiate mobile money collection
	initiateParams := bila.CollectionInitiateMobileMoneyCollectionParams{
		Amount:       100.5,
		Country:      bila.CollectionInitiateMobileMoneyCollectionParamsCountryZm,
		Operator:     bila.CollectionInitiateMobileMoneyCollectionParamsOperatorAirtel,
		Phone:        "0977433571",
		Reference:    reference,
		WalletID:     walletID,
		Bearer:       bila.CollectionInitiateMobileMoneyCollectionParamsBearerCustomer,
		CustomerName: bila.String("John Doe"),
		Narration:    bila.String("Payment for subscription"),
	}
	initiated, err := client.Collections.InitiateMobileMoneyCollection(ctx, initiateParams)
	if err != nil {
		panic(err)
	}
	fmt.Println("initiateMobileMoneyCollection:", initiated)
}
