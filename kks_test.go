package krutrim_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	krutrim "github.com/ola-krutrim/krutrim-go-sdk"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
)

const kksTestClusterKRN = "krn:kks:In-Bangalore-1:account:user:cluster:example"

type kksTestHTTPClient func(*http.Request) (*http.Response, error)

func (f kksTestHTTPClient) Do(req *http.Request) (*http.Response, error) { return f(req) }

func newKKSTestClient(t *testing.T, method, path, region, body, response string) krutrim.Client {
	t.Helper()
	return krutrim.NewClient(
		option.WithBaseURL("https://kks.example.test"),
		option.WithAPIKey("test-token"),
		option.WithMaxRetries(0),
		option.WithHTTPClient(kksTestHTTPClient(func(req *http.Request) (*http.Response, error) {
			t.Helper()
			if req.Method != method || req.URL.EscapedPath() != path {
				t.Errorf("request = %s %s, want %s %s", req.Method, req.URL.EscapedPath(), method, path)
			}
			if req.Header.Get("Authorization") != "Bearer test-token" {
				t.Errorf("client authentication options were not inherited")
			}
			if got := req.Header.Get("x-region"); got != region {
				t.Errorf("x-region = %q, want %q", got, region)
			}
			if body == "" {
				if req.Body != nil {
					t.Errorf("%s request must not contain a body", method)
				}
			} else {
				if req.Header.Get("Content-Type") != "application/json" {
					t.Error("missing JSON content type")
				}
				if req.Body == nil {
					t.Fatal("missing request body")
				}
				data, err := io.ReadAll(req.Body)
				if err != nil {
					t.Fatal(err)
				}
				assertKKSJSON(t, data, body)
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{"Content-Type": {"application/json"}},
				Body:       io.NopCloser(strings.NewReader(response)),
				Request:    req,
			}, nil
		})),
	)
}

func assertKKSJSON(t *testing.T, data []byte, want string) {
	t.Helper()
	var gotValue, wantValue any
	if err := json.Unmarshal(data, &gotValue); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(want), &wantValue); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(gotValue, wantValue) {
		t.Errorf("JSON = %s, want %s", data, want)
	}
}

