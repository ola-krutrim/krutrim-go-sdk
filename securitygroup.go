// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package krutrim

import (
	"github.com/ola-krutrim/krutrim-go-sdk/option"
)

// SecurityGroupService contains methods and other services that help with
// interacting with the sg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecurityGroupService] method instead.
type SecurityGroupService struct {
	Options []option.RequestOption
	V1      SecurityGroupV1Service
}

// NewSecurityGroupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSecurityGroupService(opts ...option.RequestOption) (r SecurityGroupService) {
	r = SecurityGroupService{}
	r.Options = opts
	r.V1 = NewSecurityGroupV1Service(opts...)
	return
}
