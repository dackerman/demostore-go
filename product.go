// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dackermanstore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dackerman/demostore-private-go/v2/internal/apijson"
	"github.com/dackerman/demostore-private-go/v2/internal/apiquery"
	"github.com/dackerman/demostore-private-go/v2/internal/requestconfig"
	"github.com/dackerman/demostore-private-go/v2/option"
	"github.com/dackerman/demostore-private-go/v2/packages/pagination"
	"github.com/dackerman/demostore-private-go/v2/packages/param"
	"github.com/dackerman/demostore-private-go/v2/packages/respjson"
)

// ProductService contains methods and other services that help with interacting
// with the stainless store API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductService] method instead.
type ProductService struct {
	Options  []option.RequestOption
	Variants ProductVariantService
}

// NewProductService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProductService(opts ...option.RequestOption) (r ProductService) {
	r = ProductService{}
	r.Options = opts
	r.Variants = NewProductVariantService(opts...)
	return
}

// Create Product
func (r *ProductService) New(ctx context.Context, params ProductNewParams, opts ...option.RequestOption) (res *Product, err error) {
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
	path := fmt.Sprintf("orgs/%s/products", params.OrgID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Read Product
func (r *ProductService) Get(ctx context.Context, productID string, query ProductGetParams, opts ...option.RequestOption) (res *Product, err error) {
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
	path := fmt.Sprintf("orgs/%s/products/%s", query.OrgID.Value, productID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Product
func (r *ProductService) Update(ctx context.Context, productID string, params ProductUpdateParams, opts ...option.RequestOption) (res *Product, err error) {
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
	path := fmt.Sprintf("orgs/%s/products/%s", params.OrgID.Value, productID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Read Products
func (r *ProductService) List(ctx context.Context, params ProductListParams, opts ...option.RequestOption) (res *pagination.OffsetPagination[Product], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.OrgID, precfg.OrgID)
	if params.OrgID.Value == "" {
		err = errors.New("missing required org_id parameter")
		return
	}
	path := fmt.Sprintf("orgs/%s/products", params.OrgID.Value)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Read Products
func (r *ProductService) ListAutoPaging(ctx context.Context, params ProductListParams, opts ...option.RequestOption) *pagination.OffsetPaginationAutoPager[Product] {
	return pagination.NewOffsetPaginationAutoPager(r.List(ctx, params, opts...))
}

// Delete Product
func (r *ProductService) Delete(ctx context.Context, productID string, body ProductDeleteParams, opts ...option.RequestOption) (res *ProductDeleteResponse, err error) {
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
	path := fmt.Sprintf("orgs/%s/products/%s", body.OrgID.Value, productID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Represents a Product record
type Product struct {
	Description string `json:"description,required"`
	ImageURL    string `json:"image_url,required"`
	Name        string `json:"name,required"`
	Price       int64  `json:"price,required"`
	ProductID   string `json:"product_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		ImageURL    respjson.Field
		Name        respjson.Field
		Price       respjson.Field
		ProductID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Product) RawJSON() string { return r.JSON.raw }
func (r *Product) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductDeleteResponse struct {
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ProductDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID       param.Opt[string] `path:"org_id,omitzero,required" json:"-"`
	Description string            `json:"description,required"`
	ImageURL    string            `json:"image_url,required"`
	Name        string            `json:"name,required"`
	Price       int64             `json:"price,required"`
	paramObj
}

func (r ProductNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProductNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProductNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductGetParams struct {
	// The ID of the organization
	//
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Opt[string] `path:"org_id,omitzero,required" json:"-"`
	paramObj
}

type ProductUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID       param.Opt[string] `path:"org_id,omitzero,required" json:"-"`
	Description string            `json:"description,required"`
	ImageURL    string            `json:"image_url,required"`
	Name        string            `json:"name,required"`
	Price       int64             `json:"price,required"`
	paramObj
}

func (r ProductUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProductUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProductUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Opt[string] `path:"org_id,omitzero,required" json:"-"`
	Limit param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Skip  param.Opt[int64]  `query:"skip,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProductListParams]'s query parameters as `url.Values`.
func (r ProductListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProductDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Opt[string] `path:"org_id,omitzero,required" json:"-"`
	paramObj
}
