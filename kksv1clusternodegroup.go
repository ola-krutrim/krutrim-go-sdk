package krutrim

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/ola-krutrim/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
	"github.com/ola-krutrim/krutrim-go-sdk/packages/param"
)

// KKSV1ClusterNodeGroupService manages node groups within a cluster.
type KKSV1ClusterNodeGroupService struct {
	Options []option.RequestOption
}

func NewKKSV1ClusterNodeGroupService(opts ...option.RequestOption) KKSV1ClusterNodeGroupService {
	return KKSV1ClusterNodeGroupService{Options: opts}
}

// Create provisions a node group in the cluster identified by its unescaped KRN.
func (r *KKSV1ClusterNodeGroupService) Create(ctx context.Context, clusterKRN string, body KKSV1ClusterNodeGroupCreateParams, opts ...option.RequestOption) (res *KKSV1ClusterNodeGroupCreateResponse, err error) {
	path, err := kksClusterPath(clusterKRN)
	if err != nil {
		return nil, err
	}
	opts = kksRequestOptions(r.Options, body.XRegion, opts)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path+"/node-groups", body, &res, opts...)
	return
}

// List returns the cluster's node groups, including an empty list when none exist.
func (r *KKSV1ClusterNodeGroupService) List(ctx context.Context, clusterKRN string, query KKSV1ClusterNodeGroupListParams, opts ...option.RequestOption) (res *KKSV1ClusterNodeGroupListResponse, err error) {
	path, err := kksClusterPath(clusterKRN)
	if err != nil {
		return nil, err
	}
	opts = kksRequestOptions(r.Options, query.XRegion, opts)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path+"/node-groups", nil, &res, opts...)
	return
}

// Get retrieves a node group using the cluster and node group's unescaped KRNs.
func (r *KKSV1ClusterNodeGroupService) Get(ctx context.Context, clusterKRN, nodeGroupKRN string, query KKSV1ClusterNodeGroupGetParams, opts ...option.RequestOption) (res *KKSV1ClusterNodeGroupGetResponse, err error) {
	path, err := kksNodeGroupPath(clusterKRN, nodeGroupKRN)
	if err != nil {
		return nil, err
	}
	opts = kksRequestOptions(r.Options, query.XRegion, opts)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// UpdateConfig updates a node group's scaling configuration using unescaped KRNs.
func (r *KKSV1ClusterNodeGroupService) UpdateConfig(ctx context.Context, clusterKRN, nodeGroupKRN string, body KKSV1ClusterNodeGroupUpdateConfigParams, opts ...option.RequestOption) (res *KKSV1ClusterNodeGroupUpdateConfigResponse, err error) {
	path, err := kksNodeGroupPath(clusterKRN, nodeGroupKRN)
	if err != nil {
		return nil, err
	}
	opts = kksRequestOptions(r.Options, body.XRegion, opts)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path+"/update-config", body, &res, opts...)
	return
}

// Delete requests deletion of a node group using its unescaped KRN.
func (r *KKSV1ClusterNodeGroupService) Delete(ctx context.Context, clusterKRN, nodeGroupKRN string, params KKSV1ClusterNodeGroupDeleteParams, opts ...option.RequestOption) error {
	path, err := kksNodeGroupPath(clusterKRN, nodeGroupKRN)
	if err != nil {
		return err
	}
	opts = kksRequestOptions(r.Options, params.XRegion, opts)
	var response []byte
	return requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &response, opts...)
}

func kksNodeGroupPath(clusterKRN, nodeGroupKRN string) (string, error) {
	path, err := kksClusterPath(clusterKRN)
	if err != nil {
		return "", err
	}
	if nodeGroupKRN == "" {
		return "", errors.New("missing required nodeGroupKRN parameter")
	}
	return path + "/node-groups/" + url.PathEscape(nodeGroupKRN), nil
}

