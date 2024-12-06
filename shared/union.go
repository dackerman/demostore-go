// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionInt int64

func (UnionInt) ImplementsProductUpdateParamsIDUnion()        {}
func (UnionInt) ImplementsProductVariantUpdateParamsIDUnion() {}

type UnionFloat float64

func (UnionFloat) ImplementsProductUpdateParamsPriceUnion()            {}
func (UnionFloat) ImplementsProductVariantUpdateParamsAddlPriceUnion() {}
