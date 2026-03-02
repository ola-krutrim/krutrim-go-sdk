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

// DescribeVpcService contains methods and other services that help with
// interacting with the krutrim-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDescribeVpcService] method instead.
type DescribeVpcService struct {
	Options []option.RequestOption
}

// NewDescribeVpcService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDescribeVpcService(opts ...option.RequestOption) (r DescribeVpcService) {
	r = DescribeVpcService{}
	r.Options = opts
	return
}

// Describe VPC Details
func (r *DescribeVpcService) Get(ctx context.Context, params DescribeVpcGetParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XRegion) {
		opts = append(opts, option.WithHeader("x-region", fmt.Sprintf("%s", params.XRegion)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "/v1/highlvlvpc/describe_vpc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return
}

type DescribeVpcGetParams struct {
	VpcID   string `query:"vpc_id,required" json:"-"`
	VpcName string `query:"vpc_name,required" json:"-"`
	XRegion string `header:"x-region,required" json:"-"`
	paramObj
}

// URLQuery serializes [DescribeVpcGetParams]'s query parameters as `url.Values`.
func (r DescribeVpcGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
