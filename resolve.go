// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bila

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/bila-go/internal/apijson"
	"github.com/stainless-sdks/bila-go/internal/requestconfig"
	"github.com/stainless-sdks/bila-go/option"
	"github.com/stainless-sdks/bila-go/packages/param"
	"github.com/stainless-sdks/bila-go/packages/respjson"
)

// Account resolution/verification endpoints
//
// ResolveService contains methods and other services that help with interacting
// with the bila API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResolveService] method instead.
type ResolveService struct {
	options []option.RequestOption
}

// NewResolveService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewResolveService(opts ...option.RequestOption) (r ResolveService) {
	r = ResolveService{}
	r.options = opts
	return
}

// Verify and retrieve bank account holder details
func (r *ResolveService) BankAccount(ctx context.Context, body ResolveBankAccountParams, opts ...option.RequestOption) (res *ResolveBankAccountResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/bila/resolve/bank-account"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Verify and retrieve mobile money account holder details
func (r *ResolveService) MobileMoney(ctx context.Context, body ResolveMobileMoneyParams, opts ...option.RequestOption) (res *ResolveMobileMoneyResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/bila/resolve/mobile-money"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ResolveBankAccountResponse struct {
	Data ResolveBankAccountResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r ResolveBankAccountResponse) RawJSON() string { return r.JSON.raw }
func (r *ResolveBankAccountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResolveBankAccountResponseData struct {
	// Account holder name
	AccountName string `json:"accountName" api:"required"`
	// Country code
	Country string `json:"country" api:"required"`
	// Bank account number
	AccountNumber string `json:"accountNumber"`
	// Bank ID
	BankID string `json:"bankId"`
	// Bank name
	BankName string `json:"bankName"`
	// Mobile money operator
	Operator string `json:"operator"`
	// Phone number
	Phone string `json:"phone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountName   respjson.Field
		Country       respjson.Field
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
func (r ResolveBankAccountResponseData) RawJSON() string { return r.JSON.raw }
func (r *ResolveBankAccountResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResolveMobileMoneyResponse struct {
	Data ResolveMobileMoneyResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r ResolveMobileMoneyResponse) RawJSON() string { return r.JSON.raw }
func (r *ResolveMobileMoneyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResolveMobileMoneyResponseData struct {
	// Account holder name
	AccountName string `json:"accountName" api:"required"`
	// Country code
	Country string `json:"country" api:"required"`
	// Bank account number
	AccountNumber string `json:"accountNumber"`
	// Bank ID
	BankID string `json:"bankId"`
	// Bank name
	BankName string `json:"bankName"`
	// Mobile money operator
	Operator string `json:"operator"`
	// Phone number
	Phone string `json:"phone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountName   respjson.Field
		Country       respjson.Field
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
func (r ResolveMobileMoneyResponseData) RawJSON() string { return r.JSON.raw }
func (r *ResolveMobileMoneyResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResolveBankAccountParams struct {
	// Bank account number
	AccountNumber string `json:"accountNumber" api:"required"`
	// Bank ID
	BankID string `json:"bankId" api:"required"`
	// Country code
	//
	// Any of "zm", "ng".
	Country ResolveBankAccountParamsCountry `json:"country,omitzero"`
	paramObj
}

func (r ResolveBankAccountParams) MarshalJSON() (data []byte, err error) {
	type shadow ResolveBankAccountParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResolveBankAccountParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Country code
type ResolveBankAccountParamsCountry string

const (
	ResolveBankAccountParamsCountryZm ResolveBankAccountParamsCountry = "zm"
	ResolveBankAccountParamsCountryNg ResolveBankAccountParamsCountry = "ng"
)

type ResolveMobileMoneyParams struct {
	// Country code
	//
	// Any of "zm", "ng".
	Country ResolveMobileMoneyParamsCountry `json:"country,omitzero" api:"required"`
	// Mobile money operator
	//
	// Any of "airtel", "mtn", "zamtel", "vodacom".
	Operator ResolveMobileMoneyParamsOperator `json:"operator,omitzero" api:"required"`
	// Mobile phone number
	Phone string `json:"phone" api:"required"`
	paramObj
}

func (r ResolveMobileMoneyParams) MarshalJSON() (data []byte, err error) {
	type shadow ResolveMobileMoneyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResolveMobileMoneyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Country code
type ResolveMobileMoneyParamsCountry string

const (
	ResolveMobileMoneyParamsCountryZm ResolveMobileMoneyParamsCountry = "zm"
	ResolveMobileMoneyParamsCountryNg ResolveMobileMoneyParamsCountry = "ng"
)

// Mobile money operator
type ResolveMobileMoneyParamsOperator string

const (
	ResolveMobileMoneyParamsOperatorAirtel  ResolveMobileMoneyParamsOperator = "airtel"
	ResolveMobileMoneyParamsOperatorMtn     ResolveMobileMoneyParamsOperator = "mtn"
	ResolveMobileMoneyParamsOperatorZamtel  ResolveMobileMoneyParamsOperator = "zamtel"
	ResolveMobileMoneyParamsOperatorVodacom ResolveMobileMoneyParamsOperator = "vodacom"
)
