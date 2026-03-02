// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package krutrim

import (
	"github.com/ola-krutrim/krutrim-go-sdk/option"
)

// KBService contains methods and other services that help with interacting with
// the kbs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKBService] method instead.
type KBService struct {
	Options []option.RequestOption
	V1      KBV1Service
}

// NewKBService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewKBService(opts ...option.RequestOption) (r KBService) {
	r = KBService{}
	r.Options = opts
	r.V1 = NewKBV1Service(opts...)
	return
}
