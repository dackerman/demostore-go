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
func (r *ProductVariantService) New(ctx context.Context, params ProductVariantNewParams, opts ...option.RequestOption) (res *ProductVariant, err error) {
	opts = append(r.Options[:], opts...)
	if params.PathProductID.Value == "" {
		err = errors.New("missing required path_product_id parameter")
		return
	}
	path := fmt.Sprintf("products/%s/variants", params.PathProductID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Read Product Variant
func (r *ProductVariantService) Get(ctx context.Context, productID int64, variantID int64, opts ...option.RequestOption) (res *ProductVariant, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("products/%v/variants/%v", productID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Product Variant
func (r *ProductVariantService) Update(ctx context.Context, variantID int64, params ProductVariantUpdateParams, opts ...option.RequestOption) (res *ProductVariant, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("products/%v/variants/%v", params.PathProductID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Read Product Variants
func (r *ProductVariantService) List(ctx context.Context, productID int64, opts ...option.RequestOption) (res *[]ProductVariant, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("products/%v/variants", productID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Product Variant
func (r *ProductVariantService) Delete(ctx context.Context, productID int64, variantID int64, opts ...option.RequestOption) (res *ProductVariantDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("products/%v/variants/%v", productID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Represents a ProductVariant record
type ProductVariant struct {
	ID        int64              `json:"id,required"`
	AddlPrice float64            `json:"addl_price,required"`
	ImageURL  string             `json:"image_url,required"`
	Name      string             `json:"name,required"`
	ProductID string             `json:"product_id,required"`
	JSON      productVariantJSON `json:"-"`
}

// productVariantJSON contains the JSON metadata for the struct [ProductVariant]
type productVariantJSON struct {
	ID          apijson.Field
	AddlPrice   apijson.Field
	ImageURL    apijson.Field
	Name        apijson.Field
	ProductID   apijson.Field
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
	PathProductID param.Field[string]  `path:"product_id,required"`
	AddlPrice     param.Field[float64] `json:"addl_price,required"`
	ImageURL      param.Field[string]  `json:"image_url,required"`
	Name          param.Field[string]  `json:"name,required"`
	BodyProductID param.Field[string]  `json:"product_id,required"`
	ID            param.Field[int64]   `json:"id"`
}

func (r ProductVariantNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductVariantUpdateParams struct {
	PathProductID param.Field[int64]                                    `path:"product_id,required"`
	ID            param.Field[ProductVariantUpdateParamsIDUnion]        `json:"id"`
	AddlPrice     param.Field[ProductVariantUpdateParamsAddlPriceUnion] `json:"addl_price"`
	ImageURL      param.Field[string]                                   `json:"image_url"`
	Name          param.Field[string]                                   `json:"name"`
	BodyProductID param.Field[string]                                   `json:"product_id"`
}

func (r ProductVariantUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductVariantUpdateParamsID struct {
	Decrement param.Field[int64] `json:"decrement"`
	Divide    param.Field[int64] `json:"divide"`
	Increment param.Field[int64] `json:"increment"`
	Multiply  param.Field[int64] `json:"multiply"`
	Set       param.Field[int64] `json:"set"`
}

func (r ProductVariantUpdateParamsID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductVariantUpdateParamsID) ImplementsProductVariantUpdateParamsIDUnion() {}

// Satisfied by [ProductVariantUpdateParamsIDIntSetInput],
// [ProductVariantUpdateParamsIDIntDivideInput],
// [ProductVariantUpdateParamsIDIntMultiplyInput],
// [ProductVariantUpdateParamsIDIntIncrementInput],
// [ProductVariantUpdateParamsIDIntDecrementInput], [shared.UnionInt],
// [ProductVariantUpdateParamsID].
type ProductVariantUpdateParamsIDUnion interface {
	ImplementsProductVariantUpdateParamsIDUnion()
}

type ProductVariantUpdateParamsIDIntSetInput struct {
	Set param.Field[int64] `json:"set,required"`
}

func (r ProductVariantUpdateParamsIDIntSetInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductVariantUpdateParamsIDIntSetInput) ImplementsProductVariantUpdateParamsIDUnion() {}

type ProductVariantUpdateParamsIDIntDivideInput struct {
	Divide param.Field[int64] `json:"divide,required"`
}

func (r ProductVariantUpdateParamsIDIntDivideInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductVariantUpdateParamsIDIntDivideInput) ImplementsProductVariantUpdateParamsIDUnion() {}

type ProductVariantUpdateParamsIDIntMultiplyInput struct {
	Multiply param.Field[int64] `json:"multiply,required"`
}

func (r ProductVariantUpdateParamsIDIntMultiplyInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductVariantUpdateParamsIDIntMultiplyInput) ImplementsProductVariantUpdateParamsIDUnion() {}

type ProductVariantUpdateParamsIDIntIncrementInput struct {
	Increment param.Field[int64] `json:"increment,required"`
}

func (r ProductVariantUpdateParamsIDIntIncrementInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductVariantUpdateParamsIDIntIncrementInput) ImplementsProductVariantUpdateParamsIDUnion() {
}

type ProductVariantUpdateParamsIDIntDecrementInput struct {
	Decrement param.Field[int64] `json:"decrement,required"`
}

func (r ProductVariantUpdateParamsIDIntDecrementInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductVariantUpdateParamsIDIntDecrementInput) ImplementsProductVariantUpdateParamsIDUnion() {
}

type ProductVariantUpdateParamsAddlPrice struct {
	Decrement param.Field[float64] `json:"decrement"`
	Divide    param.Field[float64] `json:"divide"`
	Increment param.Field[float64] `json:"increment"`
	Multiply  param.Field[float64] `json:"multiply"`
	Set       param.Field[float64] `json:"set"`
}

func (r ProductVariantUpdateParamsAddlPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductVariantUpdateParamsAddlPrice) ImplementsProductVariantUpdateParamsAddlPriceUnion() {}

// Satisfied by [ProductVariantUpdateParamsAddlPriceFloatSetInput],
// [ProductVariantUpdateParamsAddlPriceFloatDivideInput],
// [ProductVariantUpdateParamsAddlPriceFloatMultiplyInput],
// [ProductVariantUpdateParamsAddlPriceFloatIncrementInput],
// [ProductVariantUpdateParamsAddlPriceFloatDecrementInput], [shared.UnionFloat],
// [ProductVariantUpdateParamsAddlPrice].
type ProductVariantUpdateParamsAddlPriceUnion interface {
	ImplementsProductVariantUpdateParamsAddlPriceUnion()
}

type ProductVariantUpdateParamsAddlPriceFloatSetInput struct {
	Set param.Field[float64] `json:"set,required"`
}

func (r ProductVariantUpdateParamsAddlPriceFloatSetInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductVariantUpdateParamsAddlPriceFloatSetInput) ImplementsProductVariantUpdateParamsAddlPriceUnion() {
}

type ProductVariantUpdateParamsAddlPriceFloatDivideInput struct {
	Divide param.Field[float64] `json:"divide,required"`
}

func (r ProductVariantUpdateParamsAddlPriceFloatDivideInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductVariantUpdateParamsAddlPriceFloatDivideInput) ImplementsProductVariantUpdateParamsAddlPriceUnion() {
}

type ProductVariantUpdateParamsAddlPriceFloatMultiplyInput struct {
	Multiply param.Field[float64] `json:"multiply,required"`
}

func (r ProductVariantUpdateParamsAddlPriceFloatMultiplyInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductVariantUpdateParamsAddlPriceFloatMultiplyInput) ImplementsProductVariantUpdateParamsAddlPriceUnion() {
}

type ProductVariantUpdateParamsAddlPriceFloatIncrementInput struct {
	Increment param.Field[float64] `json:"increment,required"`
}

func (r ProductVariantUpdateParamsAddlPriceFloatIncrementInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductVariantUpdateParamsAddlPriceFloatIncrementInput) ImplementsProductVariantUpdateParamsAddlPriceUnion() {
}

type ProductVariantUpdateParamsAddlPriceFloatDecrementInput struct {
	Decrement param.Field[float64] `json:"decrement,required"`
}

func (r ProductVariantUpdateParamsAddlPriceFloatDecrementInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductVariantUpdateParamsAddlPriceFloatDecrementInput) ImplementsProductVariantUpdateParamsAddlPriceUnion() {
}
