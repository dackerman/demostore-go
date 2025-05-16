// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dackermanstore

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/dackerman/demostore-go/internal/apijson"
	"github.com/dackerman/demostore-go/internal/requestconfig"
	"github.com/dackerman/demostore-go/option"
	"github.com/dackerman/demostore-go/packages/param"
	"github.com/dackerman/demostore-go/packages/respjson"
)

// ProductVariantService contains methods and other services that help with
// interacting with the stainless store API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductVariantService] method instead.
type ProductVariantService struct {
	Options []option.RequestOption
}

// NewProductVariantService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProductVariantService(opts ...option.RequestOption) (r ProductVariantService) {
	r = ProductVariantService{}
	r.Options = opts
	return
}

// Create Product Variant
func (r *ProductVariantService) New(ctx context.Context, productID string, params ProductVariantNewParams, opts ...option.RequestOption) (res *ProductVariant, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.OrgID, precfg.OrgID)
	if params.OrgID.Value == "" {
		err = errors.New("missing required org_id parameter")
		return
	}
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return
	}
	path := fmt.Sprintf("orgs/%s/products/%s/variants", params.OrgID.Value, productID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Read Product Variant
func (r *ProductVariantService) Get(ctx context.Context, productID string, variantID string, query ProductVariantGetParams, opts ...option.RequestOption) (res *ProductVariant, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.OrgID, precfg.OrgID)
	if query.OrgID.Value == "" {
		err = errors.New("missing required org_id parameter")
		return
	}
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return
	}
	if variantID == "" {
		err = errors.New("missing required variant_id parameter")
		return
	}
	path := fmt.Sprintf("orgs/%s/products/%s/variants/%s", query.OrgID.Value, productID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Product Variant
func (r *ProductVariantService) Update(ctx context.Context, productID string, variantID string, params ProductVariantUpdateParams, opts ...option.RequestOption) (res *ProductVariant, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.OrgID, precfg.OrgID)
	if params.OrgID.Value == "" {
		err = errors.New("missing required org_id parameter")
		return
	}
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return
	}
	if variantID == "" {
		err = errors.New("missing required variant_id parameter")
		return
	}
	path := fmt.Sprintf("orgs/%s/products/%s/variants/%s", params.OrgID.Value, productID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Read Product Variants
func (r *ProductVariantService) List(ctx context.Context, productID string, query ProductVariantListParams, opts ...option.RequestOption) (res *[]ProductVariant, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.OrgID, precfg.OrgID)
	if query.OrgID.Value == "" {
		err = errors.New("missing required org_id parameter")
		return
	}
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return
	}
	path := fmt.Sprintf("orgs/%s/products/%s/variants", query.OrgID.Value, productID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Product Variant
func (r *ProductVariantService) Delete(ctx context.Context, productID string, variantID string, body ProductVariantDeleteParams, opts ...option.RequestOption) (res *ProductVariantDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.OrgID, precfg.OrgID)
	if body.OrgID.Value == "" {
		err = errors.New("missing required org_id parameter")
		return
	}
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return
	}
	if variantID == "" {
		err = errors.New("missing required variant_id parameter")
		return
	}
	path := fmt.Sprintf("orgs/%s/products/%s/variants/%s", body.OrgID.Value, productID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Represents a ProductVariant record
type ProductVariant struct {
	ImageURL  string `json:"image_url,required"`
	Name      string `json:"name,required"`
	Price     int64  `json:"price,required"`
	ProductID string `json:"product_id,required"`
	VariantID string `json:"variant_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Name        respjson.Field
		Price       respjson.Field
		ProductID   respjson.Field
		VariantID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductVariant) RawJSON() string { return r.JSON.raw }
func (r *ProductVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductVariantDeleteResponse struct {
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductVariantDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ProductVariantDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductVariantNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID    param.Opt[string] `path:"org_id,omitzero,required" json:"-"`
	ImageURL string            `json:"image_url,required"`
	Name     string            `json:"name,required"`
	Price    int64             `json:"price,required"`
	paramObj
}

func (r ProductVariantNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProductVariantNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProductVariantNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductVariantGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Opt[string] `path:"org_id,omitzero,required" json:"-"`
	paramObj
}

type ProductVariantUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID    param.Opt[string] `path:"org_id,omitzero,required" json:"-"`
	ImageURL string            `json:"image_url,required"`
	Name     string            `json:"name,required"`
	Price    int64             `json:"price,required"`
	paramObj
}

func (r ProductVariantUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProductVariantUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProductVariantUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductVariantListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Opt[string] `path:"org_id,omitzero,required" json:"-"`
	paramObj
}

type ProductVariantDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Opt[string] `path:"org_id,omitzero,required" json:"-"`
	paramObj
}
