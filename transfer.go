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
)

// Payout/transfer operation endpoints
//
// TransferService contains methods and other services that help with interacting
// with the bila API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransferService] method instead.
type TransferService struct {
	options []option.RequestOption
}

// NewTransferService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransferService(opts ...option.RequestOption) (r TransferService) {
	r = TransferService{}
	r.options = opts
	return
}

// Retrieve a single transfer by its UUID
func (r *TransferService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TransferGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/bila/transfers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve a paginated list of transfers/payouts for the authenticated merchant
func (r *TransferService) List(ctx context.Context, query TransferListParams, opts ...option.RequestOption) (res *TransferListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/bila/transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve transfer status by client reference
func (r *TransferService) GetStatusByReference(ctx context.Context, reference string, opts ...option.RequestOption) (res *TransferGetStatusByReferenceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if reference == "" {
		err = errors.New("missing required reference parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/bila/transfers/status/%s", url.PathEscape(reference))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Initiate a transfer to a bank account. Creates a transaction record in your
// dashboard.
func (r *TransferService) InitiateBankTransfer(ctx context.Context, body TransferInitiateBankTransferParams, opts ...option.RequestOption) (res *TransferInitiateBankTransferResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/bila/transfers/bank-account"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Initiate a transfer to a mobile money account. Creates a transaction record in
// your dashboard.
func (r *TransferService) InitiateMobileMoneyTransfer(ctx context.Context, body TransferInitiateMobileMoneyTransferParams, opts ...option.RequestOption) (res *TransferInitiateMobileMoneyTransferResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/bila/transfers/mobile-money"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type TransferGetResponse struct {
	Data TransferGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r TransferGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TransferGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferGetResponseData struct {
	// Transfer ID
	ID string `json:"id" api:"required"`
	// Transfer amount
	Amount float64 `json:"amount" api:"required"`
	// Creation timestamp (from Payment)
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Currency code
	Currency string `json:"currency" api:"required"`
	// Recipient details
	Recipient TransferGetResponseDataRecipient `json:"recipient" api:"required"`
	// Client reference
	Reference string `json:"reference" api:"required"`
	// Transfer status
	//
	// Any of "pending", "successful", "failed".
	Status string `json:"status" api:"required"`
	// Transfer type
	//
	// Any of "bank-account", "mobile-money".
	Type string `json:"type" api:"required"`
	// Completion timestamp (from Payment.processedAt)
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// Transfer narration
	Narration string `json:"narration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		CreatedAt   respjson.Field
		Currency    respjson.Field
		Recipient   respjson.Field
		Reference   respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		CompletedAt respjson.Field
		Narration   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *TransferGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recipient details
type TransferGetResponseDataRecipient struct {
	// Account holder / recipient name
	AccountName string `json:"accountName" api:"required"`
	// Bank account number (bank-account only)
	AccountNumber string `json:"accountNumber"`
	// Bank name (bank-account only)
	BankName string `json:"bankName"`
	// Mobile money operator (mobile-money only)
	Operator string `json:"operator"`
	// Phone number (mobile-money only)
	Phone string `json:"phone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountName   respjson.Field
		AccountNumber respjson.Field
		BankName      respjson.Field
		Operator      respjson.Field
		Phone         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferGetResponseDataRecipient) RawJSON() string { return r.JSON.raw }
func (r *TransferGetResponseDataRecipient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferListResponse struct {
	Data TransferListResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r TransferListResponse) RawJSON() string { return r.JSON.raw }
func (r *TransferListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferListResponseData struct {
	// List of transfers
	Data []TransferListResponseDataData `json:"data" api:"required"`
	// Pagination metadata
	Meta TransferListResponseDataMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferListResponseData) RawJSON() string { return r.JSON.raw }
func (r *TransferListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferListResponseDataData struct {
	// Transfer ID
	ID string `json:"id" api:"required"`
	// Transfer amount
	Amount float64 `json:"amount" api:"required"`
	// Creation timestamp (from Payment)
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Currency code
	Currency string `json:"currency" api:"required"`
	// Recipient details
	Recipient TransferListResponseDataDataRecipient `json:"recipient" api:"required"`
	// Client reference
	Reference string `json:"reference" api:"required"`
	// Transfer status
	//
	// Any of "pending", "successful", "failed".
	Status string `json:"status" api:"required"`
	// Transfer type
	//
	// Any of "bank-account", "mobile-money".
	Type string `json:"type" api:"required"`
	// Completion timestamp (from Payment.processedAt)
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// Transfer narration
	Narration string `json:"narration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		CreatedAt   respjson.Field
		Currency    respjson.Field
		Recipient   respjson.Field
		Reference   respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		CompletedAt respjson.Field
		Narration   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferListResponseDataData) RawJSON() string { return r.JSON.raw }
func (r *TransferListResponseDataData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recipient details
type TransferListResponseDataDataRecipient struct {
	// Account holder / recipient name
	AccountName string `json:"accountName" api:"required"`
	// Bank account number (bank-account only)
	AccountNumber string `json:"accountNumber"`
	// Bank name (bank-account only)
	BankName string `json:"bankName"`
	// Mobile money operator (mobile-money only)
	Operator string `json:"operator"`
	// Phone number (mobile-money only)
	Phone string `json:"phone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountName   respjson.Field
		AccountNumber respjson.Field
		BankName      respjson.Field
		Operator      respjson.Field
		Phone         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferListResponseDataDataRecipient) RawJSON() string { return r.JSON.raw }
func (r *TransferListResponseDataDataRecipient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata
type TransferListResponseDataMeta struct {
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
func (r TransferListResponseDataMeta) RawJSON() string { return r.JSON.raw }
func (r *TransferListResponseDataMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferGetStatusByReferenceResponse struct {
	Data TransferGetStatusByReferenceResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r TransferGetStatusByReferenceResponse) RawJSON() string { return r.JSON.raw }
func (r *TransferGetStatusByReferenceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferGetStatusByReferenceResponseData struct {
	// Transfer ID
	ID string `json:"id" api:"required"`
	// Transfer amount
	Amount float64 `json:"amount" api:"required"`
	// Creation timestamp (from Payment)
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Currency code
	Currency string `json:"currency" api:"required"`
	// Recipient details
	Recipient TransferGetStatusByReferenceResponseDataRecipient `json:"recipient" api:"required"`
	// Client reference
	Reference string `json:"reference" api:"required"`
	// Transfer status
	//
	// Any of "pending", "successful", "failed".
	Status string `json:"status" api:"required"`
	// Transfer type
	//
	// Any of "bank-account", "mobile-money".
	Type string `json:"type" api:"required"`
	// Completion timestamp (from Payment.processedAt)
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// Transfer narration
	Narration string `json:"narration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		CreatedAt   respjson.Field
		Currency    respjson.Field
		Recipient   respjson.Field
		Reference   respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		CompletedAt respjson.Field
		Narration   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferGetStatusByReferenceResponseData) RawJSON() string { return r.JSON.raw }
func (r *TransferGetStatusByReferenceResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recipient details
type TransferGetStatusByReferenceResponseDataRecipient struct {
	// Account holder / recipient name
	AccountName string `json:"accountName" api:"required"`
	// Bank account number (bank-account only)
	AccountNumber string `json:"accountNumber"`
	// Bank name (bank-account only)
	BankName string `json:"bankName"`
	// Mobile money operator (mobile-money only)
	Operator string `json:"operator"`
	// Phone number (mobile-money only)
	Phone string `json:"phone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountName   respjson.Field
		AccountNumber respjson.Field
		BankName      respjson.Field
		Operator      respjson.Field
		Phone         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferGetStatusByReferenceResponseDataRecipient) RawJSON() string { return r.JSON.raw }
func (r *TransferGetStatusByReferenceResponseDataRecipient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferInitiateBankTransferResponse struct {
	Data TransferInitiateBankTransferResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r TransferInitiateBankTransferResponse) RawJSON() string { return r.JSON.raw }
func (r *TransferInitiateBankTransferResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferInitiateBankTransferResponseData struct {
	// Transfer ID
	ID string `json:"id" api:"required"`
	// Transfer amount
	Amount float64 `json:"amount" api:"required"`
	// Creation timestamp (from Payment)
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Currency code
	Currency string `json:"currency" api:"required"`
	// Recipient details
	Recipient TransferInitiateBankTransferResponseDataRecipient `json:"recipient" api:"required"`
	// Client reference
	Reference string `json:"reference" api:"required"`
	// Transfer status
	//
	// Any of "pending", "successful", "failed".
	Status string `json:"status" api:"required"`
	// Transfer type
	//
	// Any of "bank-account", "mobile-money".
	Type string `json:"type" api:"required"`
	// Completion timestamp (from Payment.processedAt)
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// Transfer narration
	Narration string `json:"narration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		CreatedAt   respjson.Field
		Currency    respjson.Field
		Recipient   respjson.Field
		Reference   respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		CompletedAt respjson.Field
		Narration   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferInitiateBankTransferResponseData) RawJSON() string { return r.JSON.raw }
func (r *TransferInitiateBankTransferResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recipient details
type TransferInitiateBankTransferResponseDataRecipient struct {
	// Account holder / recipient name
	AccountName string `json:"accountName" api:"required"`
	// Bank account number (bank-account only)
	AccountNumber string `json:"accountNumber"`
	// Bank name (bank-account only)
	BankName string `json:"bankName"`
	// Mobile money operator (mobile-money only)
	Operator string `json:"operator"`
	// Phone number (mobile-money only)
	Phone string `json:"phone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountName   respjson.Field
		AccountNumber respjson.Field
		BankName      respjson.Field
		Operator      respjson.Field
		Phone         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferInitiateBankTransferResponseDataRecipient) RawJSON() string { return r.JSON.raw }
func (r *TransferInitiateBankTransferResponseDataRecipient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferInitiateMobileMoneyTransferResponse struct {
	Data TransferInitiateMobileMoneyTransferResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r TransferInitiateMobileMoneyTransferResponse) RawJSON() string { return r.JSON.raw }
func (r *TransferInitiateMobileMoneyTransferResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferInitiateMobileMoneyTransferResponseData struct {
	// Transfer ID
	ID string `json:"id" api:"required"`
	// Transfer amount
	Amount float64 `json:"amount" api:"required"`
	// Creation timestamp (from Payment)
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Currency code
	Currency string `json:"currency" api:"required"`
	// Recipient details
	Recipient TransferInitiateMobileMoneyTransferResponseDataRecipient `json:"recipient" api:"required"`
	// Client reference
	Reference string `json:"reference" api:"required"`
	// Transfer status
	//
	// Any of "pending", "successful", "failed".
	Status string `json:"status" api:"required"`
	// Transfer type
	//
	// Any of "bank-account", "mobile-money".
	Type string `json:"type" api:"required"`
	// Completion timestamp (from Payment.processedAt)
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// Transfer narration
	Narration string `json:"narration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		CreatedAt   respjson.Field
		Currency    respjson.Field
		Recipient   respjson.Field
		Reference   respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		CompletedAt respjson.Field
		Narration   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferInitiateMobileMoneyTransferResponseData) RawJSON() string { return r.JSON.raw }
func (r *TransferInitiateMobileMoneyTransferResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recipient details
type TransferInitiateMobileMoneyTransferResponseDataRecipient struct {
	// Account holder / recipient name
	AccountName string `json:"accountName" api:"required"`
	// Bank account number (bank-account only)
	AccountNumber string `json:"accountNumber"`
	// Bank name (bank-account only)
	BankName string `json:"bankName"`
	// Mobile money operator (mobile-money only)
	Operator string `json:"operator"`
	// Phone number (mobile-money only)
	Phone string `json:"phone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountName   respjson.Field
		AccountNumber respjson.Field
		BankName      respjson.Field
		Operator      respjson.Field
		Phone         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferInitiateMobileMoneyTransferResponseDataRecipient) RawJSON() string { return r.JSON.raw }
func (r *TransferInitiateMobileMoneyTransferResponseDataRecipient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferListParams struct {
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
	// Filter by transfer status
	//
	// Any of "pending", "successful", "failed".
	Status TransferListParamsStatus `query:"status,omitzero" json:"-"`
	// Filter by transfer type
	//
	// Any of "bank-account", "mobile-money".
	Type TransferListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TransferListParams]'s query parameters as `url.Values`.
func (r TransferListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by transfer status
type TransferListParamsStatus string

const (
	TransferListParamsStatusPending    TransferListParamsStatus = "pending"
	TransferListParamsStatusSuccessful TransferListParamsStatus = "successful"
	TransferListParamsStatusFailed     TransferListParamsStatus = "failed"
)

// Filter by transfer type
type TransferListParamsType string

const (
	TransferListParamsTypeBankAccount TransferListParamsType = "bank-account"
	TransferListParamsTypeMobileMoney TransferListParamsType = "mobile-money"
)

type TransferInitiateBankTransferParams struct {
	// Source account UUID
	AccountID string `json:"accountId" api:"required" format:"uuid"`
	// Transfer amount
	Amount float64 `json:"amount" api:"required"`
	// Unique client reference (alphanumeric, dots, underscores, hyphens)
	Reference string `json:"reference" api:"required"`
	// Bank account number (required if no transferRecipientId)
	AccountNumber param.Opt[string] `json:"accountNumber,omitzero"`
	// Bank ID (required if no transferRecipientId)
	BankID param.Opt[string] `json:"bankId,omitzero"`
	// Transfer narration
	Narration param.Opt[string] `json:"narration,omitzero"`
	// Recipient name for the transaction record
	RecipientName param.Opt[string] `json:"recipientName,omitzero"`
	// Transfer recipient UUID (use this OR accountNumber+bankId)
	TransferRecipientID param.Opt[string] `json:"transferRecipientId,omitzero" format:"uuid"`
	// Source wallet ID to debit (optional, uses main wallet if not specified)
	WalletID param.Opt[string] `json:"walletId,omitzero"`
	// Country code
	//
	// Any of "zm", "ng".
	Country TransferInitiateBankTransferParamsCountry `json:"country,omitzero"`
	paramObj
}

func (r TransferInitiateBankTransferParams) MarshalJSON() (data []byte, err error) {
	type shadow TransferInitiateBankTransferParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransferInitiateBankTransferParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Country code
type TransferInitiateBankTransferParamsCountry string

const (
	TransferInitiateBankTransferParamsCountryZm TransferInitiateBankTransferParamsCountry = "zm"
	TransferInitiateBankTransferParamsCountryNg TransferInitiateBankTransferParamsCountry = "ng"
)

type TransferInitiateMobileMoneyTransferParams struct {
	// Transfer amount
	Amount float64 `json:"amount" api:"required"`
	// Country code
	//
	// Any of "zm", "ng".
	Country TransferInitiateMobileMoneyTransferParamsCountry `json:"country,omitzero" api:"required"`
	// Mobile money operator
	//
	// Any of "airtel", "mtn", "zamtel", "vodacom".
	Operator TransferInitiateMobileMoneyTransferParamsOperator `json:"operator,omitzero" api:"required"`
	// Recipient phone number
	Phone string `json:"phone" api:"required"`
	// Unique client reference
	Reference string `json:"reference" api:"required"`
	// Transfer narration
	Narration param.Opt[string] `json:"narration,omitzero"`
	// Recipient name for the transaction record
	RecipientName param.Opt[string] `json:"recipientName,omitzero"`
	// Source wallet ID to debit (defaults to main wallet if omitted)
	WalletID param.Opt[string] `json:"walletId,omitzero" format:"uuid"`
	paramObj
}

func (r TransferInitiateMobileMoneyTransferParams) MarshalJSON() (data []byte, err error) {
	type shadow TransferInitiateMobileMoneyTransferParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransferInitiateMobileMoneyTransferParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Country code
type TransferInitiateMobileMoneyTransferParamsCountry string

const (
	TransferInitiateMobileMoneyTransferParamsCountryZm TransferInitiateMobileMoneyTransferParamsCountry = "zm"
	TransferInitiateMobileMoneyTransferParamsCountryNg TransferInitiateMobileMoneyTransferParamsCountry = "ng"
)

// Mobile money operator
type TransferInitiateMobileMoneyTransferParamsOperator string

const (
	TransferInitiateMobileMoneyTransferParamsOperatorAirtel  TransferInitiateMobileMoneyTransferParamsOperator = "airtel"
	TransferInitiateMobileMoneyTransferParamsOperatorMtn     TransferInitiateMobileMoneyTransferParamsOperator = "mtn"
	TransferInitiateMobileMoneyTransferParamsOperatorZamtel  TransferInitiateMobileMoneyTransferParamsOperator = "zamtel"
	TransferInitiateMobileMoneyTransferParamsOperatorVodacom TransferInitiateMobileMoneyTransferParamsOperator = "vodacom"
)
