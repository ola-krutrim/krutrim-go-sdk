// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package krutrim

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/ola-krutrim/krutrim-go-sdk/internal/apiquery"
	"github.com/ola-krutrim/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
	"github.com/ola-krutrim/krutrim-go-sdk/packages/param"
)

// FloatingIPListService contains methods and other services that help with
// interacting with the krutrim-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFloatingIPListService] method instead.
type FloatingIPListService struct {
	Options []option.RequestOption
}

// NewFloatingIPListService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFloatingIPListService(opts ...option.RequestOption) (r FloatingIPListService) {
	r = FloatingIPListService{}
	r.Options = opts
	return
}

// List Floating IPs
func (r *FloatingIPListService) List(ctx context.Context, params FloatingIPListListParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XRegion) {
		opts = append(opts, option.WithHeader("x-region", fmt.Sprintf("%s", params.XRegion)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "/v1/highlvlvpc/floating_ip_list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return
}

type FloatingIPListListParams struct {
	VpcID   string `query:"vpc_id,required" json:"-"`
	XRegion string `header:"x-region,required" json:"-"`
	paramObj
}

// URLQuery serializes [FloatingIPListListParams]'s query parameters as
// `url.Values`.
func (r FloatingIPListListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
