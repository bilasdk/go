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

func TestCollectionGet(t *testing.T) {
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
	_, err := client.Collections.Get(context.TODO(), "68f11209-451f-4a15-bfcd-d916eb8b09f4")
	if err != nil {
		var apierr *bila.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCollectionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Collections.List(context.TODO(), bila.CollectionListParams{
		AccountID: bila.String("68f11209-451f-4a15-bfcd-d916eb8b09f4"),
		EndDate:   bila.String("2024-12-31T23:59:59Z"),
		Page:      bila.Float(1),
		PerPage:   bila.Float(50),
		StartDate: bila.String("2024-01-01T00:00:00Z"),
		Status:    bila.CollectionListParamsStatusPending,
	})
	if err != nil {
		var apierr *bila.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCollectionGetStatusByReference(t *testing.T) {
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
	_, err := client.Collections.GetStatusByReference(context.TODO(), "collection-001")
	if err != nil {
		var apierr *bila.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCollectionInitiateMobileMoneyCollectionWithOptionalParams(t *testing.T) {
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
	_, err := client.Collections.InitiateMobileMoneyCollection(context.TODO(), bila.CollectionInitiateMobileMoneyCollectionParams{
		Amount:       100.5,
		Country:      bila.CollectionInitiateMobileMoneyCollectionParamsCountryZm,
		Operator:     bila.CollectionInitiateMobileMoneyCollectionParamsOperatorAirtel,
		Phone:        "0977433571",
		Reference:    "collection-001",
		WalletID:     "68f11209-451f-4a15-bfcd-d916eb8b09f4",
		Bearer:       bila.CollectionInitiateMobileMoneyCollectionParamsBearerCustomer,
		CustomerName: bila.String("John Doe"),
		Narration:    bila.String("Payment for subscription"),
	})
	if err != nil {
		var apierr *bila.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
