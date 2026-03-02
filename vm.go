// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package krutrim

import (
	"github.com/ola-krutrim/krutrim-go-sdk/option"
)

// VmService contains methods and other services that help with interacting with
// the krutrim-vm-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVmService] method instead.
type VmService struct {
	Options []option.RequestOption
	V1      VmV1Service
}

// NewVmService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVmService(opts ...option.RequestOption) (r VmService) {
	r = VmService{}
	r.Options = opts
	r.V1 = NewVmV1Service(opts...)
	return
}
