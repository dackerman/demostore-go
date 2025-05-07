// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dackermanstore

import (
	"context"
	"net/http"

	"github.com/dackerman/demostore-go/internal/apijson"
	"github.com/dackerman/demostore-go/internal/param"
	"github.com/dackerman/demostore-go/internal/requestconfig"
	"github.com/dackerman/demostore-go/option"
)

// FuntoolService contains methods and other services that help with interacting
// with the stainless store API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFuntoolService] method instead.
type FuntoolService struct {
	Options []option.RequestOption
}

// NewFuntoolService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFuntoolService(opts ...option.RequestOption) (r *FuntoolService) {
	r = &FuntoolService{}
	r.Options = opts
	return
}

// Set Darkmode Value
func (r *FuntoolService) SetDarkmode(ctx context.Context, body FuntoolSetDarkmodeParams, opts ...option.RequestOption) (res *FuntoolSetDarkmodeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "funtools/set_darkmode_value"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type FuntoolSetDarkmodeResponse struct {
	Success bool                           `json:"success,required"`
	JSON    funtoolSetDarkmodeResponseJSON `json:"-"`
}

// funtoolSetDarkmodeResponseJSON contains the JSON metadata for the struct
// [FuntoolSetDarkmodeResponse]
type funtoolSetDarkmodeResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FuntoolSetDarkmodeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r funtoolSetDarkmodeResponseJSON) RawJSON() string {
	return r.raw
}

type FuntoolSetDarkmodeParams struct {
	Darkmode param.Field[bool] `json:"darkmode,required"`
}

func (r FuntoolSetDarkmodeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
