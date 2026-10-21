package krutrim

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/ola-krutrim/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
	"github.com/ola-krutrim/krutrim-go-sdk/packages/param"
)

// KKSV1ClusterAddonService manages add-ons within a cluster.
type KKSV1ClusterAddonService struct {
	Options []option.RequestOption
}

func NewKKSV1ClusterAddonService(opts ...option.RequestOption) KKSV1ClusterAddonService {
	return KKSV1ClusterAddonService{Options: opts}
}

// Create installs an add-on. Omit AddonVersion to use the server's default version.
func (r *KKSV1ClusterAddonService) Create(ctx context.Context, clusterKRN string, body KKSV1ClusterAddonCreateParams, opts ...option.RequestOption) (res *KKSV1ClusterAddonCreateResponse, err error) {
	path, err := kksClusterPath(clusterKRN)
	if err != nil {
		return nil, err
	}
	opts = kksRequestOptions(r.Options, body.XRegion, opts)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path+"/addons", body, &res, opts...)
	return
}

// List returns the add-ons installed in a cluster identified by its unescaped KRN.
func (r *KKSV1ClusterAddonService) List(ctx context.Context, clusterKRN string, query KKSV1ClusterAddonListParams, opts ...option.RequestOption) (res *KKSV1ClusterAddonListResponse, err error) {
	path, err := kksClusterPath(clusterKRN)
	if err != nil {
		return nil, err
	}
	opts = kksRequestOptions(r.Options, query.XRegion, opts)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path+"/addons", nil, &res, opts...)
	return
}

// Delete requests removal of an installed add-on using its unescaped KRN.
func (r *KKSV1ClusterAddonService) Delete(ctx context.Context, clusterKRN, addonKRN string, params KKSV1ClusterAddonDeleteParams, opts ...option.RequestOption) error {
	path, err := kksClusterPath(clusterKRN)
	if err != nil {
		return err
	}
	if addonKRN == "" {
		return errors.New("missing required addonKRN parameter")
	}
	opts = kksRequestOptions(r.Options, params.XRegion, opts)
	var response []byte
	return requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path+"/addons/"+url.PathEscape(addonKRN), nil, &response, opts...)
}

type KKSV1ClusterAddonCreateParams struct {
	AddonName    string            `json:"addonName,required"`
	AddonVersion param.Opt[string] `json:"addonVersion,omitzero"`
	XRegion      string            `json:"-"`
	paramObj
}

func (r KKSV1ClusterAddonCreateParams) MarshalJSON() ([]byte, error) {
	type shadow KKSV1ClusterAddonCreateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type KKSV1ClusterAddonListParams struct {
	XRegion string `json:"-"`
}

type KKSV1ClusterAddonDeleteParams struct {
	XRegion string `json:"-"`
}

type KKSAddon struct {
	AddonKRN     string    `json:"addonKrn"`
	AddonName    string    `json:"addonName"`
	AddonVersion string    `json:"addonVersion"`
	ClusterName  string    `json:"clusterName"`
	CreatedAt    time.Time `json:"createdAt"`
	Status       string    `json:"status"`
}

type KKSV1ClusterAddonCreateResponse struct {
	Addon KKSAddon `json:"addon"`
}

// KKSAddonSummary represents an installed add-on. Unlike the create response,
// the list endpoint returns timestamps as Unix seconds instead of ISO strings.
type KKSAddonSummary struct {
	AddonKRN     string  `json:"addonKrn"`
	AddonName    string  `json:"addonName"`
	AddonVersion string  `json:"addonVersion"`
	Status       string  `json:"status"`
	Type         string  `json:"type"`
	CreatedAt    float64 `json:"createdAt"`
	ModifiedAt   float64 `json:"modifiedAt"`
}

type KKSV1ClusterAddonListResponse struct {
	Addons []KKSAddonSummary `json:"addons"`
}
