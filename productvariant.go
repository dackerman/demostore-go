// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dackermanstore

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/dackerman/demostore-go/internal/apijson"
	"github.com/dackerman/demostore-go/internal/param"
	"github.com/dackerman/demostore-go/internal/requestconfig"
	"github.com/dackerman/demostore-go/option"
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
func NewProductVariantService(opts ...option.RequestOption) (r *ProductVariantService) {
	r = &ProductVariantService{}
	r.Options = opts
	return
}

// Create Product Variant
func (r *ProductVariantService) New(ctx context.Context, productID string, body ProductVariantNewParams, opts ...option.RequestOption) (res *ProductVariant, err error) {
	opts = append(r.Options[:], opts...)
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return
	}
	path := fmt.Sprintf("products/%s/variants", productID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Read Product Variant
func (r *ProductVariantService) Get(ctx context.Context, productID string, variantID string, opts ...option.RequestOption) (res *ProductVariant, err error) {
	opts = append(r.Options[:], opts...)
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return
	}
	if variantID == "" {
		err = errors.New("missing required variant_id parameter")
		return
	}
	path := fmt.Sprintf("products/%s/variants/%s", productID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Product Variant
func (r *ProductVariantService) Update(ctx context.Context, productID string, variantID string, body ProductVariantUpdateParams, opts ...option.RequestOption) (res *ProductVariant, err error) {
	opts = append(r.Options[:], opts...)
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return
	}
	if variantID == "" {
		err = errors.New("missing required variant_id parameter")
		return
	}
	path := fmt.Sprintf("products/%s/variants/%s", productID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Read Product Variants
func (r *ProductVariantService) List(ctx context.Context, productID string, opts ...option.RequestOption) (res *[]ProductVariant, err error) {
	opts = append(r.Options[:], opts...)
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return
	}
	path := fmt.Sprintf("products/%s/variants", productID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Product Variant
func (r *ProductVariantService) Delete(ctx context.Context, productID string, variantID string, opts ...option.RequestOption) (res *ProductVariantDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return
	}
	if variantID == "" {
		err = errors.New("missing required variant_id parameter")
		return
	}
	path := fmt.Sprintf("products/%s/variants/%s", productID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Represents a ProductVariant record
type ProductVariant struct {
	ImageURL  string             `json:"image_url,required"`
	Name      string             `json:"name,required"`
	Price     int64              `json:"price,required"`
	ProductID string             `json:"product_id,required"`
	VariantID string             `json:"variant_id,required"`
	JSON      productVariantJSON `json:"-"`
}

// productVariantJSON contains the JSON metadata for the struct [ProductVariant]
type productVariantJSON struct {
	ImageURL    apijson.Field
	Name        apijson.Field
	Price       apijson.Field
	ProductID   apijson.Field
	VariantID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productVariantJSON) RawJSON() string {
	return r.raw
}

type ProductVariantDeleteResponse struct {
	Success bool                             `json:"success,required"`
	JSON    productVariantDeleteResponseJSON `json:"-"`
}

// productVariantDeleteResponseJSON contains the JSON metadata for the struct
// [ProductVariantDeleteResponse]
type productVariantDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductVariantDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productVariantDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ProductVariantNewParams struct {
	// Url of the image to display for the variant
	ImageURL param.Field[string] `json:"image_url,required"`
	Name     param.Field[string] `json:"name,required"`
	Price    param.Field[int64]  `json:"price,required"`
}

func (r ProductVariantNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductVariantUpdateParams struct {
	// Url of the image to display for the variant
	ImageURL param.Field[string] `json:"image_url,required"`
	Name     param.Field[string] `json:"name,required"`
	Price    param.Field[int64]  `json:"price,required"`
}

func (r ProductVariantUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
