// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/bilasdk/go/internal/apijson"
	"github.com/bilasdk/go/packages/param"
	"github.com/bilasdk/go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type PaginationMetaDto struct {
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
func (r PaginationMetaDto) RawJSON() string { return r.JSON.raw }
func (r *PaginationMetaDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
