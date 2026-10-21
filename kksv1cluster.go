package krutrim

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/ola-krutrim/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
	"github.com/ola-krutrim/krutrim-go-sdk/packages/param"
)

// KKSV1ClusterService manages KKS clusters and their resources.
type KKSV1ClusterService struct {
	Options   []option.RequestOption
	NodeGroup KKSV1ClusterNodeGroupService
	Addon     KKSV1ClusterAddonService
}

func NewKKSV1ClusterService(opts ...option.RequestOption) KKSV1ClusterService {
	return KKSV1ClusterService{
		Options:   opts,
		NodeGroup: NewKKSV1ClusterNodeGroupService(opts...),
		Addon:     NewKKSV1ClusterAddonService(opts...),
	}
}

// Create provisions a cluster with the requested network configuration and version.
func (r *KKSV1ClusterService) Create(ctx context.Context, body KKSV1ClusterCreateParams, opts ...option.RequestOption) (res *KKSV1ClusterCreateResponse, err error) {
	opts = kksRequestOptions(r.Options, body.XRegion, opts)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, "v1/kks/clusters", body, &res, opts...)
	return
}

// List returns clusters, optionally filtered by the region query parameter.
func (r *KKSV1ClusterService) List(ctx context.Context, query KKSV1ClusterListParams, opts ...option.RequestOption) (res *KKSV1ClusterListResponse, err error) {
	opts = kksRequestOptions(r.Options, query.XRegion, opts)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, "v1/kks/clusters", query, &res, opts...)
	return
}

// Get retrieves a cluster by its unescaped KRN.
func (r *KKSV1ClusterService) Get(ctx context.Context, clusterKRN string, query KKSV1ClusterGetParams, opts ...option.RequestOption) (res *KKSV1ClusterGetResponse, err error) {
	path, err := kksClusterPath(clusterKRN)
	if err != nil {
		return nil, err
	}
	opts = kksRequestOptions(r.Options, query.XRegion, opts)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete requests deletion of a cluster identified by its unescaped KRN.
func (r *KKSV1ClusterService) Delete(ctx context.Context, clusterKRN string, params KKSV1ClusterDeleteParams, opts ...option.RequestOption) error {
	path, err := kksClusterPath(clusterKRN)
	if err != nil {
		return err
	}
	opts = kksRequestOptions(r.Options, params.XRegion, opts)
	// Read and close the response without imposing a schema on an optional body.
	var response []byte
	return requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &response, opts...)
}

// GetKubeconfig downloads the response body unchanged, including YAML or JSON.
// The returned bytes may contain cluster credentials and should be stored securely.
func (r *KKSV1ClusterService) GetKubeconfig(ctx context.Context, clusterKRN string, query KKSV1ClusterGetKubeconfigParams, opts ...option.RequestOption) (res []byte, err error) {
	path, err := kksClusterPath(clusterKRN)
	if err != nil {
		return nil, err
	}
	defaults := append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, r.Options...)
	opts = kksRequestOptions(defaults, query.XRegion, opts)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path+"/kubeconfig", nil, &res, opts...)
	return
}

func kksClusterPath(clusterKRN string) (string, error) {
	if clusterKRN == "" {
		return "", errors.New("missing required clusterKRN parameter")
	}
	return "v1/kks/clusters/" + url.PathEscape(clusterKRN), nil
}

type KKSV1ClusterCreateParams struct {
	Name                    string                     `json:"name,required"`
	ResourcesVPCConfig      KKSResourcesVPCConfig      `json:"resourcesVpcConfig,required"`
	KubernetesNetworkConfig KKSKubernetesNetworkConfig `json:"kubernetesNetworkConfig,required"`
	Version                 string                     `json:"version,required"`
	XRegion                 string                     `json:"-"`
	paramObj
}

func (r KKSV1ClusterCreateParams) MarshalJSON() ([]byte, error) {
	type shadow KKSV1ClusterCreateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type KKSV1ClusterListParams struct {
	Region  param.Opt[string] `query:"region" json:"-"`
	XRegion string            `json:"-"`
	paramObj
}

// URLQuery encodes only the optional region filter; XRegion is a header.
func (r KKSV1ClusterListParams) URLQuery() (url.Values, error) {
	v := url.Values{}
	if !param.IsOmitted(r.Region) {
		v.Set("region", r.Region.Value)
	}
	return v, nil
}

type KKSV1ClusterGetParams struct {
	XRegion string `json:"-"`
}

type KKSV1ClusterDeleteParams struct {
	XRegion string `json:"-"`
}

type KKSV1ClusterGetKubeconfigParams struct {
	XRegion string `json:"-"`
}

type KKSResourcesVPCConfig struct {
	VPCKRN string `json:"vpcKrn"`
	// SubnetKRNs is a string, matching the KKS API rather than a JSON array.
	SubnetKRNs            string `json:"subnetKrns"`
	EndpointPrivateAccess bool   `json:"endpointPrivateAccess"`
}

type KKSKubernetesNetworkConfig struct {
	PodIPv4CIDR     string `json:"podIpv4Cidr"`
	ServiceIPv4CIDR string `json:"serviceIpv4Cidr"`
}

type KKSCluster struct {
	Name string `json:"name"`
	KRN  string `json:"krn"`
	// CreatedAt is Unix time in seconds, including the fractional seconds returned by KKS.
	CreatedAt               float64                    `json:"createdAt"`
	Version                 string                     `json:"version"`
	VersionSupported        bool                       `json:"versionSupported"`
	ResourcesVPCConfig      KKSResourcesVPCConfig      `json:"resourcesVpcConfig"`
	KubernetesNetworkConfig KKSKubernetesNetworkConfig `json:"kubernetesNetworkConfig"`
	Status                  string                     `json:"status"`
}

// KKSClusterSummary represents the different fields returned by the list endpoint.
type KKSClusterSummary struct {
	Name              string  `json:"name"`
	KRN               string  `json:"krn"`
	Status            string  `json:"status"`
	KubernetesVersion string  `json:"kubernetesVersion"`
	VersionSupported  bool    `json:"versionSupported"`
	NodegroupCount    int64   `json:"nodegroupCount"`
	CreatedAt         float64 `json:"createdAt"`
}

type KKSV1ClusterCreateResponse struct {
	Cluster KKSCluster `json:"cluster"`
}

type KKSV1ClusterGetResponse struct {
	Cluster KKSCluster `json:"cluster"`
}

type KKSV1ClusterListResponse struct {
	Clusters []KKSClusterSummary `json:"clusters"`
}
