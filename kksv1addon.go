package krutrim

import (
	"context"
	"net/http"

	"github.com/ola-krutrim/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
)

// KKSV1AddonService provides the catalog of add-ons available for installation.
// Use KKSV1ClusterAddonService for add-ons installed in a specific cluster.
type KKSV1AddonService struct {
	Options []option.RequestOption
}

func NewKKSV1AddonService(opts ...option.RequestOption) KKSV1AddonService {
	return KKSV1AddonService{Options: opts}
}

// List returns available add-ons, their versions, and cluster compatibilities.
func (r *KKSV1AddonService) List(ctx context.Context, query KKSV1AddonListParams, opts ...option.RequestOption) (res *KKSV1AddonListResponse, err error) {
	opts = kksRequestOptions(r.Options, query.XRegion, opts)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, "v1/kks/addons", nil, &res, opts...)
	return
}

type KKSV1AddonListParams struct {
	XRegion string `json:"-"`
}

type KKSAddonDefinition struct {
	AddonName     string            `json:"addonName"`
	Type          string            `json:"type"`
	Description   string            `json:"description"`
	IsDeletable   bool              `json:"isDeletable"`
	AddonVersions []KKSAddonVersion `json:"addonVersions"`
}

type KKSAddonVersion struct {
	AddonVersion    string                  `json:"addonVersion"`
	Architecture    []string                `json:"architecture"`
	Compatibilities []KKSAddonCompatibility `json:"compatibilities"`
}

type KKSAddonCompatibility struct {
	ClusterVersion   string   `json:"clusterVersion"`
	PlatformVersions []string `json:"platformVersions"`
	DefaultVersion   bool     `json:"defaultVersion"`
}

type KKSV1AddonListResponse struct {
	Addons []KKSAddonDefinition `json:"addons"`
}
