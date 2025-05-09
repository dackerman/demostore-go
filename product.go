// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dackermanstore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dackerman/demostore-go/internal/apijson"
	"github.com/dackerman/demostore-go/internal/apiquery"
	"github.com/dackerman/demostore-go/internal/param"
	"github.com/dackerman/demostore-go/internal/requestconfig"
	"github.com/dackerman/demostore-go/option"
	"github.com/dackerman/demostore-go/packages/pagination"
)

// ProductService contains methods and other services that help with interacting
// with the stainless store API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductService] method instead.
type ProductService struct {
	Options  []option.RequestOption
	Variants *ProductVariantService
}

// NewProductService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProductService(opts ...option.RequestOption) (r *ProductService) {
	r = &ProductService{}
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
	path := fmt.Sprintf("orgs/%s/products", params.OrgID)
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
	path := fmt.Sprintf("orgs/%s/products/%s", query.OrgID, productID)
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
	path := fmt.Sprintf("orgs/%s/products/%s", params.OrgID, productID)
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
	path := fmt.Sprintf("orgs/%s/products", params.OrgID)
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
	path := fmt.Sprintf("orgs/%s/products/%s", body.OrgID, productID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Represents a Product record
type Product struct {
	Description string      `json:"description,required"`
	ImageURL    string      `json:"image_url,required"`
	Name        string      `json:"name,required"`
	Price       int64       `json:"price,required"`
	ProductID   string      `json:"product_id,required"`
	JSON        productJSON `json:"-"`
}

// productJSON contains the JSON metadata for the struct [Product]
type productJSON struct {
	Description apijson.Field
	ImageURL    apijson.Field
	Name        apijson.Field
	Price       apijson.Field
	ProductID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Product) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productJSON) RawJSON() string {
	return r.raw
}

type ProductDeleteResponse struct {
	Success bool                      `json:"success,required"`
	JSON    productDeleteResponseJSON `json:"-"`
}

// productDeleteResponseJSON contains the JSON metadata for the struct
// [ProductDeleteResponse]
type productDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ProductNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID       param.Field[string] `path:"org_id,required"`
	Description param.Field[string] `json:"description,required"`
	ImageURL    param.Field[string] `json:"image_url,required"`
	Name        param.Field[string] `json:"name,required"`
	Price       param.Field[int64]  `json:"price,required"`
}

func (r ProductNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"org_id,required"`
}

type ProductUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID       param.Field[string] `path:"org_id,required"`
	Description param.Field[string] `json:"description,required"`
	ImageURL    param.Field[string] `json:"image_url,required"`
	Name        param.Field[string] `json:"name,required"`
	Price       param.Field[int64]  `json:"price,required"`
}

func (r ProductUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"org_id,required"`
	Limit param.Field[int64]  `query:"limit"`
	Skip  param.Field[int64]  `query:"skip"`
}

// URLQuery serializes [ProductListParams]'s query parameters as `url.Values`.
func (r ProductListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProductDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"org_id,required"`
}
