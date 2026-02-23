// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package krutrim

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/ola-silicon/krutrim-go-sdk/internal/apiquery"
	"github.com/ola-silicon/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-silicon/krutrim-go-sdk/option"
	"github.com/ola-silicon/krutrim-go-sdk/packages/param"
)

// DeleteVpcService contains methods and other services that help with interacting
// with the krutrim-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeleteVpcService] method instead.
type DeleteVpcService struct {
	Options []option.RequestOption
}

// NewDeleteVpcService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeleteVpcService(opts ...option.RequestOption) (r DeleteVpcService) {
	r = DeleteVpcService{}
	r.Options = opts
	return
}

// Delete a VPC
func (r *DeleteVpcService) Delete(ctx context.Context, params DeleteVpcDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XRegion) {
		opts = append(opts, option.WithHeader("x-region", fmt.Sprintf("%s", params.XRegion)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "/v1/highlvlvpc/delete_vpc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

type DeleteVpcDeleteParams struct {
	VpcID   string `query:"vpc_id,required" json:"-"`
	XRegion string `header:"x-region,required" json:"-"`
	paramObj
}

// URLQuery serializes [DeleteVpcDeleteParams]'s query parameters as `url.Values`.
func (r DeleteVpcDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
