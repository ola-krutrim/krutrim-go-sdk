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

type CreateSubnetService struct {
	Options []option.RequestOption
}

func NewCreateSubnetService(opts ...option.RequestOption) (r CreateSubnetService) {
	r = CreateSubnetService{}
	r.Options = opts
	return
}

// Create Subnet
func (r *CreateSubnetService) New(
	ctx context.Context,
	body CreateSubnetNewParams,
	opts ...option.RequestOption,
) (err error) {

	opts = slices.Concat(r.Options, opts)

	opts = append([]option.RequestOption{
		option.WithHeader("Accept", "*/*"),
	}, opts...)

	path := "/v1/highlvlvpc/create_subnet"

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
// Models
// --------------------

// Subnet Data (Request Body)
type SubnetData struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	CIDR        string `json:"cidr"`
	GatewayIP   string `json:"gateway_ip"`
	NetworkID   string `json:"network_id"`
	IPVersion   string `json:"ip_version"`
	Ingress     bool   `json:"ingress"`
	Egress      bool   `json:"egress"`
}

// --------------------
// Request Params
// --------------------

type CreateSubnetNewParams struct {
	VpcID      string     `json:"vpc_id"`
	SubnetData SubnetData `json:"subnet_data"`

	paramObj
}

// --------------------
// JSON Handling
// --------------------

func (r CreateSubnetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CreateSubnetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

func (r *CreateSubnetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
