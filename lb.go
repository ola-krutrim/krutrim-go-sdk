package krutrim

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/ola-krutrim/krutrim-go-sdk/internal/apijson"
	"github.com/ola-krutrim/krutrim-go-sdk/internal/apiquery"
	"github.com/ola-krutrim/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
	"github.com/ola-krutrim/krutrim-go-sdk/packages/param"
)

type HighlvlLoadBalancerService struct {
	Options []option.RequestOption
}

func NewHighlvlLoadBalancerService(opts ...option.RequestOption) (r HighlvlLoadBalancerService) {
	r.Options = opts
	return
}

type TargetGroupMember struct {
	Name         string `json:"name"`
	Address      string `json:"address"`
	ProtocolPort int64  `json:"protocol_port"`
	Weight       int64  `json:"weight"`
}

type HealthMonitor struct {
	Name       string `json:"name"`
	HType      string `json:"h_type"`
	Delay      int64  `json:"delay"`
	Timeout    int64  `json:"timeout"`
	MaxRetries int64  `json:"max_retries"`
	URLPath    string `json:"url_path,omitempty"`
}

type CreateTargetGroupParams struct {
	// Header
	XRegion string `json:"-"`

	// Body
	VpcID           string              `json:"vpc_id"`
	TargetGroupName string              `json:"target_group_name"`
	Members         []TargetGroupMember `json:"members"`
	HealthMonitor   HealthMonitor       `json:"health_monitor"`

	paramObj
}

type UpdateListenerData struct {
	ListenerID    string   `json:"listener_id"`
	ListenerIndex int64    `json:"listener_index"`
	ListenerName  string   `json:"listener_name"`
	Protocol      string   `json:"protocol"`
	ProtocolPort  int64    `json:"protocol_port"`
	SNIContainerRefs []string `json:"sni_container_refs,omitempty"`
}
type UpdateHealthMonitor struct {
	HealthMonitorID string `json:"health_monitor_id"`
	HType           string `json:"h_type"`
	Delay           int64  `json:"delay"`
	Timeout         int64  `json:"timeout"`
	MaxRetries      int64  `json:"max_retries"`
	URLPath         string `json:"url_path,omitempty"`
	Name            string `json:"name"`
}

type UpdateMember struct {
	KRN         string `json:"krn"`
	MemberID    string `json:"member_id"`
	MemberIndex int64  `json:"member_index"`
	Status      string `json:"status"`
}

type UpdatePoolData struct {
	PoolID          string               `json:"pool_id"`
	PoolName        string               `json:"pool_name"`
	Protocol        string               `json:"protocol"`
	LBAlgorithm     string               `json:"lb_algorithm"`
	AdminStateUp    bool                 `json:"admin_state_up"`
	TargetGroupName string               `json:"target_group_name"`
	HealthMonitor   *UpdateHealthMonitor `json:"healthmonitor_data,omitempty"`
	MemberData      []UpdateMember       `json:"member_data,omitempty"`
}

