// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

// GetVpcTaskStatusService contains methods and other services that help with
// interacting with the krutrim-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGetVpcTaskStatusService] method instead.
type GetVpcTaskStatusService struct {
	Options []option.RequestOption
}

// NewGetVpcTaskStatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGetVpcTaskStatusService(opts ...option.RequestOption) (r GetVpcTaskStatusService) {
	r = GetVpcTaskStatusService{}
	r.Options = opts
	return
}

// Get VPC Task Status
func (r *GetVpcTaskStatusService) New(ctx context.Context, body GetVpcTaskStatusNewParams, opts ...option.RequestOption) (err error) {
	if body.XRegion != "" {
		opts = append(opts,
			option.WithHeader("x-region", body.XRegion),
		)
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "/v1/highlvlvpc/get_vpc_task_status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type GetVpcTaskStatusNewParams struct {
	TaskID  string `json:"task_id"`
	XRegion string `json:"-"`
	paramObj
}

func (r GetVpcTaskStatusNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GetVpcTaskStatusNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetVpcTaskStatusNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
