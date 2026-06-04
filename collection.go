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

// Payment collection operation endpoints
//
// CollectionService contains methods and other services that help with interacting
// with the bila API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCollectionService] method instead.
type CollectionService struct {
	options []option.RequestOption
}

// NewCollectionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCollectionService(opts ...option.RequestOption) (r CollectionService) {
	r = CollectionService{}
	r.options = opts
	return
}

// Retrieve a single collection by its UUID
func (r *CollectionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *CollectionGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/bila/collections/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve a paginated list of payment collections for the authenticated merchant
func (r *CollectionService) List(ctx context.Context, query CollectionListParams, opts ...option.RequestOption) (res *CollectionListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/bila/collections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve collection status by client reference
func (r *CollectionService) GetStatusByReference(ctx context.Context, reference string, opts ...option.RequestOption) (res *CollectionGetStatusByReferenceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if reference == "" {
		err = errors.New("missing required reference parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/bila/collections/status/%s", url.PathEscape(reference))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Initiate a payment collection from a mobile money account. Creates a transaction
// record in your dashboard.
func (r *CollectionService) InitiateMobileMoneyCollection(ctx context.Context, body CollectionInitiateMobileMoneyCollectionParams, opts ...option.RequestOption) (res *CollectionInitiateMobileMoneyCollectionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/bila/collections/mobile-money"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type CollectionGetResponse struct {
	Data CollectionGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r CollectionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CollectionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionGetResponseData struct {
	// Collection ID
	ID string `json:"id" api:"required"`
	// Collection amount
	Amount float64 `json:"amount" api:"required"`
	// Collection creation timestamp
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Currency code
	Currency string `json:"currency" api:"required"`
	// Customer details
	Customer CollectionGetResponseDataCustomer `json:"customer" api:"required"`
	// Client reference
	Reference string `json:"reference" api:"required"`
	// Collection status
	//
	// Any of "pending", "successful", "failed", "otp-required", "pay-offline".
	Status string `json:"status" api:"required"`
	// Collection completion timestamp
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// Who bears the collection platform fee
	//
	// Any of "merchant", "customer".
	FeeBearer string `json:"feeBearer"`
	// Collection narration
	Narration string `json:"narration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		CreatedAt   respjson.Field
		Currency    respjson.Field
		Customer    respjson.Field
		Reference   respjson.Field
		Status      respjson.Field
		CompletedAt respjson.Field
		FeeBearer   respjson.Field
		Narration   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *CollectionGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Customer details
type CollectionGetResponseDataCustomer struct {
	// Customer name
	Name string `json:"name" api:"required"`
	// Mobile money operator
	Operator string `json:"operator" api:"required"`
	// Customer phone number
	Phone string `json:"phone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Operator    respjson.Field
		Phone       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionGetResponseDataCustomer) RawJSON() string { return r.JSON.raw }
func (r *CollectionGetResponseDataCustomer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionListResponse struct {
	Data CollectionListResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r CollectionListResponse) RawJSON() string { return r.JSON.raw }
func (r *CollectionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionListResponseData struct {
	// List of collections
	Data []CollectionListResponseDataData `json:"data" api:"required"`
	// Pagination metadata
	Meta CollectionListResponseDataMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionListResponseData) RawJSON() string { return r.JSON.raw }
func (r *CollectionListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionListResponseDataData struct {
	// Collection ID
	ID string `json:"id" api:"required"`
	// Collection amount
	Amount float64 `json:"amount" api:"required"`
	// Collection creation timestamp
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Currency code
	Currency string `json:"currency" api:"required"`
	// Customer details
	Customer CollectionListResponseDataDataCustomer `json:"customer" api:"required"`
	// Client reference
	Reference string `json:"reference" api:"required"`
	// Collection status
	//
	// Any of "pending", "successful", "failed", "otp-required", "pay-offline".
	Status string `json:"status" api:"required"`
	// Collection completion timestamp
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// Who bears the collection platform fee
	//
	// Any of "merchant", "customer".
	FeeBearer string `json:"feeBearer"`
	// Collection narration
	Narration string `json:"narration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		CreatedAt   respjson.Field
		Currency    respjson.Field
		Customer    respjson.Field
		Reference   respjson.Field
		Status      respjson.Field
		CompletedAt respjson.Field
		FeeBearer   respjson.Field
		Narration   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionListResponseDataData) RawJSON() string { return r.JSON.raw }
func (r *CollectionListResponseDataData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Customer details
type CollectionListResponseDataDataCustomer struct {
	// Customer name
	Name string `json:"name" api:"required"`
	// Mobile money operator
	Operator string `json:"operator" api:"required"`
	// Customer phone number
	Phone string `json:"phone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Operator    respjson.Field
		Phone       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionListResponseDataDataCustomer) RawJSON() string { return r.JSON.raw }
func (r *CollectionListResponseDataDataCustomer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata
type CollectionListResponseDataMeta struct {
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
func (r CollectionListResponseDataMeta) RawJSON() string { return r.JSON.raw }
func (r *CollectionListResponseDataMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionGetStatusByReferenceResponse struct {
	Data CollectionGetStatusByReferenceResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r CollectionGetStatusByReferenceResponse) RawJSON() string { return r.JSON.raw }
func (r *CollectionGetStatusByReferenceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionGetStatusByReferenceResponseData struct {
	// Collection ID
	ID string `json:"id" api:"required"`
	// Collection amount
	Amount float64 `json:"amount" api:"required"`
	// Collection creation timestamp
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Currency code
	Currency string `json:"currency" api:"required"`
	// Customer details
	Customer CollectionGetStatusByReferenceResponseDataCustomer `json:"customer" api:"required"`
	// Client reference
	Reference string `json:"reference" api:"required"`
	// Collection status
	//
	// Any of "pending", "successful", "failed", "otp-required", "pay-offline".
	Status string `json:"status" api:"required"`
	// Collection completion timestamp
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// Who bears the collection platform fee
	//
	// Any of "merchant", "customer".
	FeeBearer string `json:"feeBearer"`
	// Collection narration
	Narration string `json:"narration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		CreatedAt   respjson.Field
		Currency    respjson.Field
		Customer    respjson.Field
		Reference   respjson.Field
		Status      respjson.Field
		CompletedAt respjson.Field
		FeeBearer   respjson.Field
		Narration   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionGetStatusByReferenceResponseData) RawJSON() string { return r.JSON.raw }
func (r *CollectionGetStatusByReferenceResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Customer details
type CollectionGetStatusByReferenceResponseDataCustomer struct {
	// Customer name
	Name string `json:"name" api:"required"`
	// Mobile money operator
	Operator string `json:"operator" api:"required"`
	// Customer phone number
	Phone string `json:"phone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Operator    respjson.Field
		Phone       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionGetStatusByReferenceResponseDataCustomer) RawJSON() string { return r.JSON.raw }
func (r *CollectionGetStatusByReferenceResponseDataCustomer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionInitiateMobileMoneyCollectionResponse struct {
	Data CollectionInitiateMobileMoneyCollectionResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r CollectionInitiateMobileMoneyCollectionResponse) RawJSON() string { return r.JSON.raw }
func (r *CollectionInitiateMobileMoneyCollectionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionInitiateMobileMoneyCollectionResponseData struct {
	// Collection ID
	ID string `json:"id" api:"required"`
	// Collection amount
	Amount float64 `json:"amount" api:"required"`
	// Collection creation timestamp
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Currency code
	Currency string `json:"currency" api:"required"`
	// Customer details
	Customer CollectionInitiateMobileMoneyCollectionResponseDataCustomer `json:"customer" api:"required"`
	// Client reference
	Reference string `json:"reference" api:"required"`
	// Collection status
	//
	// Any of "pending", "successful", "failed", "otp-required", "pay-offline".
	Status string `json:"status" api:"required"`
	// Collection completion timestamp
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// Who bears the collection platform fee
	//
	// Any of "merchant", "customer".
	FeeBearer string `json:"feeBearer"`
	// Collection narration
	Narration string `json:"narration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		CreatedAt   respjson.Field
		Currency    respjson.Field
		Customer    respjson.Field
		Reference   respjson.Field
		Status      respjson.Field
		CompletedAt respjson.Field
		FeeBearer   respjson.Field
		Narration   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionInitiateMobileMoneyCollectionResponseData) RawJSON() string { return r.JSON.raw }
func (r *CollectionInitiateMobileMoneyCollectionResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Customer details
type CollectionInitiateMobileMoneyCollectionResponseDataCustomer struct {
	// Customer name
	Name string `json:"name" api:"required"`
	// Mobile money operator
	Operator string `json:"operator" api:"required"`
	// Customer phone number
	Phone string `json:"phone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Operator    respjson.Field
		Phone       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionInitiateMobileMoneyCollectionResponseDataCustomer) RawJSON() string {
	return r.JSON.raw
}
func (r *CollectionInitiateMobileMoneyCollectionResponseDataCustomer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionListParams struct {
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
	// Filter by collection status
	//
	// Any of "pending", "successful", "failed", "otp-required", "pay-offline".
	Status CollectionListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CollectionListParams]'s query parameters as `url.Values`.
func (r CollectionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by collection status
type CollectionListParamsStatus string

const (
	CollectionListParamsStatusPending     CollectionListParamsStatus = "pending"
	CollectionListParamsStatusSuccessful  CollectionListParamsStatus = "successful"
	CollectionListParamsStatusFailed      CollectionListParamsStatus = "failed"
	CollectionListParamsStatusOtpRequired CollectionListParamsStatus = "otp-required"
	CollectionListParamsStatusPayOffline  CollectionListParamsStatus = "pay-offline"
)

type CollectionInitiateMobileMoneyCollectionParams struct {
	// Collection amount
	Amount float64 `json:"amount" api:"required"`
	// Country code
	//
	// Any of "zm", "ng".
	Country CollectionInitiateMobileMoneyCollectionParamsCountry `json:"country,omitzero" api:"required"`
	// Mobile money operator
	//
	// Any of "airtel", "mtn", "zamtel", "vodacom".
	Operator CollectionInitiateMobileMoneyCollectionParamsOperator `json:"operator,omitzero" api:"required"`
	// Customer phone number
	Phone string `json:"phone" api:"required"`
	// Unique client reference
	Reference string `json:"reference" api:"required"`
	// Target wallet ID to credit
	WalletID string `json:"walletId" api:"required" format:"uuid"`
	// Customer name for the transaction record
	CustomerName param.Opt[string] `json:"customerName,omitzero"`
	// Collection narration
	Narration param.Opt[string] `json:"narration,omitzero"`
	// Who bears the transaction fee
	//
	// Any of "merchant", "customer".
	Bearer CollectionInitiateMobileMoneyCollectionParamsBearer `json:"bearer,omitzero"`
	paramObj
}

func (r CollectionInitiateMobileMoneyCollectionParams) MarshalJSON() (data []byte, err error) {
	type shadow CollectionInitiateMobileMoneyCollectionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionInitiateMobileMoneyCollectionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Country code
type CollectionInitiateMobileMoneyCollectionParamsCountry string

const (
	CollectionInitiateMobileMoneyCollectionParamsCountryZm CollectionInitiateMobileMoneyCollectionParamsCountry = "zm"
	CollectionInitiateMobileMoneyCollectionParamsCountryNg CollectionInitiateMobileMoneyCollectionParamsCountry = "ng"
)

// Mobile money operator
type CollectionInitiateMobileMoneyCollectionParamsOperator string

const (
	CollectionInitiateMobileMoneyCollectionParamsOperatorAirtel  CollectionInitiateMobileMoneyCollectionParamsOperator = "airtel"
	CollectionInitiateMobileMoneyCollectionParamsOperatorMtn     CollectionInitiateMobileMoneyCollectionParamsOperator = "mtn"
	CollectionInitiateMobileMoneyCollectionParamsOperatorZamtel  CollectionInitiateMobileMoneyCollectionParamsOperator = "zamtel"
	CollectionInitiateMobileMoneyCollectionParamsOperatorVodacom CollectionInitiateMobileMoneyCollectionParamsOperator = "vodacom"
)

// Who bears the transaction fee
type CollectionInitiateMobileMoneyCollectionParamsBearer string

const (
	CollectionInitiateMobileMoneyCollectionParamsBearerMerchant CollectionInitiateMobileMoneyCollectionParamsBearer = "merchant"
	CollectionInitiateMobileMoneyCollectionParamsBearerCustomer CollectionInitiateMobileMoneyCollectionParamsBearer = "customer"
)
