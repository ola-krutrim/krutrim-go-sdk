// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package krutrim

import (
	"github.com/ola-krutrim/krutrim-go-sdk/option"
)

// KcmService contains methods and other services that help with interacting with
// the kcertmangment API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKcmService] method instead.
type KcmService struct {
	Options []option.RequestOption
	V1      KcmV1Service
}

// NewKcmService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewKcmService(opts ...option.RequestOption) (r KcmService) {
	r = KcmService{}
	r.Options = opts
	r.V1 = NewKcmV1Service(opts...)
	return
}
