// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bila_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/bila-go"
	"github.com/stainless-sdks/bila-go/internal/testutil"
	"github.com/stainless-sdks/bila-go/option"
)

func TestTransactionGet(t *testing.T) {
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
	_, err := client.Transactions.Get(context.TODO(), "68f11209-451f-4a15-bfcd-d916eb8b09f4")
	if err != nil {
		var apierr *bila.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.List(context.TODO(), bila.TransactionListParams{
		AccountID: bila.String("68f11209-451f-4a15-bfcd-d916eb8b09f4"),
		EndDate:   bila.String("2024-12-31T23:59:59Z"),
		Page:      bila.Float(1),
		PerPage:   bila.Float(50),
		StartDate: bila.String("2024-01-01T00:00:00Z"),
		Type:      bila.TransactionListParamsTypeCredit,
	})
	if err != nil {
		var apierr *bila.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
