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

// DNSV1RecordService contains methods and other services that help with
// interacting with the dns-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSV1RecordService] method instead.
type DNSV1RecordService struct {
	Options []option.RequestOption
}

// NewDNSV1RecordService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDNSV1RecordService(opts ...option.RequestOption) (r DNSV1RecordService) {
	r = DNSV1RecordService{}
	r.Options = opts
	return
}

// Update Record
func (r *DNSV1RecordService) Update(ctx context.Context, recordid string, body DNSV1RecordUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if recordid == "" {
		err = errors.New("missing required recordid parameter")
		return
	}
	path := fmt.Sprintf("dns/v1/record/%s", url.PathEscape(recordid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Delete Record
func (r *DNSV1RecordService) Delete(ctx context.Context, recordid string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if recordid == "" {
		err = errors.New("missing required recordid parameter")
		return
	}
	path := fmt.Sprintf("dns/v1/record/%s", url.PathEscape(recordid))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// add record
func (r *DNSV1RecordService) Add(ctx context.Context, body DNSV1RecordAddParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "dns/v1/record"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type DNSV1RecordUpdateParams struct {
	Rname   param.Opt[string]               `json:"rname,omitzero"`
	Routing param.Opt[string]               `json:"routing,omitzero"`
	Type    param.Opt[string]               `json:"type,omitzero"`
	Records []DNSV1RecordUpdateParamsRecord `json:"records,omitzero"`
	paramObj
}

func (r DNSV1RecordUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DNSV1RecordUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DNSV1RecordUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DNSV1RecordUpdateParamsRecord struct {
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r DNSV1RecordUpdateParamsRecord) MarshalJSON() (data []byte, err error) {
	type shadow DNSV1RecordUpdateParamsRecord
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DNSV1RecordUpdateParamsRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DNSV1RecordAddParams struct {
	Krnid   string                       `json:"krnid" api:"required"`
	Records []DNSV1RecordAddParamsRecord `json:"records,omitzero" api:"required"`
	Rname   string                       `json:"rname" api:"required"`
	Ttl     int64                        `json:"ttl" api:"required"`
	Type    string                       `json:"type" api:"required"`
	paramObj
}

func (r DNSV1RecordAddParams) MarshalJSON() (data []byte, err error) {
	type shadow DNSV1RecordAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DNSV1RecordAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DNSV1RecordAddParamsRecord struct {
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r DNSV1RecordAddParamsRecord) MarshalJSON() (data []byte, err error) {
	type shadow DNSV1RecordAddParamsRecord
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DNSV1RecordAddParamsRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