type KKSV1ClusterNodeGroupCreateParams struct {
	Name             string               `json:"name,required"`
	InstanceTypes    string               `json:"instanceTypes,required"`
	DiskSize         int64                `json:"diskSize,required"`
	ScalingConfig    KKSScalingConfig     `json:"scalingConfig,required"`
	SubnetsKRN       string               `json:"subnetsKrn,required"`
	Labels           map[string]string    `json:"labels,omitzero"`
	Taints           []KKSTaint           `json:"taints,omitzero"`
	RemoteAccess     *KKSRemoteAccess     `json:"remoteAccess,omitzero"`
	NodeRepairConfig *KKSNodeRepairConfig `json:"nodeRepairConfig,omitzero"`
	XRegion          string               `json:"-"`
	paramObj
}

func (r KKSV1ClusterNodeGroupCreateParams) MarshalJSON() ([]byte, error) {
	type shadow KKSV1ClusterNodeGroupCreateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type KKSV1ClusterNodeGroupUpdateConfigParams struct {
	ScalingConfig KKSScalingConfig `json:"scalingConfig,required"`
	XRegion       string           `json:"-"`
	paramObj
}

func (r KKSV1ClusterNodeGroupUpdateConfigParams) MarshalJSON() ([]byte, error) {
	type shadow KKSV1ClusterNodeGroupUpdateConfigParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type KKSV1ClusterNodeGroupListParams struct {
	XRegion string `json:"-"`
}

type KKSV1ClusterNodeGroupGetParams struct {
	XRegion string `json:"-"`
}

type KKSV1ClusterNodeGroupDeleteParams struct {
	XRegion string `json:"-"`
}

type KKSScalingConfig struct {
	MinSize     int64 `json:"minSize"`
	MaxSize     int64 `json:"maxSize"`
	DesiredSize int64 `json:"desiredSize"`
}

type KKSTaint struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Effect string `json:"effect"`
}

type KKSRemoteAccess struct {
	SSHKeyKRN                string   `json:"sshKeyKrn,omitempty"`
	SourceSecurityGroupsKRNs []string `json:"sourceSecurityGroupsKrns,omitzero"`
}

type KKSNodeRepairConfig struct {
	Enabled bool `json:"enabled"`
}

type KKSNodeGroup struct {
	Name             string              `json:"name"`
	KRN              string              `json:"krn"`
	ClusterKRN       string              `json:"clusterKrn"`
	CreatedAt        float64             `json:"createdAt"`
	UpdatedAt        float64             `json:"updatedAt"`
	UserUpdatedAt    float64             `json:"userUpdatedAt"`
	InstanceTypes    string              `json:"instanceTypes"`
	RemoteAccess     KKSRemoteAccess     `json:"remoteAccess"`
	ScalingConfig    KKSScalingConfig    `json:"scalingConfig"`
	SubnetsKRN       string              `json:"subnetsKrn"`
	DiskSize         int64               `json:"diskSize"`
	Status           string              `json:"status"`
	Version          string              `json:"version"`
	Labels           map[string]string   `json:"labels"`
	Taints           []KKSTaint          `json:"taints"`
	NodeRepairConfig KKSNodeRepairConfig `json:"nodeRepairConfig"`
}

type KKSV1ClusterNodeGroupCreateResponse struct {
	Nodegroup KKSNodeGroup `json:"nodegroup"`
}

type KKSV1ClusterNodeGroupListResponse struct {
	Nodegroups []KKSNodeGroup `json:"nodegroups"`
}

type KKSV1ClusterNodeGroupGetResponse struct {
	Nodegroup KKSNodeGroup `json:"nodegroup"`
}

type KKSNodeGroupConfigUpdate struct {
	// Errors preserves error items without assuming a schema for their contents.
	Errors    []json.RawMessage `json:"errors"`
	Status    string            `json:"status"`
	CreatedAt float64           `json:"createdAt"`
	Type      string            `json:"type"`
}

type KKSV1ClusterNodeGroupUpdateConfigResponse struct {
	Update KKSNodeGroupConfigUpdate `json:"update"`
}