type UpdateRuleData struct {
	RuleID      string `json:"rule_id"`
	RuleKRN     string `json:"rule_krn"`
	CompareType string `json:"compare_type"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Key         string `json:"key"`
}
type UpdatePolicyData struct {
	PolicyID   string           `json:"policy_id"`
	PolicyName string           `json:"policy_name"`
	Action     string           `json:"action"`
	RuleData   []UpdateRuleData `json:"rule_data"`
}
type UpdateListener struct {
	ListenerData UpdateListenerData `json:"listener_data"`
	PoolData     []UpdatePoolData   `json:"pool_data"`
	PolicyData   []UpdatePolicyData `json:"policy_data,omitempty"`
}

func (r CreateTargetGroupParams) MarshalJSON() ([]byte, error) {
	type shadow CreateTargetGroupParams
	return param.MarshalObject(r, (*shadow)(&r))
}

func (r *CreateTargetGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UpdateTargetGroupParams struct {
	XRegion string `json:"-"`

	VpcID           string              `json:"vpc_id"`
	TargetGroupName string              `json:"target_group_name"`
	Members         []TargetGroupMember `json:"members,omitempty"`
	HealthMonitor   *HealthMonitor      `json:"health_monitor,omitempty"`

	paramObj
}

type TargetGroupStatusResponse struct {
	Status string `json:"status"`
}

func (r UpdateTargetGroupParams) MarshalJSON() ([]byte, error) {
	type shadow UpdateTargetGroupParams
	return param.MarshalObject(r, (*shadow)(&r))
}

func (r *HighlvlLoadBalancerService) CreateTargetGroup(
	ctx context.Context,
	body CreateTargetGroupParams,
	opts ...option.RequestOption,
) (res *TargetGroupStatusResponse, err error) {

	if body.XRegion != "" {
		opts = append(opts, option.WithHeader("x-region", body.XRegion))
	}

	opts = slices.Concat(r.Options, opts)
	path := "v1/highlvl/create_target_group"

	err = requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodPost,
		path,
		body,
		&res,
		opts...,
	)
	return
}

type GetTargetGroupListParams struct {
	XRegion string `json:"-"`

	VpcID string `query:"vpc_id,required" json:"-"`

	paramObj
}

func (r GetTargetGroupListParams) URLQuery() (url.Values, error) {
	return apiquery.Marshal(r)
}

func (r *HighlvlLoadBalancerService) GetTargetGroupList(
	ctx context.Context,
	query GetTargetGroupListParams,
	opts ...option.RequestOption,
) (res []map[string]any, err error) {

	if query.XRegion != "" {
		opts = append(opts, option.WithHeader("x-region", query.XRegion))
	}

	opts = slices.Concat(r.Options, opts)
	path := "v1/highlvl/get_tg_list"

	err = requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodGet,
		path,
		query,
		&res,
		opts...,
	)
	return
}

type GetTargetGroupParams struct {
	XRegion string `json:"-"`

	VpcID           string `query:"vpc_id,required" json:"-"`
	TargetGroupName string `query:"target_group_name,required" json:"-"`

	paramObj
}

func (r GetTargetGroupParams) URLQuery() (url.Values, error) {
	return apiquery.Marshal(r)
}

func (r *HighlvlLoadBalancerService) GetTargetGroup(
	ctx context.Context,
	query GetTargetGroupParams,
	opts ...option.RequestOption,
) (res map[string]any, err error) {

	if query.XRegion != "" {
		opts = append(opts, option.WithHeader("x-region", query.XRegion))
	}

	opts = slices.Concat(r.Options, opts)
	path := "v1/highlvl/get_target_groups"

	err = requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodGet,
		path,
		query,
		&res,
		opts...,
	)
	return
}

type GetTargetGroupNamesParams struct {
	XRegion string `json:"-"`

	VpcID string `query:"vpc_id,required" json:"-"`

	paramObj
}

func (r GetTargetGroupNamesParams) URLQuery() (url.Values, error) {
	return apiquery.Marshal(r)
}

func (r *HighlvlLoadBalancerService) GetTargetGroupNames(
	ctx context.Context,
	query GetTargetGroupNamesParams,
	opts ...option.RequestOption,
) (res []string, err error) {

	if query.XRegion != "" {
		opts = append(opts, option.WithHeader("x-region", query.XRegion))
	}

	opts = slices.Concat(r.Options, opts)
	path := "v1/highlvl/get_target_group_names"

	err = requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodGet,
		path,
		query,
		&res,
		opts...,
	)
	return
}

func (r *HighlvlLoadBalancerService) UpdateTargetGroup(
	ctx context.Context,
	body UpdateTargetGroupParams,
	opts ...option.RequestOption,
) (res map[string]any, err error) {

	if body.XRegion != "" {
		opts = append(opts, option.WithHeader("x-region", body.XRegion))
	}

	opts = slices.Concat(r.Options, opts)
	path := "v1/highlvl/updatetg"

	err = requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodPut,
		path,
		body,
		&res,
		opts...,
	)
	return
}

func (r *HighlvlLoadBalancerService) DeleteTargetGroup(
	ctx context.Context,
	vpcID string,
	targetGroupName string,
	xRegion string,
	opts ...option.RequestOption,
) (res map[string]any, err error) {

	if xRegion != "" {
		opts = append(opts, option.WithHeader("x-region", xRegion))
	}

	opts = slices.Concat(r.Options, opts)
	opts = append(opts,
		option.WithQuery("vpc_id", vpcID),
		option.WithQuery("target_group_name", targetGroupName),
	)

	path := "v1/highlvl/target_group"

	err = requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodDelete,
		path,
		nil,
		&res,
		opts...,
	)
	return
}

type LoadBalancerData struct {
	LBName      string `json:"lb_name"`
	Description string `json:"description,omitempty"`
	CreatePort  bool   `json:"create_port"`
	FloatingIP  bool   `json:"floating_ip"`
	VpcID       string `json:"vpc_id"`
	NetworkID   string `json:"network_id"`
	VipSubnetID string `json:"vip_subnet_id"`
	LBType      string `json:"lb_type"`
	Flavor      string `json:"flavor"`
}

type ListenerData struct {
	Name                   string   `json:"name,omitempty"`
	Protocol               string   `json:"protocol"`
	ProtocolPort           int64    `json:"protocol_port"`
	ListenerName           string   `json:"listener_name"`
	DefaultPool            bool     `json:"default_pool"`
	DefaultTLSContainerRef string   `json:"default_tls_container_ref,omitempty"`
	SNIContainerRefs       []string `json:"sni_container_refs,omitempty"`
}

type PoolData struct {
	PoolName        string `json:"pool_name"`
	Protocol        string `json:"protocol"`
	LBAlgorithm     string `json:"lb_algorithm"`
	AdminStateUp    bool   `json:"admin_state_up"`
	TargetGroupName string `json:"target_group_name"`
	F5              bool   `json:"F5,omitempty"`
}

type RuleData struct {
	CompareType string `json:"compare_type"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Key         string `json:"key,omitempty"`
}

