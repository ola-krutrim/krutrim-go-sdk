package krutrim

import (
	"context"
	"net/http"

	"github.com/ola-krutrim/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
)

// KKSV1FlavorService provides the instance flavors supported by KKS.
type KKSV1FlavorService struct {
	Options []option.RequestOption
}

func NewKKSV1FlavorService(opts ...option.RequestOption) KKSV1FlavorService {
	return KKSV1FlavorService{Options: opts}
}

// List returns the API's top-level array of CPU and GPU instance flavors.
func (r *KKSV1FlavorService) List(ctx context.Context, query KKSV1FlavorListParams, opts ...option.RequestOption) (res []KKSFlavor, err error) {
	opts = kksRequestOptions(r.Options, query.XRegion, opts)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, "v1/kks/flavors", nil, &res, opts...)
	return
}

type KKSV1FlavorListParams struct {
	XRegion string `json:"-"`
}

type KKSFlavor struct {
	Name      string `json:"name"`
	CPU       int64  `json:"cpu"`
	Memory    string `json:"memory"`
	GPUCount  int64  `json:"gpu_count"`
	GPUMemory int64  `json:"gpu_memory"`
}
