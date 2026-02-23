package krutrim

import (
	"github.com/ola-silicon/krutrim-go-sdk/option"
)

// VmV1Service contains methods and other services that help with interacting with
// the krutrim-vm-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVmV1Service] method instead.
type VmV1Service struct {
	Options []option.RequestOption
	Image   VmV1ImageService
}

// NewVmV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVmV1Service(opts ...option.RequestOption) (r VmV1Service) {
	r = VmV1Service{}
	r.Options = opts
	r.Image = NewVmV1ImageService(opts...)
	return
}
