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

// KBV1VolumeService contains methods and other services that help with interacting
// with the kbs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKBV1VolumeService] method instead.
type KBV1VolumeService struct {
	Options []option.RequestOption
}

// NewKBV1VolumeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewKBV1VolumeService(opts ...option.RequestOption) (r KBV1VolumeService) {
	r = KBV1VolumeService{}
	r.Options = opts
	return
}

// Create a new volume
func (r *KBV1VolumeService) New(ctx context.Context, params KBV1VolumeNewParams, opts ...option.RequestOption) (res *Kbv1VolumeNewResponse, err error) {
	if !param.IsOmitted(params.KTenantID) {
		opts = append(opts, option.WithHeader("k-tenant-id", fmt.Sprintf("%s", params.KTenantID)))
	}
	if !param.IsOmitted(params.XAccountID) {
		opts = append(opts, option.WithHeader("x-account-id", fmt.Sprintf("%s", params.XAccountID)))
	}
	params.AvailabilityZone = "nova"
	opts = slices.Concat(r.Options, opts)
	path := "kbs/v1/volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get volume details
func (r *KBV1VolumeService) Get(ctx context.Context, volumeID string, query KBV1VolumeGetParams, opts ...option.RequestOption) (res *Kbv1VolumeGetResponse, err error) {
	if !param.IsOmitted(query.KTenantID) {
		opts = append(opts, option.WithHeader("k-tenant-id", fmt.Sprintf("%s", query.KTenantID)))
	}
	opts = slices.Concat(r.Options, opts)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("kbs/v1/volumes/%s", url.PathEscape(volumeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a volume
func (r *KBV1VolumeService) Delete(ctx context.Context, volumeID string, body KBV1VolumeDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.KTenantID) {
		opts = append(opts, option.WithHeader("k-tenant-id", fmt.Sprintf("%s", body.KTenantID)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("kbs/v1/volumes/%s", url.PathEscape(volumeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Kbv1VolumeNewResponse = any

type Kbv1VolumeGetResponse = any

type KBV1VolumeNewParams struct {
	AvailabilityZone string                    `json:"availability_zone,required"`
	Name             string                    `json:"name,required"`
	Size             int64                     `json:"size,required"`
	Source           KBV1VolumeNewParamsSource `json:"source,omitzero,required"`
	Volumetype       string                    `json:"volumetype,required"`
	KTenantID        string                    `header:"k-tenant-id,required" json:"-"`
	XAccountID       string                    `header:"x-account-id,required" json:"-"`
	Description      param.Opt[string]         `json:"description,omitzero"`
	Multiattach      param.Opt[bool]           `json:"multiattach,omitzero"`
	Metadata         map[string]string         `json:"metadata,omitzero"`
	paramObj
}

func (r KBV1VolumeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow KBV1VolumeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KBV1VolumeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KBV1VolumeNewParamsSource struct {
	ID   param.Opt[string] `json:"id,omitzero"`
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r KBV1VolumeNewParamsSource) MarshalJSON() (data []byte, err error) {
	type shadow KBV1VolumeNewParamsSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KBV1VolumeNewParamsSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KBV1VolumeGetParams struct {
	KTenantID string `header:"k-tenant-id,required" json:"-"`
	paramObj
}

type KBV1VolumeDeleteParams struct {
	KTenantID string `header:"k-tenant-id,required" json:"-"`
	paramObj
}
