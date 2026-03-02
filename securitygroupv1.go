// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package krutrim

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/ola-krutrim/krutrim-go-sdk/internal/apijson"
	"github.com/ola-krutrim/krutrim-go-sdk/internal/apiquery"
	"github.com/ola-krutrim/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
	"github.com/ola-krutrim/krutrim-go-sdk/packages/param"
	"github.com/ola-krutrim/krutrim-go-sdk/packages/respjson"
)

// SecurityGroupV1Service contains methods and other services that help with
// interacting with the sg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecurityGroupV1Service] method instead.
type SecurityGroupV1Service struct {
	Options []option.RequestOption
}

// NewSecurityGroupV1Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSecurityGroupV1Service(opts ...option.RequestOption) (r SecurityGroupV1Service) {
	r = SecurityGroupV1Service{}
	r.Options = opts
	return
}

// Create a security group
func (r *SecurityGroupV1Service) New(ctx context.Context, params SecurityGroupV1NewParams, opts ...option.RequestOption) (res *SecurityGroup, err error) {
	if !param.IsOmitted(params.XRegion) {
		opts = append(opts, option.WithHeader("x-region", fmt.Sprintf("%s", params.XRegion)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "securityGroup/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List security groups for a VPC
func (r *SecurityGroupV1Service) List(ctx context.Context, vpcKrn string, params SecurityGroupV1ListParams, opts ...option.RequestOption) (res *SecurityGroupV1ListResponse, err error) {
	if !param.IsOmitted(params.XRegion) {
		opts = append(opts, option.WithHeader("x-region", fmt.Sprintf("%s", params.XRegion)))
	}
	opts = slices.Concat(r.Options, opts)
	if vpcKrn == "" {
		err = errors.New("missing required vpcKrn parameter")
		return
	}
	path := fmt.Sprintf("securityGroup/v1/%s", url.PathEscape(vpcKrn))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete a security group
func (r *SecurityGroupV1Service) Delete(ctx context.Context, securityGroupKrn string, body SecurityGroupV1DeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XRegion) {
		opts = append(opts, option.WithHeader("x-region", fmt.Sprintf("%s", body.XRegion)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if securityGroupKrn == "" {
		err = errors.New("missing required securityGroupKrn parameter")
		return
	}
	path := fmt.Sprintf("securityGroup/v1/%s", url.PathEscape(securityGroupKrn))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Attach rule to a security group
func (r *SecurityGroupV1Service) AttachRule(ctx context.Context, params SecurityGroupV1AttachRuleParams, opts ...option.RequestOption) (res *SecurityGroup, err error) {
	if !param.IsOmitted(params.XRegion) {
		opts = append(opts, option.WithHeader("x-region", fmt.Sprintf("%s", params.XRegion)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "securityGroup/v1/attachrule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Create a security group rule
func (r *SecurityGroupV1Service) NewRule(ctx context.Context, params SecurityGroupV1NewRuleParams, opts ...option.RequestOption) (res *SecurityGroupV1NewRuleResponse, err error) {
	if !param.IsOmitted(params.XRegion) {
		opts = append(opts, option.WithHeader("x-region", fmt.Sprintf("%s", params.XRegion)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "securityGroup/v1/rule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Detach rule from a security group
func (r *SecurityGroupV1Service) DetachRule(ctx context.Context, params SecurityGroupV1DetachRuleParams, opts ...option.RequestOption) (res *SecurityGroup, err error) {
	if !param.IsOmitted(params.XRegion) {
		opts = append(opts, option.WithHeader("x-region", fmt.Sprintf("%s", params.XRegion)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "securityGroup/v1/deattachrule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type SecurityGroup struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt" format:"date-time"`
	Description string    `json:"description"`
	Name        string    `json:"name"`
	UpdatedAt   time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityGroup) RawJSON() string { return r.JSON.raw }
func (r *SecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityGroupV1ListResponse struct {
	Items []SecurityGroup `json:"items"`
	Limit int64           `json:"limit"`
	Page  int64           `json:"page"`
	Total int64           `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Limit       respjson.Field
		Page        respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityGroupV1ListResponse) RawJSON() string { return r.JSON.raw }
func (r *SecurityGroupV1ListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityGroupV1NewRuleResponse struct {
	ID             string    `json:"id"`
	CreatedAt      time.Time `json:"createdAt" format:"date-time"`
	Description    string    `json:"description"`
	Direction      string    `json:"direction"`
	Ethertypes     string    `json:"ethertypes"`
	PortMaxRange   int64     `json:"portMaxRange"`
	PortMinRange   int64     `json:"portMinRange"`
	Protocol       string    `json:"protocol"`
	RemoteIPPrefix string    `json:"remoteIPPrefix"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Description    respjson.Field
		Direction      respjson.Field
		Ethertypes     respjson.Field
		PortMaxRange   respjson.Field
		PortMinRange   respjson.Field
		Protocol       respjson.Field
		RemoteIPPrefix respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityGroupV1NewRuleResponse) RawJSON() string { return r.JSON.raw }
func (r *SecurityGroupV1NewRuleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityGroupV1NewParams struct {
	Name        string            `json:"name,required"`
	Vpcid       string            `json:"vpcid,required"`
	XRegion     string            `header:"x-region,required" json:"-"`
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r SecurityGroupV1NewParams) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupV1NewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SecurityGroupV1NewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityGroupV1ListParams struct {
	XRegion string            `header:"x-region,required" json:"-"`
	Limit   param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Page    param.Opt[int64]  `query:"page,omitzero" json:"-"`
	SortBy  param.Opt[string] `query:"sortBy,omitzero" json:"-"`
	// Any of "asc", "desc".
	SortOrder SecurityGroupV1ListParamsSortOrder `query:"sortOrder,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SecurityGroupV1ListParams]'s query parameters as
// `url.Values`.
func (r SecurityGroupV1ListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SecurityGroupV1ListParamsSortOrder string

const (
	SecurityGroupV1ListParamsSortOrderAsc  SecurityGroupV1ListParamsSortOrder = "asc"
	SecurityGroupV1ListParamsSortOrderDesc SecurityGroupV1ListParamsSortOrder = "desc"
)

type SecurityGroupV1DeleteParams struct {
	XRegion string `header:"x-region,required" json:"-"`
	paramObj
}

type SecurityGroupV1AttachRuleParams struct {
	Ruleid     string `json:"ruleid,required"`
	Securityid string `json:"securityid,required"`
	Vpcid      string `json:"vpcid,required"`
	XRegion    string `header:"x-region,required" json:"-"`
	paramObj
}

func (r SecurityGroupV1AttachRuleParams) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupV1AttachRuleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SecurityGroupV1AttachRuleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityGroupV1NewRuleParams struct {
	// Any of "ingress", "egress".
	Direction SecurityGroupV1NewRuleParamsDirection `json:"direction,required"`
	// Any of "IPv4", "IPv6".
	Ethertypes     SecurityGroupV1NewRuleParamsEthertypes `json:"ethertypes,required"`
	Protocol       string                                 `json:"protocol,required"`
	Vpcid          string                                 `json:"vpcid,required"`
	XRegion        string                                 `header:"x-region,required" json:"-"`
	Description    param.Opt[string]                      `json:"description,omitzero"`
	PortMaxRange   int64                                  `json:"portMaxRange"`
	PortMinRange   int64                                  `json:"portMinRange"`
	RemoteIPPrefix string                                 `json:"remoteIPPrefix"`
	paramObj
}

func (r SecurityGroupV1NewRuleParams) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupV1NewRuleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SecurityGroupV1NewRuleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityGroupV1NewRuleParamsDirection string

const (
	SecurityGroupV1NewRuleParamsDirectionIngress SecurityGroupV1NewRuleParamsDirection = "ingress"
	SecurityGroupV1NewRuleParamsDirectionEgress  SecurityGroupV1NewRuleParamsDirection = "egress"
)

type SecurityGroupV1NewRuleParamsEthertypes string

const (
	SecurityGroupV1NewRuleParamsEthertypesIPv4 SecurityGroupV1NewRuleParamsEthertypes = "IPv4"
	SecurityGroupV1NewRuleParamsEthertypesIPv6 SecurityGroupV1NewRuleParamsEthertypes = "IPv6"
)

type SecurityGroupV1DetachRuleParams struct {
	Ruleid     string `json:"ruleid,required"`
	Securityid string `json:"securityid,required"`
	Vpcid      string `json:"vpcid,required"`
	XRegion    string `header:"x-region,required" json:"-"`
	paramObj
}

func (r SecurityGroupV1DetachRuleParams) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupV1DetachRuleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SecurityGroupV1DetachRuleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
