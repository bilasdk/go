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

	"github.com/bilasdk/go/internal/apijson"
	"github.com/bilasdk/go/internal/apiquery"
	"github.com/bilasdk/go/internal/requestconfig"
	"github.com/bilasdk/go/option"
	"github.com/bilasdk/go/packages/param"
	"github.com/bilasdk/go/packages/respjson"
	"github.com/bilasdk/go/shared"
)

// Account/wallet management endpoints
//
// AccountService contains methods and other services that help with interacting
// with the bila API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r AccountService) {
	r = AccountService{}
	r.options = opts
	return
}

// Retrieve a single account by its UUID
func (r *AccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/bila/accounts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve a paginated list of accounts/wallets for the authenticated merchant
func (r *AccountService) List(ctx context.Context, query AccountListParams, opts ...option.RequestOption) (res *AccountListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/bila/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the balance of a specific account
func (r *AccountService) GetBalance(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountGetBalanceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/bila/accounts/%s/balance", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AccountDetailsDto struct {
	// Account holder name
	AccountName string `json:"accountName" api:"required"`
	// Account detail type
	Type string `json:"type" api:"required"`
	// Till number (for mobile money)
	TillNumber string `json:"tillNumber"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountName respjson.Field
		Type        respjson.Field
		TillNumber  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountDetailsDto) RawJSON() string { return r.JSON.raw }
func (r *AccountDetailsDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountResponseDto struct {
	// Account UUID
	ID string `json:"id" api:"required" format:"uuid"`
	// Account creation timestamp
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Currency code
	Currency string `json:"currency" api:"required"`
	// Account details
	Details AccountDetailsDto `json:"details" api:"required"`
	// Account status
	//
	// Any of "active", "inactive", "suspended".
	Status AccountResponseDtoStatus `json:"status" api:"required"`
	// Account type
	//
	// Any of "main", "sub", "virtual".
	Type AccountResponseDtoType `json:"type" api:"required"`
	// Available balance
	AvailableBalance string `json:"availableBalance"`
	// Ledger balance
	LedgerBalance string `json:"ledgerBalance"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		Currency         respjson.Field
		Details          respjson.Field
		Status           respjson.Field
		Type             respjson.Field
		AvailableBalance respjson.Field
		LedgerBalance    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountResponseDto) RawJSON() string { return r.JSON.raw }
func (r *AccountResponseDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account status
type AccountResponseDtoStatus string

const (
	AccountResponseDtoStatusActive    AccountResponseDtoStatus = "active"
	AccountResponseDtoStatusInactive  AccountResponseDtoStatus = "inactive"
	AccountResponseDtoStatusSuspended AccountResponseDtoStatus = "suspended"
)

// Account type
type AccountResponseDtoType string

const (
	AccountResponseDtoTypeMain    AccountResponseDtoType = "main"
	AccountResponseDtoTypeSub     AccountResponseDtoType = "sub"
	AccountResponseDtoTypeVirtual AccountResponseDtoType = "virtual"
)

type AccountGetResponse struct {
	// Response message
	Message string `json:"message" api:"required"`
	// Request success status
	Status bool               `json:"status" api:"required"`
	Data   AccountResponseDto `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Status      respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListResponse struct {
	// Response message
	Message string `json:"message" api:"required"`
	// Request success status
	Status bool                    `json:"status" api:"required"`
	Data   AccountListResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Status      respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountListResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListResponseData struct {
	// List of accounts
	Data []AccountResponseDto `json:"data" api:"required"`
	// Pagination metadata
	Meta shared.PaginationMetaDto `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountListResponseData) RawJSON() string { return r.JSON.raw }
func (r *AccountListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGetBalanceResponse struct {
	// Response message
	Message string `json:"message" api:"required"`
	// Request success status
	Status bool                          `json:"status" api:"required"`
	Data   AccountGetBalanceResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Status      respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountGetBalanceResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountGetBalanceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGetBalanceResponseData struct {
	// Available balance
	AvailableBalance string `json:"availableBalance" api:"required"`
	// Currency code
	Currency string `json:"currency" api:"required"`
	// Ledger balance
	LedgerBalance string `json:"ledgerBalance" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailableBalance respjson.Field
		Currency         respjson.Field
		LedgerBalance    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountGetBalanceResponseData) RawJSON() string { return r.JSON.raw }
func (r *AccountGetBalanceResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListParams struct {
	// Page number (default: 1)
	Page param.Opt[float64] `query:"page,omitzero" json:"-"`
	// Items per page (default: 50)
	PerPage param.Opt[float64] `query:"perPage,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
