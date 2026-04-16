// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package krutrim

import (
	"github.com/ola-krutrim/krutrim-go-sdk/option"
)

// KcmV1Service contains methods and other services that help with interacting with
// the krutrim API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKcmV1Service] method instead.
type KcmV1Service struct {
	Options []option.RequestOption
	Certs   KcmV1CertService
	CertTag KcmV1CertTagService
}

// NewKcmV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewKcmV1Service(opts ...option.RequestOption) (r KcmV1Service) {
	r = KcmV1Service{}
	r.Options = opts
	r.Certs = NewKcmV1CertService(opts...)
	r.CertTag = NewKcmV1CertTagService(opts...)
	return
}
