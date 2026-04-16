// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package krutrim

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/ola-krutrim/krutrim-go-sdk/internal/apijson"
	"github.com/ola-krutrim/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
	"github.com/ola-krutrim/krutrim-go-sdk/packages/param"
)

// DNSV1ZoneVpcService contains methods and other services that help with
// interacting with the dns-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSV1ZoneVpcService] method instead.
type DNSV1ZoneVpcService struct {
	Options []option.RequestOption
}

// NewDNSV1ZoneVpcService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDNSV1ZoneVpcService(opts ...option.RequestOption) (r DNSV1ZoneVpcService) {
	r = DNSV1ZoneVpcService{}
	r.Options = opts
	return
}

// add vpc
func (r *DNSV1ZoneVpcService) Add(ctx context.Context, zoneid string, body DNSV1ZoneVpcAddParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if zoneid == "" {
		err = errors.New("missing required zoneid parameter")
		return
	}
	path := fmt.Sprintf("dns/v1/zone/%s/vpc", url.PathEscape(zoneid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Remove VPC from Private Zone
func (r *DNSV1ZoneVpcService) Remove(ctx context.Context, zoneid string, body DNSV1ZoneVpcRemoveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if zoneid == "" {
		err = errors.New("missing required zoneid parameter")
		return
	}
	path := fmt.Sprintf("dns/v1/zone/%s/vpc", url.PathEscape(zoneid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

type DNSV1ZoneVpcAddParams struct {
	Vpcinfo []string `json:"vpcinfo,omitzero"`
	paramObj
}

func (r DNSV1ZoneVpcAddParams) MarshalJSON() (data []byte, err error) {
	type shadow DNSV1ZoneVpcAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DNSV1ZoneVpcAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DNSV1ZoneVpcRemoveParams struct {
	Vpcinfo []string `json:"vpcinfo,omitzero"`
	paramObj
}

func (r DNSV1ZoneVpcRemoveParams) MarshalJSON() (data []byte, err error) {
	type shadow DNSV1ZoneVpcRemoveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DNSV1ZoneVpcRemoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
