// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.
// ⚠️ Modified for better customer usability

package krutrim

import (
	"context"
	"net/http"
	"slices"

	"github.com/ola-krutrim/krutrim-go-sdk/internal/apijson"
	"github.com/ola-krutrim/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
	"github.com/ola-krutrim/krutrim-go-sdk/packages/param"
)

// --------------------
// Service
// --------------------

type CreateVpcAsyncService struct {
	Options []option.RequestOption
}

func NewCreateVpcAsyncService(opts ...option.RequestOption) (r CreateVpcAsyncService) {
	r = CreateVpcAsyncService{}
	r.Options = opts
	return
}

// Create VPC (Asynchronous)
func (r *CreateVpcAsyncService) New(
	ctx context.Context,
	body CreateVpcAsyncNewParams,
	opts ...option.RequestOption,
) (err error) {
	if body.XRegion != "" {
		opts = append(opts,
			option.WithHeader("x-region", body.XRegion),
		)
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{
		option.WithHeader("Accept", "*/*"),
	}, opts...)

	path := "/v1/highlvlvpc/create_vpc_async"

	err = requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodPost,
		path,
		body,
		nil,
		opts...,
	)

	return
}

// --------------------
// Strongly Typed Models
// --------------------

// VPC
type Vpc struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Enabled     bool   `json:"enabled"`
}

// Network
type Network struct {
	Name         string `json:"name"`
	AdminStateUp bool   `json:"admin_state_up"`
}

// Subnet
type Subnet struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	CIDR        string `json:"cidr"`
	GatewayIP   string `json:"gateway_ip"`
	IPVersion   string `json:"ip_version"`
	Ingress     bool   `json:"ingress"`
	Egress      bool   `json:"egress"`
}

// --------------------
// Request Params
// --------------------

type CreateVpcAsyncNewParams struct {
	XRegion string  `json:"-"`
	Network Network `json:"network"`
	Subnet  Subnet  `json:"subnet"`
	Vpc     Vpc     `json:"vpc"`
	paramObj
}

// --------------------
// JSON Handling
// --------------------

func (r CreateVpcAsyncNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CreateVpcAsyncNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

func (r *CreateVpcAsyncNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
