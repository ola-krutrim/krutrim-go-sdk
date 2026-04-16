// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package krutrim

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/ola-krutrim/krutrim-go-sdk/internal/apiquery"
	shimjson "github.com/ola-krutrim/krutrim-go-sdk/internal/encoding/json"
	"github.com/ola-krutrim/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
)

// KcmV1CertTagService contains methods and other services that help with
// interacting with the krutrim API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKcmV1CertTagService] method instead.
type KcmV1CertTagService struct {
	Options []option.RequestOption
}

// NewKcmV1CertTagService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewKcmV1CertTagService(opts ...option.RequestOption) (r KcmV1CertTagService) {
	r = KcmV1CertTagService{}
	r.Options = opts
	return
}

// Add tags to certificate
func (r *KcmV1CertTagService) Add(ctx context.Context, certID string, body KcmV1CertTagAddParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if certID == "" {
		err = errors.New("missing required certId parameter")
		return
	}
	path := fmt.Sprintf("kcm/v1/certs/tags/%s", url.PathEscape(certID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get tag by name
func (r *KcmV1CertTagService) GetByName(ctx context.Context, query KcmV1CertTagGetByNameParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "kcm/v1/certs/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type KcmV1CertTagAddParams struct {
	Body map[string]string
	paramObj
}

func (r KcmV1CertTagAddParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *KcmV1CertTagAddParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type KcmV1CertTagGetByNameParams struct {
	CertID  string `query:"certId,required" json:"-"`
	TagName string `query:"tagName,required" json:"-"`
	paramObj
}

// URLQuery serializes [KcmV1CertTagGetByNameParams]'s query parameters as
// `url.Values`.
func (r KcmV1CertTagGetByNameParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
