// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package krutrim

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strconv"

	"github.com/ola-krutrim/krutrim-go-sdk/internal/apijson"
	"github.com/ola-krutrim/krutrim-go-sdk/internal/apiquery"
	"github.com/ola-krutrim/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
	"github.com/ola-krutrim/krutrim-go-sdk/packages/param"
)

// HighlvlvpcService contains methods and other services that help with interacting
// with the krutrim-vm-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHighlvlvpcService] method instead.
type HighlvlvpcService struct {
	Options []option.RequestOption
}

// NewHighlvlvpcService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHighlvlvpcService(opts ...option.RequestOption) (r HighlvlvpcService) {
	r = HighlvlvpcService{}
	r.Options = opts
	return
}

// Create a new VM instance in a VPC
func (r *HighlvlvpcService) NewInstance(ctx context.Context, body HighlvlvpcNewInstanceParams, opts ...option.RequestOption) (res *HighlvlvpcNewInstanceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "vm/v1/create_instance"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a VM instance
func (r *HighlvlvpcService) DeleteInstance(ctx context.Context, pathkrn string, deleteVolume bool, opts ...option.RequestOption) (res *HighlvlvpcDeleteInstanceResponse, err error) {

	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)

	path := "vm/v1/delete_instance"

	// Attach query params
	opts = append(opts, option.WithQuery("deleteVolume", strconv.FormatBool(deleteVolume)))
	opts = append(opts, option.WithQuery("instanceKrn", pathkrn))

	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get details of a specific VM instance by KRN
func (r *HighlvlvpcService) GetInstance(ctx context.Context, params HighlvlvpcGetInstanceParams, opts ...option.RequestOption) (res *HighlvlvpcGetInstanceResponse, err error) {
	if !param.IsOmitted(params.XRegion) {
		opts = append(opts, option.WithHeader("x-region", fmt.Sprintf("%s", params.XRegion)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/highlvlvpc/instance"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Search VM instances in a VPC
func (r *HighlvlvpcService) SearchInstances(ctx context.Context, query HighlvlvpcSearchInstancesParams, opts ...option.RequestOption) (res *HighlvlvpcSearchInstancesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "vm/v1/search_instances"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type HighlvlvpcNewInstanceResponse map[string]any

type HighlvlvpcDeleteInstanceResponse map[string]any

type HighlvlvpcGetInstanceResponse map[string]any

type HighlvlvpcSearchInstancesResponse map[string]any

type HighlvlvpcNewInstanceParams struct {
	// Required fields
	InstanceName string `json:"instanceName"`
	InstanceType string `json:"instanceType"`
	NetworkID    string `json:"network_id"`
	Region       string `json:"region"`
	SubnetID     string `json:"subnet_id"`
	VpcID        string `json:"vpc_id"`
	Volumetype   string `json:"volumetype"`
	VolumeName   string `json:"volume_name"`

	// Optional booleans
	IsGPU               *bool `json:"isGpu,omitempty"`
	FloatingIP          *bool `json:"floating_ip,omitempty"`
	DeleteOnTermination *bool `json:"delete_on_termination,omitempty"`

	// Optional strings
	ImageKrn   *string `json:"image_krn,omitempty"`
	PortKrn    *string `json:"port_krn,omitempty"`
	SshkeyName *string `json:"sshkey_name,omitempty"`
	UserData   *string `json:"user_data,omitempty"`

	// Optional numbers
	VolumeSize *int64 `json:"volume_size,omitempty"`

	// Optional collections
	SecurityGroups []string                         `json:"security_groups,omitempty"`
	Tags           []HighlvlvpcNewInstanceParamsTag `json:"tags,omitempty"`
	Volumes        []any                            `json:"volumes,omitempty"`
	paramObj
}

func (r HighlvlvpcNewInstanceParams) MarshalJSON() (data []byte, err error) {
	type shadow HighlvlvpcNewInstanceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HighlvlvpcNewInstanceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HighlvlvpcNewInstanceParamsTag struct {
	Key   param.Opt[string] `json:"key,omitzero"`
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r HighlvlvpcNewInstanceParamsTag) MarshalJSON() (data []byte, err error) {
	type shadow HighlvlvpcNewInstanceParamsTag
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HighlvlvpcNewInstanceParamsTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HighlvlvpcGetInstanceParams struct {
	// KRN identifier of the instance
	Krn     string `query:"krn,required" json:"-"`
	XRegion string `header:"x-region,required" json:"-"`
	paramObj
}

// URLQuery serializes [HighlvlvpcGetInstanceParams]'s query parameters as
// `url.Values`.
func (r HighlvlvpcGetInstanceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HighlvlvpcSearchInstancesParams struct {
	// VPC identifier (KRN encoded)
	VpcID string `query:"vpc_id,required" json:"-"`
	// Number of items per page
	Limit int64 `query:"limit" json:"-"`
	// Page number for pagination
	Page int64 `query:"page" json:"-"`
	paramObj
}

// URLQuery serializes [HighlvlvpcSearchInstancesParams]'s query parameters as
// `url.Values`.
func (r HighlvlvpcSearchInstancesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
type InstanceAction string

const (
	InstanceActionStart  InstanceAction = "start"
	InstanceActionStop   InstanceAction = "stop"
	InstanceActionReboot InstanceAction = "reboot"
)

type HighlvlvpcInstanceActionParams struct {
	Action InstanceAction `json:"action"`
	paramObj
}

func (r HighlvlvpcInstanceActionParams) MarshalJSON() ([]byte, error) {
	type shadow HighlvlvpcInstanceActionParams
	return param.MarshalObject(r, (*shadow)(&r))
}

func (r *HighlvlvpcInstanceActionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (r *HighlvlvpcService) InstanceAction(
	ctx context.Context,
	instanceKrn string,
	region string,
	body HighlvlvpcInstanceActionParams,
	opts ...option.RequestOption,
) (res []byte, err error) {

	opts = slices.Concat(r.Options, opts)

	opts = append(opts,
		option.WithHeader("Accept", "*/*"),
		option.WithHeader("x-region", region),
	)

	path := fmt.Sprintf("vm/v1/instance/%s", instanceKrn)

	err = requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodPut,
		path,
		body,
		&res, // ← IMPORTANT CHANGE
		opts...,
	)
	return
}

type VolumeAttachInput struct {
	InstanceID     string `json:"instanceId"`
	MountPartition string `json:"mountPartition"`
}

type VolumeAttachParams struct {
	Input VolumeAttachInput `json:"input"`
	paramObj
}

func (r VolumeAttachParams) MarshalJSON() ([]byte, error) {
	type shadow VolumeAttachParams
	return param.MarshalObject(r, (*shadow)(&r))
}

func (r *VolumeAttachParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (r *HighlvlvpcService) AttachVolume(
	ctx context.Context,
	volumeKrn string,
	tenantID string,
	body VolumeAttachParams,
	opts ...option.RequestOption,
) (res map[string]any, err error) {

	opts = slices.Concat(r.Options, opts)

	opts = append(opts,
		option.WithHeader("Accept", "*/*"),
		option.WithHeader("k-tenant-id", tenantID),
	)

	path := fmt.Sprintf("kbs/v1/volumes/%s/action", volumeKrn)
	opts = append(opts, option.WithQuery("op", "attach"))

	err = requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodPost,
		path,
		body,
		&res,
		opts...,
	)
	return
}
type VolumeDetachInput struct {
	InstanceID   string `json:"instanceId"`
	AttachmentID string `json:"attachment_id"`
}

type VolumeDetachParams struct {
	Input VolumeDetachInput `json:"input"`
	paramObj
}

func (r VolumeDetachParams) MarshalJSON() ([]byte, error) {
	type shadow VolumeDetachParams
	return param.MarshalObject(r, (*shadow)(&r))
}

func (r *VolumeDetachParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (r *HighlvlvpcService) DetachVolume(
	ctx context.Context,
	volumeKrn string,
	tenantID string,
	body VolumeDetachParams,
	opts ...option.RequestOption,
) (res map[string]any, err error) {

	opts = slices.Concat(r.Options, opts)

	opts = append(opts,
		option.WithHeader("Accept", "*/*"),
		option.WithHeader("k-tenant-id", tenantID),
	)

	path := fmt.Sprintf("kbs/v1/volumes/%s/action", volumeKrn)
	opts = append(opts, option.WithQuery("op", "detach"))

	err = requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodPost,
		path,
		body,
		&res,
		opts...,
	)
	return
}
