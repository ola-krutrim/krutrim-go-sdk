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

// DeleteSubnetService contains methods and other services that help with
// interacting with the krutrim-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeleteSubnetService] method instead.
type DeleteSubnetService struct {
	Options []option.RequestOption
}

// NewDeleteSubnetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeleteSubnetService(opts ...option.RequestOption) (r DeleteSubnetService) {
	r = DeleteSubnetService{}
	r.Options = opts
	return
}

// Deletes a specific subnet within a VPC using its Subnet ID (KRN).
func (r *DeleteSubnetService) Delete(ctx context.Context, params DeleteSubnetDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XRegion) {
		opts = append(opts, option.WithHeader("x-region", fmt.Sprintf("%s", params.XRegion)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "/v1/highlvlvpc/delete_subnet"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

type DeleteSubnetDeleteParams struct {
	// The Krutrim Resource Name (KRN) of the subnet to delete.
	SubnetID string `query:"subnet_id,required" json:"-"`
	VpcID    string `query:"vpc_id,required" json:"-"`
	XRegion  string `header:"x-region,required" json:"-"`
	paramObj
}

// URLQuery serializes [DeleteSubnetDeleteParams]'s query parameters as
// `url.Values`.
func (r DeleteSubnetDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
