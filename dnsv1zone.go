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

// DNSV1ZoneService contains methods and other services that help with interacting
// with the dns-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSV1ZoneService] method instead.
type DNSV1ZoneService struct {
	Options []option.RequestOption
	Vpc     DNSV1ZoneVpcService
}

// NewDNSV1ZoneService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDNSV1ZoneService(opts ...option.RequestOption) (r DNSV1ZoneService) {
	r = DNSV1ZoneService{}
	r.Options = opts
	r.Vpc = NewDNSV1ZoneVpcService(opts...)
	return
}

// Delete Zone
func (r *DNSV1ZoneService) Delete(ctx context.Context, zoneid string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if zoneid == "" {
		err = errors.New("missing required zoneid parameter")
		return
	}
	path := fmt.Sprintf("dns/v1/zone/%s", url.PathEscape(zoneid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// add zone
func (r *DNSV1ZoneService) Add(ctx context.Context, body DNSV1ZoneAddParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "dns/v1/zone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get Zone by ID
func (r *DNSV1ZoneService) Get(ctx context.Context, zoneid string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if zoneid == "" {
		err = errors.New("missing required zoneid parameter")
		return
	}
	path := fmt.Sprintf("dns/v1/zone/%s", url.PathEscape(zoneid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

type DNSV1ZoneAddParams struct {
	Type     string `json:"type" api:"required"`
	Vpcid    string `json:"vpcid" api:"required"`
	Zonename string `json:"zonename" api:"required"`
	paramObj
}

func (r DNSV1ZoneAddParams) MarshalJSON() (data []byte, err error) {
	type shadow DNSV1ZoneAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DNSV1ZoneAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
