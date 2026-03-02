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
	"github.com/ola-krutrim/krutrim-go-sdk/internal/apiquery"
	"github.com/ola-krutrim/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
	"github.com/ola-krutrim/krutrim-go-sdk/packages/param"
)

// SshkeyService contains methods and other services that help with interacting
// with the sshkeys API.
//
// Note, unlike clients, this service does not read variables from the environment,
// automatically. You should not instantiate this service directly, and instead use
// the [NewSshkeyService] method instead.
type SshkeyService struct {
	Options []option.RequestOption
}

// NewSshkeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSshkeyService(opts ...option.RequestOption) (r SshkeyService) {
	r = SshkeyService{}
	r.Options = opts
	return
}

// Creates a new SSH public key for the authenticated customer.
func (r *SshkeyService) New(ctx context.Context, params SshkeyNewParams, opts ...option.RequestOption) (err error) {
	// if !param.IsOmitted(params.KCustomerID) {
	// 	opts = append(opts, option.WithHeader("k-customer-id", fmt.Sprintf("%s", params.KCustomerID)))
	// }
	if !param.IsOmitted(params.XRegion) {
		opts = append(opts, option.WithHeader("x-region", fmt.Sprintf("%s", params.XRegion)))
	}
	// if !param.IsOmitted(params.XRef) {
	// 	opts = append(opts, option.WithHeader("x-ref", fmt.Sprintf("%s", params.XRef.Value)))
	// }
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v2/sshkeys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Deletes an SSH key identified by its unique ID.
func (r *SshkeyService) Delete(ctx context.Context, sshKeyID string, opts ...option.RequestOption) (err error) {
	// if !param.IsOmitted(body.KCustomerID) {
	// 	opts = append(opts, option.WithHeader("k-customer-id", fmt.Sprintf("%s", body.KCustomerID)))
	// }
	// if !param.IsOmitted(body.XRegion) {
	// 	opts = append(opts, option.WithHeader("x-region", fmt.Sprintf("%s", body.XRegion)))
	// }
	// if !param.IsOmitted(body.XRef) {
	// 	opts = append(opts, option.WithHeader("x-ref", fmt.Sprintf("%s", body.XRef.Value)))
	// }
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if sshKeyID == "" {
		err = errors.New("missing required sshKeyId parameter")
		return
	}
	path := fmt.Sprintf("v2/sshkeys/%s", sshKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Returns SSH keys belonging to the authenticated customer filtered by key name.
func (r *SshkeyService) Search(ctx context.Context, params SshkeySearchParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v2/sshkeys/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return
}

type SshkeyNewParams struct {
	// Friendly name for the SSH key
	KeyName string `json:"keyName,required"`
	// SSH public key in OpenSSH format
	PublicKey string `json:"publicKey,required"`
	XRegion   string `header:"x-region,required" json:"-"`
	// UUID of the user who created the key
	// CreatedBy string `json:"createdBy,required"`
	paramObj
}

func (r SshkeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SshkeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SshkeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SshkeySearchParams struct {
	// Partial SSH key name to search for
	KeyName string `query:"keyName,required" json:"-"`
	paramObj
}

// URLQuery serializes [SshkeySearchParams]'s query parameters as `url.Values`.
func (r SshkeySearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