type PolicyData struct {
	PolicyName       string     `json:"policy_name"`
	Action           string     `json:"action"`
	RedirectURL      string     `json:"redirect_url,omitempty"`
	RedirectPoolName string     `json:"redirect_pool_name,omitempty"`
	RuleData         []RuleData `json:"rule_data"`
}

type Listener struct {
	ListenerData ListenerData `json:"listener_data"`
	PoolData     []PoolData   `json:"pool_data"`
	PolicyData   []PolicyData `json:"policy_data,omitempty"`
}

type CreateLoadBalancerParams struct {
	// Header
	XRegion string `json:"-"`

	// Body
	LoadBalancerData LoadBalancerData `json:"loadbalancer_data"`
	SecurityGroups   []string         `json:"security_groups"`
	ListenerCount    int              `json:"listener_count"`
	Listeners        []Listener       `json:"listeners"`

	paramObj
}


func (r CreateLoadBalancerParams) MarshalJSON() ([]byte, error) {
	type shadow CreateLoadBalancerParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type UpdateLoadBalancerParams struct {
	// Header
	XRegion string `json:"-"`

	// Body
	LBKrn            string                  `json:"lb_krn"`
	SecurityGroups   []string                `json:"security_groups"`
	LoadBalancerData *UpdateLoadBalancerData `json:"loadbalancer_data,omitempty"`
	Listeners        []UpdateListener        `json:"listeners,omitempty"`

	paramObj
}

type UpdateLoadBalancerData struct {
	LBName      string `json:"lb_name,omitempty"`
	Description string `json:"description,omitempty"`
	CreatePort  *bool  `json:"create_port,omitempty"`

	VpcID       string `json:"vpc_id,omitempty"`
	NetworkID   string `json:"network_id,omitempty"`
	VipSubnetID string `json:"vip_subnet_id,omitempty"`
}

func (r UpdateLoadBalancerParams) MarshalJSON() ([]byte, error) {
	type shadow UpdateLoadBalancerParams
	return param.MarshalObject(r, (*shadow)(&r))
}

func (r *HighlvlLoadBalancerService) CreateLoadBalancer(
	ctx context.Context,
	body CreateLoadBalancerParams,
	opts ...option.RequestOption,
) (res map[string]any, err error) {

	if body.XRegion != "" {
		opts = append(opts, option.WithHeader("x-region", body.XRegion))
	}

	opts = slices.Concat(r.Options, opts)
	path := "v3/highlvl/create_load_balancer_orchestration"

	err = requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodPost,
		path,
		body,
		&res,
		opts...,
	)
	return
}

