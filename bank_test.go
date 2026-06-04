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

func TestBankListWithOptionalParams(t *testing.T) {
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
	_, err := client.Banks.List(context.TODO(), bila.BankListParams{
		Country: bila.String("zm"),
	})
	if err != nil {
		var apierr *bila.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
