// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package krutrim

import (
	"github.com/ola-krutrim/krutrim-go-sdk/option"
)

// KBV1Service contains methods and other services that help with interacting with
// the kbs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKBV1Service] method instead.
type KBV1Service struct {
	Options []option.RequestOption
	Volumes KBV1VolumeService
}

// NewKBV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewKBV1Service(opts ...option.RequestOption) (r KBV1Service) {
	r = KBV1Service{}
	r.Options = opts
	r.Volumes = NewKBV1VolumeService(opts...)
	return
}