func (r *HighlvlLoadBalancerService) GetTaskStatus(
	ctx context.Context,
	taskID string,
	xRegion string,
	opts ...option.RequestOption,
) (res map[string]any, err error) {

	if xRegion != "" {
		opts = append(opts, option.WithHeader("x-region", xRegion))
	}

	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v3/highlvl/task_status/%s", taskID)

	err = requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodGet,
		path,
		nil,
		&res,
		opts...,
	)
	return
}

// 9. Get Load Balancer List
func (r *HighlvlLoadBalancerService) GetLoadBalancerList(
	ctx context.Context,
	vpcID string,
	page int64,
	limit int64,
	xRegion string,
	opts ...option.RequestOption,
) (res map[string]any, err error) {

	if xRegion != "" {
		opts = append(opts, option.WithHeader("x-region", xRegion))
	}

	opts = slices.Concat(r.Options, opts)
	opts = append(opts,
		option.WithQuery("page", fmt.Sprintf("%d", page)),
		option.WithQuery("limit", fmt.Sprintf("%d", limit)),
	)

	path := fmt.Sprintf("v3/highlvl/get_lb_list_new/%s", vpcID)

	err = requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodGet,
		path,
		nil,
		&res,
		opts...,
	)
	return
}

// 10. Get Load Balancer Details
func (r *HighlvlLoadBalancerService) GetLoadBalancerDetails(
	ctx context.Context,
	lbKrn string,
	xRegion string,
	opts ...option.RequestOption,
) (res map[string]any, err error) {

	if xRegion != "" {
		opts = append(opts, option.WithHeader("x-region", xRegion))
	}

	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v3/highlvl/lb_details_new/%s", lbKrn)

	err = requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodGet,
		path,
		nil,
		&res,
		opts...,
	)
	return
}

// 11. Update Load Balancer
func (r *HighlvlLoadBalancerService) UpdateLoadBalancer(
	ctx context.Context,
	lbKrn string,
	body UpdateLoadBalancerParams,
	opts ...option.RequestOption,
) (res map[string]any, err error) {

	if body.XRegion != "" {
		opts = append(opts, option.WithHeader("x-region", body.XRegion))
	}

	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v3/highlvl/update_load_balancer/%s", lbKrn)

	err = requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodPut,
		path,
		body,
		&res,
		opts...,
	)
	return
}

// 12. Delete Load Balancer
func (r *HighlvlLoadBalancerService) DeleteLoadBalancer(
	ctx context.Context,
	lbKrn string,
	xRegion string,
	opts ...option.RequestOption,
) (res map[string]any, err error) {

	if xRegion != "" {
		opts = append(opts, option.WithHeader("x-region", xRegion))
	}

	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v3/highlvl/loadbalancer/%s", lbKrn)

	err = requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodDelete,
		path,
		nil,
		&res,
		opts...,
	)
	return
}
