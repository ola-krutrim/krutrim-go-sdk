// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package krutrim

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/ola-krutrim/krutrim-go-sdk/internal/apiform"
	"github.com/ola-krutrim/krutrim-go-sdk/internal/apiquery"
	"github.com/ola-krutrim/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
	"github.com/ola-krutrim/krutrim-go-sdk/packages/param"
)

// KcmV1CertService contains methods and other services that help with interacting
// with the krutrim API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKcmV1CertService] method instead.
type KcmV1CertService struct {
	Options []option.RequestOption
	Tags    KcmV1CertTagService
}

// NewKcmV1CertService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewKcmV1CertService(opts ...option.RequestOption) (r KcmV1CertService) {
	r = KcmV1CertService{}
	r.Options = opts
	r.Tags = NewKcmV1CertTagService(opts...)
	return
}

// Get certificate by ID
func (r *KcmV1CertService) Get(ctx context.Context, certID string, opts ...option.RequestOption) (res *KcmV1CertGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if certID == "" {
		err = errors.New("missing required certId parameter")
		return
	}
	// Fix: Use query parameter, not path parameter
	path := "kcm/v1/certs"
	query := url.Values{}
	query.Set("certId", certID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update certificate bundle
func (r *KcmV1CertService) Update(ctx context.Context, certID string, params KcmV1CertUpdateParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XVpcID) {
		opts = append(opts, option.WithHeader("X-Vpc-Id", fmt.Sprintf("%v", params.XVpcID)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if certID == "" {
		err = errors.New("missing required certId parameter")
		return
	}
	path := fmt.Sprintf("kcm/v1/certs/%s", url.PathEscape(certID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return
}

// List certificates or get certificate by query
func (r *KcmV1CertService) List(ctx context.Context, query KcmV1CertListParams, opts ...option.RequestOption) (res *KcmV1CertGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "kcm/v1/certs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete certificate
func (r *KcmV1CertService) Delete(ctx context.Context, certID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if certID == "" {
		err = errors.New("missing required certId parameter")
		return
	}
	path := fmt.Sprintf("kcm/v1/certs/%s", url.PathEscape(certID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get expiring certificates
func (r *KcmV1CertService) GetExpiring(ctx context.Context, query KcmV1CertGetExpiringParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "kcm/v1/certs/expiringIn"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Import certificate
func (r *KcmV1CertService) Import(ctx context.Context, params KcmV1CertImportParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XVpcID) {
		opts = append(opts, option.WithHeader("X-Vpc-Id", fmt.Sprintf("%v", params.XVpcID)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "kcm/v1/certs/import"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

type KcmV1CertUpdateParams struct {
	XVpcID   string           `header:"X-Vpc-Id,required" json:"-"`
	Flag     param.Opt[int64] `json:"flag,omitzero"`
	CertFile io.Reader        `json:"certFile,omitzero" format:"binary"`
	paramObj
}

func (r KcmV1CertUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type KcmV1CertListParams struct {
	CertID param.Opt[string] `query:"certId,omitzero" json:"-"`
	Lbtype param.Opt[int64]  `query:"lbtype,omitzero" json:"-"`
	VpcID  param.Opt[string] `query:"vpcId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [KcmV1CertListParams]'s query parameters as `url.Values`.
func (r KcmV1CertListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type KcmV1CertGetExpiringParams struct {
	Date  time.Time `query:"date,required" format:"date-time" json:"-"`
	VpcID string    `query:"vpcId,required" json:"-"`
	paramObj
}

// URLQuery serializes [KcmV1CertGetExpiringParams]'s query parameters as
// `url.Values`.
func (r KcmV1CertGetExpiringParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type KcmV1CertImportParams struct {
	CertFile io.Reader         `json:"certFile,omitzero,required" format:"binary"`
	Name     string            `json:"name,required"`
	XVpcID   string            `header:"X-Vpc-Id,required" json:"-"`
	Flag     param.Opt[int64]  `json:"flag,omitzero"`
	Tags     map[string]string `json:"tags,omitzero"`
	paramObj
}

func (r KcmV1CertImportParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type KcmV1CertGetResponse struct {
	Message      string                 `json:"message"`
	Certificate  string                 `json:"certificate"`
	Krn          string                 `json:"krn"`
	Expiration   string                 `json:"expiration"`
	CommonName   string                 `json:"commonName"`
	SANs         []string               `json:"SANs"`
	SerialNumber string                 `json:"Serial Number"`
	Issuer       string                 `json:"issuer"`
	SignatureAlg string                 `json:"Signature Algorithm"`
	TenantID     string                 `json:"tenantId"`
	Tags         map[string]interface{} `json:"tags"`
}
