// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bila

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/bila-go/internal/apijson"
	"github.com/stainless-sdks/bila-go/internal/apiquery"
	"github.com/stainless-sdks/bila-go/internal/requestconfig"
	"github.com/stainless-sdks/bila-go/option"
	"github.com/stainless-sdks/bila-go/packages/param"
	"github.com/stainless-sdks/bila-go/packages/respjson"
)

// Transaction history endpoints
//
// TransactionService contains methods and other services that help with
// interacting with the bila API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionService] method instead.
type TransactionService struct {
	options []option.RequestOption
}

// NewTransactionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionService(opts ...option.RequestOption) (r TransactionService) {
	r = TransactionService{}
	r.options = opts
	return
}

// Retrieve a single transaction by its UUID
func (r *TransactionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TransactionGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/bila/transactions/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve a paginated list of transactions
func (r *TransactionService) List(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) (res *TransactionListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/bila/transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type TransactionGetResponse struct {
	Data TransactionGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r TransactionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TransactionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionGetResponseData struct {
	// Transaction UUID
	ID string `json:"id" api:"required"`
	// Account / wallet ID
	AccountID string `json:"accountId" api:"required"`
	// Transaction amount
	Amount float64 `json:"amount" api:"required"`
	// Balance after transaction
	BalanceAfter float64 `json:"balanceAfter" api:"required"`
	// Balance before transaction
	BalanceBefore float64 `json:"balanceBefore" api:"required"`
	// Transaction timestamp
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Currency code
	Currency string `json:"currency" api:"required"`
	// Transaction status
	//
	// Any of "pending", "successful", "failed", "cancelled".
	Status string `json:"status" api:"required"`
	// Transaction type
	//
	// Any of "credit", "debit".
	Type string `json:"type" api:"required"`
	// Transaction description
	Description string `json:"description"`
	// Client reference
	Reference string `json:"reference"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AccountID     respjson.Field
		Amount        respjson.Field
		BalanceAfter  respjson.Field
		BalanceBefore respjson.Field
		CreatedAt     respjson.Field
		Currency      respjson.Field
		Status        respjson.Field
		Type          respjson.Field
		Description   respjson.Field
		Reference     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *TransactionGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionListResponse struct {
	Data TransactionListResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r TransactionListResponse) RawJSON() string { return r.JSON.raw }
func (r *TransactionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionListResponseData struct {
	// List of transactions
	Data []TransactionListResponseDataData `json:"data" api:"required"`
	// Pagination metadata
	Meta TransactionListResponseDataMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionListResponseData) RawJSON() string { return r.JSON.raw }
func (r *TransactionListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionListResponseDataData struct {
	// Transaction UUID
	ID string `json:"id" api:"required"`
	// Account / wallet ID
	AccountID string `json:"accountId" api:"required"`
	// Transaction amount
	Amount float64 `json:"amount" api:"required"`
	// Balance after transaction
	BalanceAfter float64 `json:"balanceAfter" api:"required"`
	// Balance before transaction
	BalanceBefore float64 `json:"balanceBefore" api:"required"`
	// Transaction timestamp
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Currency code
	Currency string `json:"currency" api:"required"`
	// Transaction status
	//
	// Any of "pending", "successful", "failed", "cancelled".
	Status string `json:"status" api:"required"`
	// Transaction type
	//
	// Any of "credit", "debit".
	Type string `json:"type" api:"required"`
	// Transaction description
	Description string `json:"description"`
	// Client reference
	Reference string `json:"reference"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AccountID     respjson.Field
		Amount        respjson.Field
		BalanceAfter  respjson.Field
		BalanceBefore respjson.Field
		CreatedAt     respjson.Field
		Currency      respjson.Field
		Status        respjson.Field
		Type          respjson.Field
		Description   respjson.Field
		Reference     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionListResponseDataData) RawJSON() string { return r.JSON.raw }
func (r *TransactionListResponseDataData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata
type TransactionListResponseDataMeta struct {
	// Current page number
	CurrentPage float64 `json:"currentPage" api:"required"`
	// Total number of pages
	PageCount float64 `json:"pageCount" api:"required"`
	// Items per page
	PerPage float64 `json:"perPage" api:"required"`
	// Total number of records
	Total float64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentPage respjson.Field
		PageCount   respjson.Field
		PerPage     respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionListResponseDataMeta) RawJSON() string { return r.JSON.raw }
func (r *TransactionListResponseDataMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionListParams struct {
	// Filter by account ID
	AccountID param.Opt[string] `query:"accountId,omitzero" json:"-"`
	// Filter by end date (ISO 8601)
	EndDate param.Opt[string] `query:"endDate,omitzero" json:"-"`
	// Page number (default: 1)
	Page param.Opt[float64] `query:"page,omitzero" json:"-"`
	// Items per page (default: 50)
	PerPage param.Opt[float64] `query:"perPage,omitzero" json:"-"`
	// Filter by start date (ISO 8601)
	StartDate param.Opt[string] `query:"startDate,omitzero" json:"-"`
	// Filter by transaction type
	//
	// Any of "credit", "debit".
	Type TransactionListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TransactionListParams]'s query parameters as `url.Values`.
func (r TransactionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by transaction type
type TransactionListParamsType string

const (
	TransactionListParamsTypeCredit TransactionListParamsType = "credit"
	TransactionListParamsTypeDebit  TransactionListParamsType = "debit"
)
