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

// Transfer recipient management endpoints
//
// TransferRecipientService contains methods and other services that help with
// interacting with the bila API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransferRecipientService] method instead.
type TransferRecipientService struct {
	options []option.RequestOption
}

// NewTransferRecipientService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransferRecipientService(opts ...option.RequestOption) (r TransferRecipientService) {
	r = TransferRecipientService{}
	r.options = opts
	return
}

// Retrieve a single transfer recipient by its UUID
func (r *TransferRecipientService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TransferRecipientGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/bila/transfer-recipients/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve a paginated list of saved transfer recipients
func (r *TransferRecipientService) List(ctx context.Context, query TransferRecipientListParams, opts ...option.RequestOption) (res *TransferRecipientListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/bila/transfer-recipients"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Create a new bank account transfer recipient
func (r *TransferRecipientService) NewBankAccount(ctx context.Context, body TransferRecipientNewBankAccountParams, opts ...option.RequestOption) (res *TransferRecipientNewBankAccountResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/bila/transfer-recipients/bank-account"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a new mobile money transfer recipient
func (r *TransferRecipientService) NewMobileMoney(ctx context.Context, body TransferRecipientNewMobileMoneyParams, opts ...option.RequestOption) (res *TransferRecipientNewMobileMoneyResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/bila/transfer-recipients/mobile-money"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type RecipientResponseDto struct {
	// Recipient UUID
	ID string `json:"id" api:"required" format:"uuid"`
	// Account holder name
	AccountName string `json:"accountName" api:"required"`
	// Country code
	Country string `json:"country" api:"required"`
	// Creation timestamp
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Transfer recipient type
	//
	// Any of "bank-account", "mobile-money".
	Type RecipientResponseDtoType `json:"type" api:"required"`
	// Bank account number (bank-account only)
	AccountNumber string `json:"accountNumber"`
	// Bank ID (bank-account only)
	BankID string `json:"bankId"`
	// Bank name (bank-account only)
	BankName string `json:"bankName"`
	// Mobile money operator (mobile-money only)
	Operator string `json:"operator"`
	// Phone number (mobile-money only)
	Phone string `json:"phone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AccountName   respjson.Field
		Country       respjson.Field
		CreatedAt     respjson.Field
		Type          respjson.Field
		AccountNumber respjson.Field
		BankID        respjson.Field
		BankName      respjson.Field
		Operator      respjson.Field
		Phone         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecipientResponseDto) RawJSON() string { return r.JSON.raw }
func (r *RecipientResponseDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transfer recipient type
type RecipientResponseDtoType string

const (
	RecipientResponseDtoTypeBankAccount RecipientResponseDtoType = "bank-account"
	RecipientResponseDtoTypeMobileMoney RecipientResponseDtoType = "mobile-money"
)

type TransferRecipientGetResponse struct {
	// Response message
	Message string `json:"message" api:"required"`
	// Request success status
	Status bool                 `json:"status" api:"required"`
	Data   RecipientResponseDto `json:"data"`
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
func (r TransferRecipientGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TransferRecipientGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferRecipientListResponse struct {
	// Response message
	Message string `json:"message" api:"required"`
	// Request success status
	Status bool                              `json:"status" api:"required"`
	Data   TransferRecipientListResponseData `json:"data"`
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
func (r TransferRecipientListResponse) RawJSON() string { return r.JSON.raw }
func (r *TransferRecipientListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferRecipientListResponseData struct {
	// List of recipients
	Data []RecipientResponseDto `json:"data" api:"required"`
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
func (r TransferRecipientListResponseData) RawJSON() string { return r.JSON.raw }
func (r *TransferRecipientListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferRecipientNewBankAccountResponse struct {
	// Response message
	Message string `json:"message" api:"required"`
	// Request success status
	Status bool                 `json:"status" api:"required"`
	Data   RecipientResponseDto `json:"data"`
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
func (r TransferRecipientNewBankAccountResponse) RawJSON() string { return r.JSON.raw }
func (r *TransferRecipientNewBankAccountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferRecipientNewMobileMoneyResponse struct {
	// Response message
	Message string `json:"message" api:"required"`
	// Request success status
	Status bool                 `json:"status" api:"required"`
	Data   RecipientResponseDto `json:"data"`
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
func (r TransferRecipientNewMobileMoneyResponse) RawJSON() string { return r.JSON.raw }
func (r *TransferRecipientNewMobileMoneyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferRecipientListParams struct {
	// Page number (default: 1)
	Page param.Opt[float64] `query:"page,omitzero" json:"-"`
	// Items per page (default: 50)
	PerPage param.Opt[float64] `query:"perPage,omitzero" json:"-"`
	// Filter by recipient type
	//
	// Any of "bank-account", "mobile-money".
	Type TransferRecipientListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TransferRecipientListParams]'s query parameters as
// `url.Values`.
func (r TransferRecipientListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by recipient type
type TransferRecipientListParamsType string

const (
	TransferRecipientListParamsTypeBankAccount TransferRecipientListParamsType = "bank-account"
	TransferRecipientListParamsTypeMobileMoney TransferRecipientListParamsType = "mobile-money"
)

type TransferRecipientNewBankAccountParams struct {
	// Bank account number
	AccountNumber string `json:"accountNumber" api:"required"`
	// Bank ID
	BankID string `json:"bankId" api:"required"`
	// Account holder name (optional, will be resolved)
	AccountName param.Opt[string] `json:"accountName,omitzero"`
	// Country code
	//
	// Any of "zm".
	Country TransferRecipientNewBankAccountParamsCountry `json:"country,omitzero"`
	paramObj
}

func (r TransferRecipientNewBankAccountParams) MarshalJSON() (data []byte, err error) {
	type shadow TransferRecipientNewBankAccountParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransferRecipientNewBankAccountParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Country code
type TransferRecipientNewBankAccountParamsCountry string

const (
	TransferRecipientNewBankAccountParamsCountryZm TransferRecipientNewBankAccountParamsCountry = "zm"
)

type TransferRecipientNewMobileMoneyParams struct {
	// Country code
	//
	// Any of "zm".
	Country TransferRecipientNewMobileMoneyParamsCountry `json:"country,omitzero" api:"required"`
	// Mobile money operator
	//
	// Any of "airtel", "mtn", "zamtel".
	Operator TransferRecipientNewMobileMoneyParamsOperator `json:"operator,omitzero" api:"required"`
	// Mobile phone number
	Phone string `json:"phone" api:"required"`
	// Account holder name (optional, will be resolved)
	AccountName param.Opt[string] `json:"accountName,omitzero"`
	paramObj
}

func (r TransferRecipientNewMobileMoneyParams) MarshalJSON() (data []byte, err error) {
	type shadow TransferRecipientNewMobileMoneyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransferRecipientNewMobileMoneyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Country code
type TransferRecipientNewMobileMoneyParamsCountry string

const (
	TransferRecipientNewMobileMoneyParamsCountryZm TransferRecipientNewMobileMoneyParamsCountry = "zm"
)

// Mobile money operator
type TransferRecipientNewMobileMoneyParamsOperator string

const (
	TransferRecipientNewMobileMoneyParamsOperatorAirtel TransferRecipientNewMobileMoneyParamsOperator = "airtel"
	TransferRecipientNewMobileMoneyParamsOperatorMtn    TransferRecipientNewMobileMoneyParamsOperator = "mtn"
	TransferRecipientNewMobileMoneyParamsOperatorZamtel TransferRecipientNewMobileMoneyParamsOperator = "zamtel"
)
