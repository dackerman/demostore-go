// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"
	"strconv"

	"github.com/dackerman/demostore-go/v2/internal/apijson"
	"github.com/dackerman/demostore-go/v2/internal/requestconfig"
	"github.com/dackerman/demostore-go/v2/option"
	"github.com/dackerman/demostore-go/v2/packages/param"
	"github.com/dackerman/demostore-go/v2/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type OffsetPagination[T any] struct {
	Data []T   `json:"data"`
	Next int64 `json:"next"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Next        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r OffsetPagination[T]) RawJSON() string { return r.JSON.raw }
func (r *OffsetPagination[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *OffsetPagination[T]) GetNextPage() (res *OffsetPagination[T], err error) {
	cfg := r.cfg.Clone(r.cfg.Context)

	next := r.Next
	length := int64(len(r.Data))

	if length > 0 && next != 0 {
		err = cfg.Apply(option.WithQuery("skip", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *OffsetPagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &OffsetPagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type OffsetPaginationAutoPager[T any] struct {
	page *OffsetPagination[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewOffsetPaginationAutoPager[T any](page *OffsetPagination[T], err error) *OffsetPaginationAutoPager[T] {
	return &OffsetPaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *OffsetPaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *OffsetPaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *OffsetPaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *OffsetPaginationAutoPager[T]) Index() int {
	return r.run
}
