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

// Webhook configuration and delivery history
//
// WebhookService contains methods and other services that help with interacting
// with the bila API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.options = opts
	return
}

// Create a webhook config
func (r *WebhookService) New(ctx context.Context, body WebhookNewParams, opts ...option.RequestOption) (res *WebhookNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/bila/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a webhook config
func (r *WebhookService) Update(ctx context.Context, id string, body WebhookUpdateParams, opts ...option.RequestOption) (res *WebhookUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/bila/webhooks/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List webhook configs
func (r *WebhookService) List(ctx context.Context, opts ...option.RequestOption) (res *WebhookListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/bila/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Deactivate a webhook
func (r *WebhookService) Deactivate(ctx context.Context, id string, opts ...option.RequestOption) (res *BilaResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/bila/webhooks/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Get delivery history
func (r *WebhookService) GetDeliveries(ctx context.Context, id string, query WebhookGetDeliveriesParams, opts ...option.RequestOption) (res *WebhookGetDeliveriesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/bila/webhooks/%s/deliveries", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List webhook event types
func (r *WebhookService) ListEvents(ctx context.Context, opts ...option.RequestOption) (res *WebhookListEventsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/bila/webhooks/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Rotate webhook signing secret
func (r *WebhookService) RotateSecret(ctx context.Context, id string, opts ...option.RequestOption) (res *WebhookRotateSecretResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/bila/webhooks/%s/rotate-secret", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type WebhookNewResponse struct {
	Data WebhookNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r WebhookNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewResponseData struct {
	// Webhook config UUID
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Subscribed event types
	Events []string `json:"events" api:"required"`
	// Whether the webhook is active
	IsActive bool `json:"isActive" api:"required"`
	// Merchant UUID
	MerchantID string `json:"merchantId" api:"required" format:"uuid"`
	// Signing secret; plaintext only on create/rotate-secret, otherwise masked
	Secret    string    `json:"secret" api:"required"`
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Webhook endpoint URL
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Events      respjson.Field
		IsActive    respjson.Field
		MerchantID  respjson.Field
		Secret      respjson.Field
		UpdatedAt   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateResponse struct {
	Data WebhookUpdateResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r WebhookUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateResponseData struct {
	// Webhook config UUID
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Subscribed event types
	Events []string `json:"events" api:"required"`
	// Whether the webhook is active
	IsActive bool `json:"isActive" api:"required"`
	// Merchant UUID
	MerchantID string `json:"merchantId" api:"required" format:"uuid"`
	// Signing secret; plaintext only on create/rotate-secret, otherwise masked
	Secret    string    `json:"secret" api:"required"`
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Webhook endpoint URL
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Events      respjson.Field
		IsActive    respjson.Field
		MerchantID  respjson.Field
		Secret      respjson.Field
		UpdatedAt   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListResponse struct {
	Data []WebhookListResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r WebhookListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListResponseData struct {
	// Webhook config UUID
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Subscribed event types
	Events []string `json:"events" api:"required"`
	// Whether the webhook is active
	IsActive bool `json:"isActive" api:"required"`
	// Merchant UUID
	MerchantID string `json:"merchantId" api:"required" format:"uuid"`
	// Signing secret; plaintext only on create/rotate-secret, otherwise masked
	Secret    string    `json:"secret" api:"required"`
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Webhook endpoint URL
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Events      respjson.Field
		IsActive    respjson.Field
		MerchantID  respjson.Field
		Secret      respjson.Field
		UpdatedAt   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookGetDeliveriesResponse struct {
	Data WebhookGetDeliveriesResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r WebhookGetDeliveriesResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookGetDeliveriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookGetDeliveriesResponseData struct {
	// List of webhook deliveries
	Data []WebhookGetDeliveriesResponseDataData `json:"data" api:"required"`
	// Pagination metadata
	Meta WebhookGetDeliveriesResponseDataMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookGetDeliveriesResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookGetDeliveriesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookGetDeliveriesResponseDataData struct {
	// Delivery UUID
	ID string `json:"id" api:"required" format:"uuid"`
	// Number of delivery attempts
	Attempts  float64   `json:"attempts" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// When the delivery succeeded
	DeliveredAt time.Time `json:"deliveredAt" api:"required" format:"date-time"`
	// Webhook event type
	EventType string `json:"eventType" api:"required"`
	// When the delivery permanently failed
	FailedAt time.Time `json:"failedAt" api:"required" format:"date-time"`
	// Maximum delivery attempts
	MaxAttempts float64 `json:"maxAttempts" api:"required"`
	// When the next retry is scheduled
	NextRetryAt time.Time `json:"nextRetryAt" api:"required" format:"date-time"`
	// Event payload JSON as stored for delivery
	Payload map[string]any `json:"payload" api:"required"`
	// Response body from the merchant endpoint (truncated)
	ResponseBody string `json:"responseBody" api:"required"`
	// HTTP status code from the merchant endpoint
	ResponseStatus float64 `json:"responseStatus" api:"required"`
	// Delivery status
	//
	// Any of "QUEUED", "DELIVERED", "FAILED", "RETRYING".
	Status string `json:"status" api:"required"`
	// Webhook config UUID
	WebhookConfigID string `json:"webhookConfigId" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Attempts        respjson.Field
		CreatedAt       respjson.Field
		DeliveredAt     respjson.Field
		EventType       respjson.Field
		FailedAt        respjson.Field
		MaxAttempts     respjson.Field
		NextRetryAt     respjson.Field
		Payload         respjson.Field
		ResponseBody    respjson.Field
		ResponseStatus  respjson.Field
		Status          respjson.Field
		WebhookConfigID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookGetDeliveriesResponseDataData) RawJSON() string { return r.JSON.raw }
func (r *WebhookGetDeliveriesResponseDataData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata
type WebhookGetDeliveriesResponseDataMeta struct {
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
func (r WebhookGetDeliveriesResponseDataMeta) RawJSON() string { return r.JSON.raw }
func (r *WebhookGetDeliveriesResponseDataMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListEventsResponse struct {
	Data []string `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r WebhookListEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookListEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookRotateSecretResponse struct {
	Data WebhookRotateSecretResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BilaResponse
}

// Returns the unmodified JSON received from the API
func (r WebhookRotateSecretResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookRotateSecretResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookRotateSecretResponseData struct {
	// New signing secret (64-character hex, shown once)
	Secret string `json:"secret" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookRotateSecretResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookRotateSecretResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewParams struct {
	// Event types to subscribe to
	//
	// Any of "order.created", "order.paid", "order.cancelled", "stock.low",
	// "payment.created", "payment.completed", "payment.failed", "collection.pending",
	// "collection.completed", "collection.failed", "withdrawal.created",
	// "withdrawal.completed", "withdrawal.failed", "transaction.updated",
	// "transfer.pending", "transfer.completed", "transfer.failed",
	// "settlement.completed".
	Events []string `json:"events,omitzero" api:"required"`
	// Webhook endpoint URL
	URL string `json:"url" api:"required"`
	paramObj
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateParams struct {
	// Whether the webhook is active
	IsActive param.Opt[bool] `json:"isActive,omitzero"`
	// Webhook endpoint URL
	URL param.Opt[string] `json:"url,omitzero"`
	// Event types to subscribe to
	//
	// Any of "order.created", "order.paid", "order.cancelled", "stock.low",
	// "payment.created", "payment.completed", "payment.failed", "collection.pending",
	// "collection.completed", "collection.failed", "withdrawal.created",
	// "withdrawal.completed", "withdrawal.failed", "transaction.updated",
	// "transfer.pending", "transfer.completed", "transfer.failed",
	// "settlement.completed".
	Events []string `json:"events,omitzero"`
	paramObj
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookGetDeliveriesParams struct {
	// ISO 8601 end of createdAt range (inclusive)
	EndDate param.Opt[string] `query:"endDate,omitzero" json:"-"`
	// Filter by event type
	EventType param.Opt[string] `query:"eventType,omitzero" json:"-"`
	// Page number
	Page param.Opt[float64] `query:"page,omitzero" json:"-"`
	// Items per page
	PerPage param.Opt[float64] `query:"perPage,omitzero" json:"-"`
	// ISO 8601 start of createdAt range (inclusive)
	StartDate param.Opt[string] `query:"startDate,omitzero" json:"-"`
	// Filter by status (QUEUED, DELIVERED, FAILED, RETRYING)
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookGetDeliveriesParams]'s query parameters as
// `url.Values`.
func (r WebhookGetDeliveriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
