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
func (r *ProductService) New(ctx context.Context, body ProductNewParams, opts ...option.RequestOption) (res *Product, err error) {
	opts = append(r.Options[:], opts...)
	path := "products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Read Product
func (r *ProductService) Get(ctx context.Context, productID string, opts ...option.RequestOption) (res *Product, err error) {
	opts = append(r.Options[:], opts...)
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return
	}
	path := fmt.Sprintf("products/%s", productID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Product
func (r *ProductService) Update(ctx context.Context, productID string, body ProductUpdateParams, opts ...option.RequestOption) (res *Product, err error) {
	opts = append(r.Options[:], opts...)
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return
	}
	path := fmt.Sprintf("products/%s", productID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Read Products
func (r *ProductService) List(ctx context.Context, query ProductListParams, opts ...option.RequestOption) (res *ProductListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete Product
func (r *ProductService) Delete(ctx context.Context, productID string, opts ...option.RequestOption) (res *ProductDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return
	}
	path := fmt.Sprintf("products/%s", productID)
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

type ProductListResponse struct {
	Data []Product               `json:"data,required"`
	Next int64                   `json:"next,required,nullable"`
	JSON productListResponseJSON `json:"-"`
}

// productListResponseJSON contains the JSON metadata for the struct
// [ProductListResponse]
type productListResponseJSON struct {
	Data        apijson.Field
	Next        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListResponseJSON) RawJSON() string {
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
	Description param.Field[string] `json:"description,required"`
	ImageURL    param.Field[string] `json:"image_url,required"`
	Name        param.Field[string] `json:"name,required"`
	Price       param.Field[int64]  `json:"price,required"`
}

func (r ProductNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductUpdateParams struct {
	Description param.Field[string] `json:"description,required"`
	ImageURL    param.Field[string] `json:"image_url,required"`
	Name        param.Field[string] `json:"name,required"`
	Price       param.Field[int64]  `json:"price,required"`
}

func (r ProductUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductListParams struct {
	Limit param.Field[int64] `query:"limit"`
	Skip  param.Field[int64] `query:"skip"`
}

// URLQuery serializes [ProductListParams]'s query parameters as `url.Values`.
func (r ProductListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
