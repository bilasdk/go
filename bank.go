// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bila

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/bila-go/internal/apijson"
	"github.com/stainless-sdks/bila-go/internal/apiquery"
	"github.com/stainless-sdks/bila-go/internal/requestconfig"
	"github.com/stainless-sdks/bila-go/option"
	"github.com/stainless-sdks/bila-go/packages/param"
	"github.com/stainless-sdks/bila-go/packages/respjson"
)

// Bank reference data endpoints
//
// BankService contains methods and other services that help with interacting with
// the bila API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBankService] method instead.
type BankService struct {
	options []option.RequestOption
}

// NewBankService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBankService(opts ...option.RequestOption) (r BankService) {
	r = BankService{}
	r.options = opts
	return
}

// Retrieve a list of all supported banks and financial institutions
func (r *BankService) List(ctx context.Context, query BankListParams, opts ...option.RequestOption) (res *BankListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/bila/banks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type BankListResponse struct {
	Data []BankListResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r BankListResponse) RawJSON() string { return r.JSON.raw }
func (r *BankListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BankListResponseData struct {
	// Bank ID
	ID string `json:"id" api:"required"`
	// Bank code
	Code string `json:"code" api:"required"`
	// Country code
	Country string `json:"country" api:"required"`
	// Bank name
	Name string `json:"name" api:"required"`
	// Bank type
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Code        respjson.Field
		Country     respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BankListResponseData) RawJSON() string { return r.JSON.raw }
func (r *BankListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BankListParams struct {
	// Filter banks by country code
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BankListParams]'s query parameters as `url.Values`.
func (r BankListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
