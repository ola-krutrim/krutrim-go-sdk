// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package krutrim

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/ola-krutrim/krutrim-go-sdk/internal/apiquery"
	"github.com/ola-krutrim/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
)

// DNSV1Service contains methods and other services that help with interacting with
// the dns-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSV1Service] method instead.
type DNSV1Service struct {
	Options []option.RequestOption
	Record  DNSV1RecordService
	Zone    DNSV1ZoneService
}

// NewDNSV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDNSV1Service(opts ...option.RequestOption) (r DNSV1Service) {
	r = DNSV1Service{}
	r.Options = opts
	r.Record = NewDNSV1RecordService(opts...)
	r.Zone = NewDNSV1ZoneService(opts...)
	return
}

// Get Records
func (r *DNSV1Service) GetRecords(ctx context.Context, zoneid string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if zoneid == "" {
		err = errors.New("missing required zoneid parameter")
		return
	}
	path := fmt.Sprintf("dns/v1/records/%s", url.PathEscape(zoneid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// getZones
func (r *DNSV1Service) GetZones(ctx context.Context, query DNSV1GetZonesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "dns/v1/zones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type DNSV1GetZonesParams struct {
	Region string `query:"region" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [DNSV1GetZonesParams]'s query parameters as `url.Values`.
func (r DNSV1GetZonesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
