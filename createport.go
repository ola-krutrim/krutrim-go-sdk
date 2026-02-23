// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package krutrim

import (
	"context"
	"net/http"
	"slices"

	"github.com/ola-silicon/krutrim-go-sdk/internal/apijson"

	"github.com/ola-silicon/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-silicon/krutrim-go-sdk/option"
	"github.com/ola-silicon/krutrim-go-sdk/packages/param"
)

// CreatePortService contains methods and other services that help with interacting
// with the krutrim-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCreatePortService] method instead.
type CreatePortService struct {
	Options []option.RequestOption
}

// NewCreatePortService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCreatePortService(opts ...option.RequestOption) (r CreatePortService) {
	r = CreatePortService{}
	r.Options = opts
	return
}

// Create Port (Reserve IP)
func (r *CreatePortService) New(ctx context.Context, body CreatePortNewParams, opts ...option.RequestOption) (err error) {
	if body.XRegion != "" {
		opts = append(opts,
			option.WithHeader("x-region", body.XRegion),
		)
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "/v1/highlvlvpc/create_port"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type CreatePortNewParams struct {
	XRegion    string `json:"-"`
	FloatingIP bool   `json:"floating_ip,omitzero"`
	Name       string `json:"name,omitzero"`
	NetworkID  string `json:"network_id,omitzero"`
	SubnetID   string `json:"subnet_id,omitzero"`
	VpcID      string `json:"vpc_id,omitzero"`
	paramObj
}

func (r CreatePortNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CreatePortNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatePortNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