func TestKKSClusterCreate(t *testing.T) {
	client := newKKSTestClient(t, http.MethodPost, "/v1/kks/clusters", "In-Bangalore-1", `{
		"name":"test", "version":"1.36",
		"resourcesVpcConfig":{"vpcKrn":"vpc-krn","subnetKrns":"subnet-krn","endpointPrivateAccess":false},
		"kubernetesNetworkConfig":{"podIpv4Cidr":"10.244.0.0/16","serviceIpv4Cidr":"10.96.0.0/12"}
	}`, `{"cluster":{
		"name":"test","krn":"cluster-krn","createdAt":1792574430.7298555,
		"version":"1.36","versionSupported":false,"status":"CREATING",
		"resourcesVpcConfig":{"vpcKrn":"vpc-krn","subnetKrns":"subnet-krn","endpointPrivateAccess":false},
		"kubernetesNetworkConfig":{"podIpv4Cidr":"10.244.0.0/16","serviceIpv4Cidr":"10.96.0.0/12"}
	}}`)
	res, err := client.KKS.V1.Cluster.Create(context.Background(), krutrim.KKSV1ClusterCreateParams{
		Name: "test", Version: "1.36", XRegion: "In-Bangalore-1",
		ResourcesVPCConfig: krutrim.KKSResourcesVPCConfig{VPCKRN: "vpc-krn", SubnetKRNs: "subnet-krn"},
		KubernetesNetworkConfig: krutrim.KKSKubernetesNetworkConfig{
			PodIPv4CIDR: "10.244.0.0/16", ServiceIPv4CIDR: "10.96.0.0/12",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if res == nil || res.Cluster.Name != "test" || res.Cluster.KRN != "cluster-krn" ||
		res.Cluster.CreatedAt != 1792574430.7298555 || res.Cluster.VersionSupported ||
		res.Cluster.Status != "CREATING" || res.Cluster.ResourcesVPCConfig.SubnetKRNs != "subnet-krn" ||
		res.Cluster.KubernetesNetworkConfig.PodIPv4CIDR != "10.244.0.0/16" {
		t.Fatalf("unexpected cluster response: %+v", res)
	}
}

func TestKKSClusterList(t *testing.T) {
	for _, region := range []string{"", "In-Bangalore-1", "region /?&"} {
		t.Run("region="+region, func(t *testing.T) {
			client := newKKSTestClient(t, http.MethodGet, "/v1/kks/clusters", "header-region", "", `{
				"clusters":[{"name":"test","krn":"cluster-krn","status":"PROVISIONED",
				"kubernetesVersion":"1.36","versionSupported":true,"nodegroupCount":1,"createdAt":1792574430}]
			}`)
			params := krutrim.KKSV1ClusterListParams{XRegion: "header-region"}
			if region != "" {
				params.Region = krutrim.String(region)
			}
			var raw *http.Response
			res, err := client.KKS.V1.Cluster.List(context.Background(), params, option.WithResponseInto(&raw))
			if err != nil {
				t.Fatal(err)
			}
			wantQuery := url.Values{}
			if region != "" {
				wantQuery.Set("region", region)
			}
			if raw.Request.URL.RawQuery != wantQuery.Encode() {
				t.Errorf("query = %q, want %q", raw.Request.URL.RawQuery, wantQuery.Encode())
			}
			if res == nil || len(res.Clusters) != 1 || res.Clusters[0].KubernetesVersion != "1.36" ||
				!res.Clusters[0].VersionSupported || res.Clusters[0].NodegroupCount != 1 || res.Clusters[0].CreatedAt != 1792574430 {
				t.Fatalf("unexpected list response: %+v", res)
			}
		})
	}
}

func TestKKSClusterGet(t *testing.T) {
	// Reserved path characters must stay inside a single path segment.
	clusterKRN := kksTestClusterKRN + "/part?query=value#fragment"
	client := newKKSTestClient(t, http.MethodGet, "/v1/kks/clusters/"+url.PathEscape(clusterKRN), "In-Bangalore-1", "",
		`{"cluster":{"name":"test","version":"1.36","versionSupported":true,"status":"PROVISIONED","createdAt":1792574430.729855}}`)
	res, err := client.KKS.V1.Cluster.Get(context.Background(), clusterKRN, krutrim.KKSV1ClusterGetParams{XRegion: "In-Bangalore-1"})
	if err != nil {
		t.Fatal(err)
	}
	if res == nil || res.Cluster.Status != "PROVISIONED" || !res.Cluster.VersionSupported || res.Cluster.CreatedAt != 1792574430.729855 {
		t.Fatalf("unexpected get response: %+v", res)
	}
}

func TestKKSNodeGroupCreate(t *testing.T) {
	for _, full := range []bool{false, true} {
		name := "minimal"
		if full {
			name = "all optional fields"
		}
		t.Run(name, func(t *testing.T) {
			body := krutrim.KKSV1ClusterNodeGroupCreateParams{
				Name: "default-pool", InstanceTypes: "CPU-4x-16GB", DiskSize: 166, XRegion: "In-Bangalore-1",
				ScalingConfig: krutrim.KKSScalingConfig{MinSize: 1, MaxSize: 5, DesiredSize: 2}, SubnetsKRN: "subnet-krn",
			}
			want := `{"name":"default-pool","instanceTypes":"CPU-4x-16GB","diskSize":166,
				"scalingConfig":{"minSize":1,"maxSize":5,"desiredSize":2},"subnetsKrn":"subnet-krn"}`
			groups := "null"
			if full {
				body.Labels = map[string]string{}
				body.Taints = []krutrim.KKSTaint{}
				body.RemoteAccess = &krutrim.KKSRemoteAccess{SourceSecurityGroupsKRNs: []string{}}
				body.NodeRepairConfig = &krutrim.KKSNodeRepairConfig{Enabled: true}
				want = `{"name":"default-pool","instanceTypes":"CPU-4x-16GB","diskSize":166,
					"scalingConfig":{"minSize":1,"maxSize":5,"desiredSize":2},"subnetsKrn":"subnet-krn",
					"labels":{},"taints":[],"remoteAccess":{"sourceSecurityGroupsKrns":[]},"nodeRepairConfig":{"enabled":true}}`
				groups = "[]"
			}
			client := newKKSTestClient(t, http.MethodPost, "/v1/kks/clusters/"+kksTestClusterKRN+"/node-groups", "In-Bangalore-1", want,
				`{"nodegroup":{"name":"default-pool","krn":"nodegroup-krn","clusterKrn":"cluster-krn",
				"createdAt":1792574509.8295214,"updatedAt":1792574509.8295214,"userUpdatedAt":1792574509.8288865,
				"instanceTypes":"CPU-4x-16GB","diskSize":166,"subnetsKrn":"subnet-krn","version":"1.36","status":"CREATING",
				"scalingConfig":{"minSize":1,"maxSize":5,"desiredSize":2},"remoteAccess":{"sshKeyKrn":"","sourceSecurityGroupsKrns":`+groups+`}}}`)
			res, err := client.KKS.V1.Cluster.NodeGroup.Create(context.Background(), kksTestClusterKRN, body)
			if err != nil {
				t.Fatal(err)
			}
			if res == nil || res.Nodegroup.KRN != "nodegroup-krn" || res.Nodegroup.CreatedAt != 1792574509.8295214 ||
				res.Nodegroup.UserUpdatedAt != 1792574509.8288865 || res.Nodegroup.ScalingConfig.MaxSize != 5 {
				t.Fatalf("unexpected node group response: %+v", res)
			}
			if (res.Nodegroup.RemoteAccess.SourceSecurityGroupsKRNs != nil) != full {
				t.Error("null and empty security group responses were not preserved")
			}
		})
	}
}

func TestKKSNodeGroupExplicitZeroValues(t *testing.T) {
	body := krutrim.KKSV1ClusterNodeGroupCreateParams{
		Name: "zero-pool", InstanceTypes: "CPU-4x-16GB", DiskSize: 100, SubnetsKRN: "subnet-krn",
		ScalingConfig:    krutrim.KKSScalingConfig{MaxSize: 4},
		Labels:           map[string]string{"role": "worker"},
		Taints:           []krutrim.KKSTaint{{Key: "dedicated", Value: "batch", Effect: "NoSchedule"}},
		RemoteAccess:     &krutrim.KKSRemoteAccess{SSHKeyKRN: "ssh-key", SourceSecurityGroupsKRNs: []string{"security-group"}},
		NodeRepairConfig: &krutrim.KKSNodeRepairConfig{Enabled: false},
	}
	data, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	assertKKSJSON(t, data, `{"name":"zero-pool","instanceTypes":"CPU-4x-16GB","diskSize":100,"subnetsKrn":"subnet-krn",
		"scalingConfig":{"minSize":0,"maxSize":4,"desiredSize":0},"labels":{"role":"worker"},
		"taints":[{"key":"dedicated","value":"batch","effect":"NoSchedule"}],
		"remoteAccess":{"sshKeyKrn":"ssh-key","sourceSecurityGroupsKrns":["security-group"]},"nodeRepairConfig":{"enabled":false}}`)
}

func TestKKSNodeGroupList(t *testing.T) {
	client := newKKSTestClient(t, http.MethodGet, "/v1/kks/clusters/"+kksTestClusterKRN+"/node-groups", "", "", `{
		"nodegroups":[
		{"krn":"first","name":"default-pool","status":"RUNNING","instanceTypes":"CPU-4x-16GB","diskSize":166,"version":"1.36","scalingConfig":{"minSize":1,"maxSize":5,"desiredSize":2}},
		{"krn":"second","name":"failed-pool","status":"FAILED","instanceTypes":"CPU-4x-16GB","diskSize":100,"version":"1.36","scalingConfig":{"minSize":1,"maxSize":4,"desiredSize":2}}]
	}`)
	res, err := client.KKS.V1.Cluster.NodeGroup.List(context.Background(), kksTestClusterKRN, krutrim.KKSV1ClusterNodeGroupListParams{})
	if err != nil {
		t.Fatal(err)
	}
	if res == nil || len(res.Nodegroups) != 2 || res.Nodegroups[0].Status != "RUNNING" ||
		res.Nodegroups[1].Status != "FAILED" || res.Nodegroups[1].DiskSize != 100 || res.Nodegroups[1].ScalingConfig.MaxSize != 4 {
		t.Fatalf("unexpected node group list response: %+v", res)
	}
}

func TestKKSAddonCreate(t *testing.T) {
	for _, version := range []string{"", "1.0.0"} {
		t.Run("version="+version, func(t *testing.T) {
			body := krutrim.KKSV1ClusterAddonCreateParams{AddonName: "kube-proxy", XRegion: "In-Bangalore-1"}
			want := `{"addonName":"kube-proxy"}`
			if version != "" {
				body.AddonName = "csi"
				body.AddonVersion = krutrim.String(version)
				want = `{"addonName":"csi","addonVersion":"1.0.0"}`
			}
			client := newKKSTestClient(t, http.MethodPost, "/v1/kks/clusters/"+kksTestClusterKRN+"/addons", "In-Bangalore-1", want,
				`{"addon":{"addonKrn":"addon-krn","addonName":"`+body.AddonName+`","addonVersion":"1.33.0","clusterName":"test","createdAt":"2026-10-21T09:28:24Z","status":"CREATING"}}`)
			res, err := client.KKS.V1.Cluster.Addon.Create(context.Background(), kksTestClusterKRN, body)
			if err != nil {
				t.Fatal(err)
			}
			if res == nil || res.Addon.AddonKRN != "addon-krn" || res.Addon.AddonName != body.AddonName ||
				res.Addon.Status != "CREATING" || !res.Addon.CreatedAt.Equal(time.Date(2026, 10, 21, 9, 28, 24, 0, time.UTC)) {
				t.Fatalf("unexpected add-on response: %+v", res)
			}
		})
	}
}

func TestKKSKubeconfig(t *testing.T) {
	for _, contentType := range []string{"application/yaml", "application/octet-stream", "text/plain", "application/json"} {
		t.Run(contentType, func(t *testing.T) {
			want := "apiVersion: v1\nkind: Config\nclusters: []\n"
			if contentType == "application/json" {
				want = `{"apiVersion":"v1","kind":"Config","clusters":[]}`
			}
			client := krutrim.NewClient(
				option.WithBaseURL("https://kks.example.test"), option.WithAPIKey("test-token"),
				option.WithHTTPClient(kksTestHTTPClient(func(req *http.Request) (*http.Response, error) {
					if req.Method != http.MethodGet || req.URL.Path != "/v1/kks/clusters/"+kksTestClusterKRN+"/kubeconfig" ||
						req.Header.Get("Accept") != "*/*" || req.Header.Get("Authorization") != "Bearer test-token" || req.Body != nil {
						t.Errorf("unexpected kubeconfig request: %+v", req)
					}
					return &http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": {contentType}}, Body: io.NopCloser(strings.NewReader(want))}, nil
				})),
			)
			res, err := client.KKS.V1.Cluster.GetKubeconfig(context.Background(), kksTestClusterKRN, krutrim.KKSV1ClusterGetKubeconfigParams{})
			if err != nil {
				t.Fatal(err)
			}
			if string(res) != want {
				t.Errorf("kubeconfig = %q, want %q", res, want)
			}
		})
	}
}

func TestKKSEmptyLists(t *testing.T) {
	client := newKKSTestClient(t, http.MethodGet, "/v1/kks/clusters", "", "", `{"clusters":[]}`)
	clusters, err := client.KKS.V1.Cluster.List(context.Background(), krutrim.KKSV1ClusterListParams{})
	if err != nil || clusters == nil || len(clusters.Clusters) != 0 {
		t.Fatalf("empty clusters = %+v, err = %v", clusters, err)
	}
	client = newKKSTestClient(t, http.MethodGet, "/v1/kks/clusters/"+kksTestClusterKRN+"/node-groups", "", "", `{"nodegroups":[]}`)
	nodes, err := client.KKS.V1.Cluster.NodeGroup.List(context.Background(), kksTestClusterKRN, krutrim.KKSV1ClusterNodeGroupListParams{})
	if err != nil || nodes == nil || len(nodes.Nodegroups) != 0 {
		t.Fatalf("empty node groups = %+v, err = %v", nodes, err)
	}
}

func TestKKSMissingClusterKRN(t *testing.T) {
	client := krutrim.NewClient(option.WithHTTPClient(kksTestHTTPClient(func(req *http.Request) (*http.Response, error) {
		t.Fatal("missing KRN must fail before sending a request")
		return nil, nil
	})))
	calls := map[string]func() error{
		"get": func() error {
			_, err := client.KKS.V1.Cluster.Get(context.Background(), "", krutrim.KKSV1ClusterGetParams{})
			return err
		},
		"delete cluster": func() error {
			return client.KKS.V1.Cluster.Delete(context.Background(), "", krutrim.KKSV1ClusterDeleteParams{})
		},
		"kubeconfig": func() error {
			_, err := client.KKS.V1.Cluster.GetKubeconfig(context.Background(), "", krutrim.KKSV1ClusterGetKubeconfigParams{})
			return err
		},
		"create node group": func() error {
			_, err := client.KKS.V1.Cluster.NodeGroup.Create(context.Background(), "", krutrim.KKSV1ClusterNodeGroupCreateParams{})
			return err
		},
		"list node groups": func() error {
			_, err := client.KKS.V1.Cluster.NodeGroup.List(context.Background(), "", krutrim.KKSV1ClusterNodeGroupListParams{})
			return err
		},
		"get node group": func() error {
			_, err := client.KKS.V1.Cluster.NodeGroup.Get(context.Background(), "", "nodegroup", krutrim.KKSV1ClusterNodeGroupGetParams{})
			return err
		},
		"update node group config": func() error {
			_, err := client.KKS.V1.Cluster.NodeGroup.UpdateConfig(context.Background(), "", "nodegroup", krutrim.KKSV1ClusterNodeGroupUpdateConfigParams{})
			return err
		},
		"delete node group": func() error {
			return client.KKS.V1.Cluster.NodeGroup.Delete(context.Background(), "", "nodegroup", krutrim.KKSV1ClusterNodeGroupDeleteParams{})
		},
		"create addon": func() error {
			_, err := client.KKS.V1.Cluster.Addon.Create(context.Background(), "", krutrim.KKSV1ClusterAddonCreateParams{})
			return err
		},
		"list addons": func() error {
			_, err := client.KKS.V1.Cluster.Addon.List(context.Background(), "", krutrim.KKSV1ClusterAddonListParams{})
			return err
		},
		"delete addon": func() error {
			return client.KKS.V1.Cluster.Addon.Delete(context.Background(), "", "addon", krutrim.KKSV1ClusterAddonDeleteParams{})
		},
	}
	for name, call := range calls {
		t.Run(name, func(t *testing.T) {
			if err := call(); err == nil || !strings.Contains(err.Error(), "clusterKRN") {
				t.Fatalf("expected missing cluster KRN error, got %v", err)
			}
		})
	}
}

func TestKKSRequestOptions(t *testing.T) {
	client := newKKSTestClient(t, http.MethodPost, "/v1/kks/clusters/"+kksTestClusterKRN+"/addons", "override-region", `{"addonName":"csi"}`, `{"addon":{}}`)
	_, err := client.KKS.V1.Cluster.Addon.Create(context.Background(), kksTestClusterKRN,
		krutrim.KKSV1ClusterAddonCreateParams{AddonName: "csi", XRegion: "parameter-region"},
		option.WithHeader("x-region", "override-region"))
	if err != nil {
		t.Fatal(err)
	}
	// A region on the parent client is used when no parameter overrides it.
	client = newKKSTestClient(t, http.MethodGet, "/v1/kks/clusters", "client-region", "", `{"clusters":[]}`)
	client.KKS = krutrim.NewKKSService(append(client.Options, option.WithHeader("x-region", "client-region"))...)
	_, err = client.KKS.V1.Cluster.List(context.Background(), krutrim.KKSV1ClusterListParams{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestKKSErrorPropagation(t *testing.T) {
	client := krutrim.NewClient(
		option.WithBaseURL("https://kks.example.test"), option.WithMaxRetries(0),
		option.WithHTTPClient(kksTestHTTPClient(func(req *http.Request) (*http.Response, error) {
			if err := req.Context().Err(); err != nil {
				return nil, err
			}
			return &http.Response{StatusCode: http.StatusForbidden, Body: io.NopCloser(strings.NewReader(`{"message":"denied"}`))}, nil
		})),
	)
	_, err := client.KKS.V1.Cluster.Get(context.Background(), kksTestClusterKRN, krutrim.KKSV1ClusterGetParams{})
	var apiErr *krutrim.Error
	if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusForbidden {
		t.Fatalf("expected SDK API error with status 403, got %v", err)
	}
	_, err = client.KKS.V1.Cluster.NodeGroup.UpdateConfig(context.Background(), kksTestClusterKRN, "nodegroup-krn", krutrim.KKSV1ClusterNodeGroupUpdateConfigParams{
		ScalingConfig: krutrim.KKSScalingConfig{MinSize: 2, MaxSize: 5, DesiredSize: 2},
	})
	if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusForbidden {
		t.Fatalf("expected update-config API error with status 403, got %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = client.KKS.V1.Cluster.List(ctx, krutrim.KKSV1ClusterListParams{})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context cancellation, got %v", err)
	}
}

func TestKKSNodeGroupGet(t *testing.T) {
	clusterKRN := kksTestClusterKRN + "/cluster?part=1#fragment"
	nodeGroupKRN := "krn:kks:In-Bangalore-1:account:user:nodegroup:example/part?tag=1#node"
	path := "/v1/kks/clusters/" + url.PathEscape(clusterKRN) + "/node-groups/" + url.PathEscape(nodeGroupKRN)
	client := newKKSTestClient(t, http.MethodGet, path, "In-Bangalore-1", "", `{"nodegroup":{
		"name":"default-pool","krn":"nodegroup-krn","clusterKrn":"cluster-krn",
		"createdAt":1792574509.829521,"updatedAt":1792574511.36277,"userUpdatedAt":1792574509.828886,
		"instanceTypes":"CPU-4x-16GB","remoteAccess":{"sshKeyKrn":"","sourceSecurityGroupsKrns":[]},
		"scalingConfig":{"desiredSize":2,"maxSize":5,"minSize":1},"subnetsKrn":"subnet-krn",
		"diskSize":166,"status":"RUNNING","version":"1.36"}}`)
	res, err := client.KKS.V1.Cluster.NodeGroup.Get(context.Background(), clusterKRN, nodeGroupKRN, krutrim.KKSV1ClusterNodeGroupGetParams{XRegion: "In-Bangalore-1"})
	if err != nil {
		t.Fatal(err)
	}
	if res == nil || res.Nodegroup.KRN != "nodegroup-krn" || res.Nodegroup.ClusterKRN != "cluster-krn" ||
		res.Nodegroup.CreatedAt != 1792574509.829521 || res.Nodegroup.UpdatedAt != 1792574511.36277 ||
		res.Nodegroup.UserUpdatedAt != 1792574509.828886 || res.Nodegroup.Status != "RUNNING" ||
		res.Nodegroup.DiskSize != 166 || res.Nodegroup.ScalingConfig.DesiredSize != 2 || res.Nodegroup.RemoteAccess.SourceSecurityGroupsKRNs == nil {
		t.Fatalf("unexpected node group: %+v", res)
	}
}

func TestKKSNodeGroupUpdateConfig(t *testing.T) {
	clusterKRN := kksTestClusterKRN + "/cluster?part=1#fragment"
	nodeGroupKRN := "krn:kks:In-Bangalore-1:account:user:nodegroup:example/part?tag=1#node"
	path := "/v1/kks/clusters/" + url.PathEscape(clusterKRN) + "/node-groups/" + url.PathEscape(nodeGroupKRN) + "/update-config"
	for _, test := range []struct {
		name    string
		scaling krutrim.KKSScalingConfig
		body    string
		errors  string
		status  string
	}{
		{"sample", krutrim.KKSScalingConfig{MinSize: 2, MaxSize: 5, DesiredSize: 2},
			`{"scalingConfig":{"minSize":2,"maxSize":5,"desiredSize":2}}`, `[]`, "Success"},
		{"explicit zero sizes", krutrim.KKSScalingConfig{MinSize: 0, MaxSize: 5, DesiredSize: 0},
			`{"scalingConfig":{"minSize":0,"maxSize":5,"desiredSize":0}}`, `[]`, "Success"},
		{"unspecified error item schema", krutrim.KKSScalingConfig{MinSize: 2, MaxSize: 5, DesiredSize: 2},
			`{"scalingConfig":{"minSize":2,"maxSize":5,"desiredSize":2}}`, `[{"code":"example","message":"details"},"additional details"]`, "Failed"},
	} {
		t.Run(test.name, func(t *testing.T) {
			response := `{"update":{"errors":` + test.errors + `,"status":"` + test.status + `","createdAt":1792579469.0122135,"type":"ConfigUpdate"}}`
			client := newKKSTestClient(t, http.MethodPost, path, "In-Bangalore-1", test.body, response)
			res, err := client.KKS.V1.Cluster.NodeGroup.UpdateConfig(context.Background(), clusterKRN, nodeGroupKRN, krutrim.KKSV1ClusterNodeGroupUpdateConfigParams{
				ScalingConfig: test.scaling,
				XRegion:       "parameter-region",
			}, option.WithHeader("x-region", "In-Bangalore-1"))
			if err != nil {
				t.Fatal(err)
			}
			if res == nil || res.Update.Status != test.status || res.Update.Type != "ConfigUpdate" ||
				res.Update.CreatedAt != 1792579469.0122135 || res.Update.Errors == nil {
				t.Fatalf("unexpected config update: %+v", res)
			}
			data, err := json.Marshal(res)
			if err != nil {
				t.Fatal(err)
			}
			assertKKSJSON(t, data, response)
		})
	}
}

func TestKKSFlavorList(t *testing.T) {
	want := `[
		{"name":"CPU-4x-16GB","cpu":4,"memory":"16Gi","gpu_count":0,"gpu_memory":0},
		{"name":"CPU-8x-32GB","cpu":8,"memory":"32Gi","gpu_count":0,"gpu_memory":0},
		{"name":"CPU-16x-64GB","cpu":16,"memory":"64Gi","gpu_count":0,"gpu_memory":0},
		{"name":"CPU-32x-128GB","cpu":32,"memory":"128Gi","gpu_count":0,"gpu_memory":0},
		{"name":"H100-NVLINK-1x","cpu":24,"memory":"200Gi","gpu_count":1,"gpu_memory":80}
	]`
	client := newKKSTestClient(t, http.MethodGet, "/v1/kks/flavors", "In-Bangalore-1", "", want)
	res, err := client.KKS.V1.Flavor.List(context.Background(), krutrim.KKSV1FlavorListParams{XRegion: "In-Bangalore-1"})
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 5 || res[0].CPU != 4 || res[0].GPUCount != 0 || res[4].GPUCount != 1 || res[4].GPUMemory != 80 || res[4].Memory != "200Gi" {
		t.Fatalf("unexpected flavors: %+v", res)
	}
	data, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	assertKKSJSON(t, data, want)
}

func TestKKSAddonCatalogList(t *testing.T) {
	// Preserve the complete supplied catalog response, including both defaults for kube-proxy.
	fixture, err := os.ReadFile("testdata/kks/addons.json")
	if err != nil {
		t.Fatal(err)
	}
	client := newKKSTestClient(t, http.MethodGet, "/v1/kks/addons", "In-Bangalore-1", "", string(fixture))
	res, err := client.KKS.V1.Addon.List(context.Background(), krutrim.KKSV1AddonListParams{XRegion: "In-Bangalore-1"})
	if err != nil {
		t.Fatal(err)
	}
	if res == nil || len(res.Addons) != 7 || res.Addons[0].AddonName != "csi" || !res.Addons[0].IsDeletable ||
		len(res.Addons[5].AddonVersions) != 2 || res.Addons[5].AddonVersions[0].Compatibilities[1].DefaultVersion ||
		!res.Addons[5].AddonVersions[1].Compatibilities[1].DefaultVersion {
		t.Fatalf("unexpected add-on catalog: %+v", res)
	}
	data, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	assertKKSJSON(t, data, string(fixture))
}

func TestKKSClusterAddonList(t *testing.T) {
	clusterKRN := kksTestClusterKRN + "/part?tag=1#cluster"
	want := `{"addons":[
		{"addonKrn":"dns-krn","addonName":"coredns","addonVersion":"1.12.4","status":"ACTIVE","type":"dns","createdAt":1792574510,"modifiedAt":1792574510},
		{"addonKrn":"csi-krn","addonName":"csi","addonVersion":"1.0.0","status":"ACTIVE","type":"storage","createdAt":1792574510,"modifiedAt":1792574510},
		{"addonKrn":"proxy-krn","addonName":"kube-proxy","addonVersion":"1.33.0","status":"ACTIVE","type":"networking","createdAt":1792574904,"modifiedAt":1792574904},
		{"addonKrn":"cilium-krn","addonName":"cilium","addonVersion":"1.18.2","status":"ACTIVE","type":"networking","createdAt":1792574510,"modifiedAt":1792574510},
		{"addonKrn":"metrics-krn","addonName":"metrics-server","addonVersion":"3.13.0","status":"ACTIVE","type":"monitoring","createdAt":1792574510,"modifiedAt":1792574510},
		{"addonKrn":"autoscaler-krn","addonName":"cluster-autoscaler","addonVersion":"9.50.1","status":"ACTIVE","type":"autoscaling","createdAt":1792574510,"modifiedAt":1792574510}
	]}`
	client := newKKSTestClient(t, http.MethodGet, "/v1/kks/clusters/"+url.PathEscape(clusterKRN)+"/addons", "In-Bangalore-1", "", want)
	res, err := client.KKS.V1.Cluster.Addon.List(context.Background(), clusterKRN, krutrim.KKSV1ClusterAddonListParams{XRegion: "In-Bangalore-1"})
	if err != nil {
		t.Fatal(err)
	}
	if res == nil || len(res.Addons) != 6 || res.Addons[0].Type != "dns" || res.Addons[0].CreatedAt != 1792574510 || res.Addons[2].ModifiedAt != 1792574904 {
		t.Fatalf("unexpected installed add-ons: %+v", res)
	}
	data, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	assertKKSJSON(t, data, want)
}

func TestKKSEmptyCatalogsAndAddons(t *testing.T) {
	client := newKKSTestClient(t, http.MethodGet, "/v1/kks/flavors", "", "", `[]`)
	flavors, err := client.KKS.V1.Flavor.List(context.Background(), krutrim.KKSV1FlavorListParams{})
	if err != nil || flavors == nil || len(flavors) != 0 {
		t.Fatalf("empty flavors = %+v, err = %v", flavors, err)
	}
	client = newKKSTestClient(t, http.MethodGet, "/v1/kks/addons", "", "", `{"addons":[]}`)
	catalog, err := client.KKS.V1.Addon.List(context.Background(), krutrim.KKSV1AddonListParams{})
	if err != nil || catalog == nil || len(catalog.Addons) != 0 {
		t.Fatalf("empty catalog = %+v, err = %v", catalog, err)
	}
	client = newKKSTestClient(t, http.MethodGet, "/v1/kks/clusters/"+kksTestClusterKRN+"/addons", "", "", `{"addons":[]}`)
	installed, err := client.KKS.V1.Cluster.Addon.List(context.Background(), kksTestClusterKRN, krutrim.KKSV1ClusterAddonListParams{})
	if err != nil || installed == nil || len(installed.Addons) != 0 {
		t.Fatalf("empty installed add-ons = %+v, err = %v", installed, err)
	}
}

type kksTrackedBody struct {
	*strings.Reader
	closed bool
}

func (r *kksTrackedBody) Close() error { r.closed = true; return nil }

func TestKKSDelete(t *testing.T) {
	clusterKRN := kksTestClusterKRN + "/cluster?tag=1#fragment"
	resourceKRN := "krn:kks::account:user:resource:example/part?tag=1#fragment"
	clusterPath := "/v1/kks/clusters/" + url.PathEscape(clusterKRN)
	operations := []struct {
		name string
		path string
		call func(*krutrim.Client, ...option.RequestOption) error
	}{
		{"cluster", clusterPath, func(client *krutrim.Client, opts ...option.RequestOption) error {
			return client.KKS.V1.Cluster.Delete(context.Background(), clusterKRN, krutrim.KKSV1ClusterDeleteParams{XRegion: "parameter-region"}, opts...)
		}},
		{"node group", clusterPath + "/node-groups/" + url.PathEscape(resourceKRN), func(client *krutrim.Client, opts ...option.RequestOption) error {
			return client.KKS.V1.Cluster.NodeGroup.Delete(context.Background(), clusterKRN, resourceKRN, krutrim.KKSV1ClusterNodeGroupDeleteParams{XRegion: "parameter-region"}, opts...)
		}},
		{"add-on", clusterPath + "/addons/" + url.PathEscape(resourceKRN), func(client *krutrim.Client, opts ...option.RequestOption) error {
			return client.KKS.V1.Cluster.Addon.Delete(context.Background(), clusterKRN, resourceKRN, krutrim.KKSV1ClusterAddonDeleteParams{XRegion: "parameter-region"}, opts...)
		}},
	}
	responses := []struct {
		name        string
		status      int
		contentType string
		body        string
	}{
		{"no content", http.StatusNoContent, "", ""},
		{"empty JSON body", http.StatusOK, "application/json", ""},
		{"accepted with JSON", http.StatusAccepted, "application/json", `{"message":"deletion requested"}`},
		{"plain text", http.StatusOK, "text/plain", "deletion requested"},
		{"forbidden", http.StatusForbidden, "application/json", `{"message":"denied"}`},
	}
	for _, operation := range operations {
		for _, response := range responses {
			t.Run(operation.name+"/"+response.name, func(t *testing.T) {
				body := &kksTrackedBody{Reader: strings.NewReader(response.body)}
				client := krutrim.NewClient(
					option.WithBaseURL("https://kks.example.test"), option.WithAPIKey("test-token"), option.WithMaxRetries(0),
					option.WithHTTPClient(kksTestHTTPClient(func(req *http.Request) (*http.Response, error) {
						if req.Method != http.MethodDelete || req.URL.EscapedPath() != operation.path || req.URL.RawQuery != "" || req.Body != nil {
							t.Errorf("unexpected delete request: %s %s", req.Method, req.URL)
						}
						if req.Header.Get("Authorization") != "Bearer test-token" || req.Header.Get("x-region") != "override-region" {
							t.Error("authentication or region options were not applied")
						}
						return &http.Response{StatusCode: response.status, Header: http.Header{"Content-Type": {response.contentType}}, Body: body}, nil
					})),
				)
				var raw *http.Response
				err := operation.call(&client, option.WithHeader("x-region", "override-region"), option.WithResponseInto(&raw))
				if response.status >= 400 {
					var apiErr *krutrim.Error
					if !errors.As(err, &apiErr) || apiErr.StatusCode != response.status {
						t.Fatalf("expected SDK API error, got %v", err)
					}
				} else if err != nil {
					t.Fatal(err)
				}
				if raw == nil || raw.StatusCode != response.status {
					t.Errorf("raw response option was not preserved: %+v", raw)
				}
				if !body.closed || body.Len() != 0 {
					t.Error("delete response body was not read and closed")
				}
			})
		}
	}
}

func TestKKSMissingResourceKRN(t *testing.T) {
	client := krutrim.NewClient(option.WithHTTPClient(kksTestHTTPClient(func(req *http.Request) (*http.Response, error) {
		t.Fatal("missing resource KRN must fail before sending a request")
		return nil, nil
	})))
	tests := []struct {
		name  string
		field string
		call  func() error
	}{
		{"get node group", "nodeGroupKRN", func() error {
			_, err := client.KKS.V1.Cluster.NodeGroup.Get(context.Background(), kksTestClusterKRN, "", krutrim.KKSV1ClusterNodeGroupGetParams{})
			return err
		}},
		{"update node group config", "nodeGroupKRN", func() error {
			_, err := client.KKS.V1.Cluster.NodeGroup.UpdateConfig(context.Background(), kksTestClusterKRN, "", krutrim.KKSV1ClusterNodeGroupUpdateConfigParams{})
			return err
		}},
		{"delete node group", "nodeGroupKRN", func() error {
			return client.KKS.V1.Cluster.NodeGroup.Delete(context.Background(), kksTestClusterKRN, "", krutrim.KKSV1ClusterNodeGroupDeleteParams{})
		}},
		{"delete add-on", "addonKRN", func() error {
			return client.KKS.V1.Cluster.Addon.Delete(context.Background(), kksTestClusterKRN, "", krutrim.KKSV1ClusterAddonDeleteParams{})
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.call(); err == nil || !strings.Contains(err.Error(), test.field) {
				t.Fatalf("expected missing %s error, got %v", test.field, err)
			}
		})
	}
}
