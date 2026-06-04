// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bila_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/bilasdk/go"
	"github.com/bilasdk/go/internal/testutil"
	"github.com/bilasdk/go/option"
)

func TestTransferGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bila.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Transfers.Get(context.TODO(), "68f11209-451f-4a15-bfcd-d916eb8b09f4")
	if err != nil {
		var apierr *bila.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransferListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bila.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Transfers.List(context.TODO(), bila.TransferListParams{
		AccountID: bila.String("68f11209-451f-4a15-bfcd-d916eb8b09f4"),
		EndDate:   bila.String("2024-12-31T23:59:59Z"),
		Page:      bila.Float(1),
		PerPage:   bila.Float(50),
		StartDate: bila.String("2024-01-01T00:00:00Z"),
		Status:    bila.TransferListParamsStatusPending,
		Type:      bila.TransferListParamsTypeBankAccount,
	})
	if err != nil {
		var apierr *bila.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransferGetStatusByReference(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bila.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Transfers.GetStatusByReference(context.TODO(), "transfer-001")
	if err != nil {
		var apierr *bila.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransferInitiateBankTransferWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bila.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Transfers.InitiateBankTransfer(context.TODO(), bila.TransferInitiateBankTransferParams{
		AccountID:           "68f11209-451f-4a15-bfcd-d916eb8b09f4",
		Amount:              1000,
		Reference:           "transfer-001",
		AccountNumber:       bila.String("1234567890"),
		BankID:              bila.String("bank-001"),
		Country:             bila.TransferInitiateBankTransferParamsCountryZm,
		Narration:           bila.String("Payment for services"),
		RecipientName:       bila.String("Jane Doe"),
		TransferRecipientID: bila.String("68f11209-451f-4a15-bfcd-d916eb8b09f4"),
		WalletID:            bila.String("68f11209-451f-4a15-bfcd-d916eb8b09f4"),
	})
	if err != nil {
		var apierr *bila.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransferInitiateMobileMoneyTransferWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bila.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Transfers.InitiateMobileMoneyTransfer(context.TODO(), bila.TransferInitiateMobileMoneyTransferParams{
		Amount:        250,
		Country:       bila.TransferInitiateMobileMoneyTransferParamsCountryZm,
		Operator:      bila.TransferInitiateMobileMoneyTransferParamsOperatorAirtel,
		Phone:         "0977433571",
		Reference:     "mobile-transfer-001",
		Narration:     bila.String("Mobile money payout"),
		RecipientName: bila.String("Jane Doe"),
		WalletID:      bila.String("68f11209-451f-4a15-bfcd-d916eb8b09f4"),
	})
	if err != nil {
		var apierr *bila.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
