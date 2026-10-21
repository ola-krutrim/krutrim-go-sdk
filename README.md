# Krutrim Client Go API Library

<!-- x-release-please-start-version -->

<a href="https://pkg.go.dev/github.com/ola-krutrim/krutrim-go-sdk"><img src="https://pkg.go.dev/badge/github.com/ola-krutrim/krutrim-go-sdk.svg" alt="Go Reference"></a>

<!-- x-release-please-end -->

The Krutrim Client Go library provides convenient access to the Krutrim Client REST API
from applications written in Go.

It is generated with [Stainless](https://www.stainless.com/).

## Installation

```go
import (
	"github.com/ola-krutrim/krutrim-go-sdk" // imported as krutrimclient
)
```

Or to pin the version:

```sh
go get -u 'github.com/ola-krutrim/krutrim-go-sdk@v0.0.1'
```

## Requirements

This library requires Go 1.22+.

## Usage

The full API of this library can be found in [api.md](api.md).

```go

package krutrim

import (
	"context"

	"github.com/ola-krutrim/krutrim-go-sdk"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
)

func main() {
	client := krutrimclient.NewClient(
		option.WithAPIKey("My API Key"), // defaults to os.LookupEnv("KRUTRIM_CLIENT_API_KEY")
	)
	err := client.DeleteSubnet.Delete(context.TODO(), krutrimclient.DeleteSubnetDeleteParams{
		SubnetID: "REPLACE_ME",
		VpcID:    "vpc_id",
		XRegion:  "x-region",
	})
	if err != nil {
		panic(err.Error())
	}
}

```

### Krutrim Kubernetes Service (KKS)

KKS is available through `client.KKS.V1`. Cluster operations use
`client.KKS.V1.Cluster`, while the flavor and add-on catalogs use
`client.KKS.V1.Flavor` and `client.KKS.V1.Addon`. All methods accept request
options and use the client's existing authentication, base URL, retries, and
context handling. Set `KRUTRIM_CLIENT_API_KEY` to your bearer token (without the
`Bearer ` prefix). Pass raw, unescaped KRNs to SDK methods.

| Operation | SDK method | HTTP endpoint |
| --- | --- | --- |
| Create cluster | `client.KKS.V1.Cluster.Create` | `POST /v1/kks/clusters` |
| List clusters | `client.KKS.V1.Cluster.List` | `GET /v1/kks/clusters` |
| Get cluster | `client.KKS.V1.Cluster.Get` | `GET /v1/kks/clusters/{clusterKRN}` |
| Delete cluster | `client.KKS.V1.Cluster.Delete` | `DELETE /v1/kks/clusters/{clusterKRN}` |
| Download kubeconfig | `client.KKS.V1.Cluster.GetKubeconfig` | `GET /v1/kks/clusters/{clusterKRN}/kubeconfig` |
| Create node group | `client.KKS.V1.Cluster.NodeGroup.Create` | `POST /v1/kks/clusters/{clusterKRN}/node-groups` |
| List node groups | `client.KKS.V1.Cluster.NodeGroup.List` | `GET /v1/kks/clusters/{clusterKRN}/node-groups` |
| Get node group | `client.KKS.V1.Cluster.NodeGroup.Get` | `GET /v1/kks/clusters/{clusterKRN}/node-groups/{nodeGroupKRN}` |
| Update node group config | `client.KKS.V1.Cluster.NodeGroup.UpdateConfig` | `POST /v1/kks/clusters/{clusterKRN}/node-groups/{nodeGroupKRN}/update-config` |
| Delete node group | `client.KKS.V1.Cluster.NodeGroup.Delete` | `DELETE /v1/kks/clusters/{clusterKRN}/node-groups/{nodeGroupKRN}` |
| Create add-on | `client.KKS.V1.Cluster.Addon.Create` | `POST /v1/kks/clusters/{clusterKRN}/addons` |
| List installed add-ons | `client.KKS.V1.Cluster.Addon.List` | `GET /v1/kks/clusters/{clusterKRN}/addons` |
| Delete add-on | `client.KKS.V1.Cluster.Addon.Delete` | `DELETE /v1/kks/clusters/{clusterKRN}/addons/{addonKRN}` |
| List available add-ons | `client.KKS.V1.Addon.List` | `GET /v1/kks/addons` |
| List flavors | `client.KKS.V1.Flavor.List` | `GET /v1/kks/flavors` |

For example, with imports for `context`, `log`, `os`, `krutrim`, and `option`:

```go
client := krutrim.NewClient(
	option.WithHeader("x-region", "In-Bangalore-1"),
)
ctx := context.Background()

created, err := client.KKS.V1.Cluster.Create(ctx, krutrim.KKSV1ClusterCreateParams{
	Name:    "test",
	Version: "1.36",
	ResourcesVPCConfig: krutrim.KKSResourcesVPCConfig{
		VPCKRN:                "YOUR_VPC_KRN",
		SubnetKRNs:            "YOUR_SUBNET_KRN",
		EndpointPrivateAccess: false,
	},
	KubernetesNetworkConfig: krutrim.KKSKubernetesNetworkConfig{
		PodIPv4CIDR:     "10.244.0.0/16",
		ServiceIPv4CIDR: "10.96.0.0/12",
	},
})
if err != nil {
	log.Fatal(err)
}
clusterKRN := created.Cluster.KRN

clusters, err := client.KKS.V1.Cluster.List(ctx, krutrim.KKSV1ClusterListParams{
	Region: krutrim.String("In-Bangalore-1"), // Omit to list without a region filter.
})
if err != nil {
	log.Fatal(err)
}
_ = clusters.Clusters

cluster, err := client.KKS.V1.Cluster.Get(ctx, clusterKRN, krutrim.KKSV1ClusterGetParams{})
if err != nil {
	log.Fatal(err)
}
_ = cluster.Cluster.Status
```

Creation is asynchronous. Check cluster status with `Get` and wait for
`PROVISIONED` before adding node groups or add-ons. The following calls are
independent examples to run once the cluster is ready:

```go
nodegroup, err := client.KKS.V1.Cluster.NodeGroup.Create(ctx, clusterKRN, krutrim.KKSV1ClusterNodeGroupCreateParams{
	Name:          "default-pool",
	InstanceTypes: "CPU-4x-16GB",
	DiskSize:      166,
	ScalingConfig: krutrim.KKSScalingConfig{MinSize: 1, MaxSize: 5, DesiredSize: 2},
	SubnetsKRN:    "YOUR_SUBNET_KRN",
	Labels:        map[string]string{},
	Taints:        []krutrim.KKSTaint{},
	RemoteAccess: &krutrim.KKSRemoteAccess{
		SourceSecurityGroupsKRNs: []string{},
	},
	NodeRepairConfig: &krutrim.KKSNodeRepairConfig{Enabled: true},
})
if err != nil {
	log.Fatal(err)
}
_ = nodegroup.Nodegroup.KRN

nodegroups, err := client.KKS.V1.Cluster.NodeGroup.List(ctx, clusterKRN, krutrim.KKSV1ClusterNodeGroupListParams{})
if err != nil {
	log.Fatal(err)
}
_ = nodegroups.Nodegroups

addon, err := client.KKS.V1.Cluster.Addon.Create(ctx, clusterKRN, krutrim.KKSV1ClusterAddonCreateParams{
	AddonName:    "csi",
	AddonVersion: krutrim.String("1.0.0"), // Omit to use the server's default.
})
if err != nil {
	log.Fatal(err)
}
_ = addon.Addon.AddonKRN

kubeconfig, err := client.KKS.V1.Cluster.GetKubeconfig(ctx, clusterKRN, krutrim.KKSV1ClusterGetKubeconfigParams{})
if err != nil {
	log.Fatal(err)
}
if err := os.WriteFile("kubeconfig", kubeconfig, 0600); err != nil {
	log.Fatal(err)
}
```

`XRegion` on each method's parameter struct sets the `x-region` header; the list
method's `Region` sets the separate `region` query parameter. Request options can
override either the client header or `XRegion`. `SubnetKRNs`, `SubnetsKRN`, and
`InstanceTypes` are strings, matching the API payloads. Optional node-group fields
are omitted when unset, while explicitly empty maps and slices are sent as `{}`
and `[]`. Use a pointer to `KKSNodeRepairConfig{Enabled: false}` to explicitly
disable repair. Kubeconfig is returned as raw bytes, preserving the response format.

`Flavor.List` returns a slice because `/v1/kks/flavors` returns a top-level JSON
array. `Addon.List` returns the available catalog, including descriptions,
versions, architectures, and cluster compatibility. `Cluster.Addon.List` returns
installed add-ons with Unix timestamps; the add-on create response uses a
`time.Time` for its ISO timestamp. Delete methods return an error and accept both
empty and non-empty success responses. Use `option.WithResponseInto` to inspect
the HTTP status when needed.

`Cluster.NodeGroup.UpdateConfig` accepts a `KKSV1ClusterNodeGroupUpdateConfigParams`
with `ScalingConfig` and optional `XRegion`. The response's `Update` contains
`Status`, `Type`, fractional Unix `CreatedAt`, and `Errors`. Error items are
preserved as `json.RawMessage` because their schema is unspecified.

Each KKS operation has its own standalone example. Fill in the resource KRNs
directly in the file, for example:

```go
vpcKRN := "enter the vpckrn"
subnetKRN := "enter the subnetkrn"
region := "In-Bangalore-1" // Replace with your region.
```

Only `KRUTRIM_CLIENT_API_KEY` is read from the environment by the examples;
the examples also load `.env` when present. Set the key and run one file at a time:

```sh
export KRUTRIM_CLIENT_API_KEY='YOUR_BEARER_TOKEN'
go run ./examples/kks/list_clusters.go
go run ./examples/kks/create_cluster.go
go run ./examples/kks/create_addon.go
```

The region is set directly in each file. Fill in these resource variables for the
operation you want to run:

| Example | Resource variables to fill in |
| --- | --- |
| [create_cluster.go](examples/kks/create_cluster.go) | `vpcKRN`, `subnetKRN` |
| [list_clusters.go](examples/kks/list_clusters.go) | None |
| [get_cluster.go](examples/kks/get_cluster.go) | `clusterKRN` |
| [delete_cluster.go](examples/kks/delete_cluster.go) | `clusterKRN` |
| [get_kubeconfig.go](examples/kks/get_kubeconfig.go) | `clusterKRN` |
| [create_node_group.go](examples/kks/create_node_group.go) | `clusterKRN`, `subnetKRN` |
| [list_node_groups.go](examples/kks/list_node_groups.go) | `clusterKRN` |
| [get_node_group.go](examples/kks/get_node_group.go) | `clusterKRN`, `nodeGroupKRN` |
| [update_node_group_config.go](examples/kks/update_node_group_config.go) | `clusterKRN`, `nodeGroupKRN` |
| [delete_node_group.go](examples/kks/delete_node_group.go) | `clusterKRN`, `nodeGroupKRN` |
| [create_addon.go](examples/kks/create_addon.go) | `clusterKRN` |
| [list_cluster_addons.go](examples/kks/list_cluster_addons.go) | `clusterKRN` |
| [delete_addon.go](examples/kks/delete_addon.go) | `clusterKRN`, `addonKRN` |
| [list_addons.go](examples/kks/list_addons.go) | None |
| [list_flavors.go](examples/kks/list_flavors.go) | None |

These files use `//go:build ignore` so their independent `main` functions are
excluded from package builds. Explicit `go run` or `go test` commands naming one
example file still compile it normally.

### Request fields

The krutrimclient library uses the [`omitzero`](https://tip.golang.org/doc/go1.24#encodingjsonpkgencodingjson)
semantics from the Go 1.24+ `encoding/json` release for request fields.

Required primitive fields (`int64`, `string`, etc.) feature the tag <code>\`json:"...,required"\`</code>. These
fields are always serialized, even their zero values.

Optional primitive types are wrapped in a `param.Opt[T]`. These fields can be set with the provided constructors, `krutrimclient.String(string)`, `krutrimclient.Int(int64)`, etc.

Any `param.Opt[T]`, map, slice, struct or string enum uses the
tag <code>\`json:"...,omitzero"\`</code>. Its zero value is considered omitted.

The `param.IsOmitted(any)` function can confirm the presence of any `omitzero` field.

```go
p := krutrimclient.ExampleParams{
	ID:   "id_xxx",                    // required property
	Name: krutrimclient.String("..."), // optional property

	Point: krutrimclient.Point{
		X: 0,                    // required field will serialize as 0
		Y: krutrimclient.Int(1), // optional field will serialize as 1
		// ... omitted non-required fields will not be serialized
	},

	Origin: krutrimclient.Origin{}, // the zero value of [Origin] is considered omitted
}
```

To send `null` instead of a `param.Opt[T]`, use `param.Null[T]()`.
To send `null` instead of a struct `T`, use `param.NullStruct[T]()`.

```go
p.Name = param.Null[string]()       // 'null' instead of string
p.Point = param.NullStruct[Point]() // 'null' instead of struct

param.IsNull(p.Name)  // true
param.IsNull(p.Point) // true
```

Request structs contain a `.SetExtraFields(map[string]any)` method which can send non-conforming
fields in the request body. Extra fields overwrite any struct fields with a matching
key. For security reasons, only use `SetExtraFields` with trusted data.

To send a custom value instead of a struct, use `param.Override[T](value)`.

```go
// In cases where the API specifies a given type,
// but you want to send something else, use [SetExtraFields]:
p.SetExtraFields(map[string]any{
	"x": 0.01, // send "x" as a float instead of int
})

// Send a number instead of an object
custom := param.Override[krutrimclient.FooParams](12)
```

### Request unions

Unions are represented as a struct with fields prefixed by "Of" for each of its variants,
only one field can be non-zero. The non-zero field will be serialized.

Sub-properties of the union can be accessed via methods on the union struct.
These methods return a mutable pointer to the underlying data, if present.

```go
// Only one field can be non-zero, use param.IsOmitted() to check if a field is set
type AnimalUnionParam struct {
	OfCat *Cat `json:",omitzero,inline`
	OfDog *Dog `json:",omitzero,inline`
}

animal := AnimalUnionParam{
	OfCat: &Cat{
		Name: "Whiskers",
		Owner: PersonParam{
			Address: AddressParam{Street: "3333 Coyote Hill Rd", Zip: 0},
		},
	},
}

// Mutating a field
if address := animal.GetOwner().GetAddress(); address != nil {
	address.ZipCode = 94304
}
```

### Response objects

All fields in response structs are ordinary value types (not pointers or wrappers).
Response structs also include a special `JSON` field containing metadata about
each property.

```go
type Animal struct {
	Name   string `json:"name,nullable"`
	Owners int    `json:"owners"`
	Age    int    `json:"age"`
	JSON   struct {
		Name        respjson.Field
		Owner       respjson.Field
		Age         respjson.Field
		ExtraFields map[string]respjson.Field
	} `json:"-"`
}
```

To handle optional data, use the `.Valid()` method on the JSON field.
`.Valid()` returns true if a field is not `null`, not present, or couldn't be marshaled.

If `.Valid()` is false, the corresponding field will simply be its zero value.

```go
raw := `{"owners": 1, "name": null}`

var res Animal
json.Unmarshal([]byte(raw), &res)

// Accessing regular fields

res.Owners // 1
res.Name   // ""
res.Age    // 0

// Optional field checks

res.JSON.Owners.Valid() // true
res.JSON.Name.Valid()   // false
res.JSON.Age.Valid()    // false

// Raw JSON values

res.JSON.Owners.Raw()                  // "1"
res.JSON.Name.Raw() == "null"          // true
res.JSON.Name.Raw() == respjson.Null   // true
res.JSON.Age.Raw() == ""               // true
res.JSON.Age.Raw() == respjson.Omitted // true
```

These `.JSON` structs also include an `ExtraFields` map containing
any properties in the json response that were not specified
in the struct. This can be useful for API features not yet
present in the SDK.

```go
body := res.JSON.ExtraFields["my_unexpected_field"].Raw()
```

### Response Unions

In responses, unions are represented by a flattened struct containing all possible fields from each of the
object variants.
To convert it to a variant use the `.AsFooVariant()` method or the `.AsAny()` method if present.

If a response value union contains primitive values, primitive fields will be alongside
the properties but prefixed with `Of` and feature the tag `json:"...,inline"`.

```go
type AnimalUnion struct {
	// From variants [Dog], [Cat]
	Owner Person `json:"owner"`
	// From variant [Dog]
	DogBreed string `json:"dog_breed"`
	// From variant [Cat]
	CatBreed string `json:"cat_breed"`
	// ...

	JSON struct {
		Owner respjson.Field
		// ...
	} `json:"-"`
}

// If animal variant
if animal.Owner.Address.ZipCode == "" {
	panic("missing zip code")
}

// Switch on the variant
switch variant := animal.AsAny().(type) {
case Dog:
case Cat:
default:
	panic("unexpected type")
}
```

### RequestOptions

This library uses the functional options pattern. Functions defined in the
`option` package return a `RequestOption`, which is a closure that mutates a
`RequestConfig`. These options can be supplied to the client or at individual
requests. For example:

```go
client := krutrimclient.NewClient(
	// Adds a header to every request made by the client
	option.WithHeader("X-Some-Header", "custom_header_info"),
)

client.DeleteSubnet.Delete(context.TODO(), ...,
	// Override the header
	option.WithHeader("X-Some-Header", "some_other_custom_header_info"),
	// Add an undocumented field to the request body, using sjson syntax
	option.WithJSONSet("some.json.path", map[string]string{"my": "object"}),
)
```

The request option `option.WithDebugLog(nil)` may be helpful while debugging.

See the [full list of request options](https://pkg.go.dev/github.com/ola-krutrim/krutrim-go-sdk/option).

### Pagination

This library provides some conveniences for working with paginated list endpoints.

You can use `.ListAutoPaging()` methods to iterate through items across all pages:

Or you can use simple `.List()` methods to fetch a single page and receive a standard response object
with additional helper methods like `.GetNextPage()`, e.g.:

### Errors

When the API returns a non-success status code, we return an error with type
`*krutrimclient.Error`. This contains the `StatusCode`, `*http.Request`, and
`*http.Response` values of the request, as well as the JSON of the error body
(much like other response objects in the SDK).

To handle errors, we recommend that you use the `errors.As` pattern:

```go
err := client.DeleteSubnet.Delete(context.TODO(), krutrimclient.DeleteSubnetDeleteParams{
	SubnetID: "REPLACE_ME",
	VpcID:    "vpc_id",
	XRegion:  "x-region",
})
if err != nil {
	var apierr *krutrimclient.Error
	if errors.As(err, &apierr) {
		println(string(apierr.DumpRequest(true)))  // Prints the serialized HTTP request
		println(string(apierr.DumpResponse(true))) // Prints the serialized HTTP response
	}
	panic(err.Error()) // GET "/delete_subnet": 400 Bad Request { ... }
}
```

When other errors occur, they are returned unwrapped; for example,
if HTTP transport fails, you might receive `*url.Error` wrapping `*net.OpError`.

### Timeouts

Requests do not time out by default; use context to configure a timeout for a request lifecycle.

Note that if a request is [retried](#retries), the context timeout does not start over.
To set a per-retry timeout, use `option.WithRequestTimeout()`.

```go
// This sets the timeout for the request, including all the retries.
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()
client.DeleteSubnet.Delete(
	ctx,
	krutrimclient.DeleteSubnetDeleteParams{
		SubnetID: "REPLACE_ME",
		VpcID:    "vpc_id",
		XRegion:  "x-region",
	},
	// This sets the per-retry timeout
	option.WithRequestTimeout(20*time.Second),
)
```

### File uploads

Request parameters that correspond to file uploads in multipart requests are typed as
`io.Reader`. The contents of the `io.Reader` will by default be sent as a multipart form
part with the file name of "anonymous_file" and content-type of "application/octet-stream".

The file name and content-type can be customized by implementing `Name() string` or `ContentType()
string` on the run-time type of `io.Reader`. Note that `os.File` implements `Name() string`, so a
file returned by `os.Open` will be sent with the file name on disk.

We also provide a helper `krutrimclient.File(reader io.Reader, filename string, contentType string)`
which can be used to wrap any `io.Reader` with the appropriate file name and content type.

### Retries

Certain errors will be automatically retried 2 times by default, with a short exponential backoff.
We retry by default all connection errors, 408 Request Timeout, 409 Conflict, 429 Rate Limit,
and >=500 Internal errors.

You can use the `WithMaxRetries` option to configure or disable this:

```go
// Configure the default for all requests:
client := krutrimclient.NewClient(
	option.WithMaxRetries(0), // default is 2
)

// Override per-request:
client.DeleteSubnet.Delete(
	context.TODO(),
	krutrimclient.DeleteSubnetDeleteParams{
		SubnetID: "REPLACE_ME",
		VpcID:    "vpc_id",
		XRegion:  "x-region",
	},
	option.WithMaxRetries(5),
)
```

### Accessing raw response data (e.g. response headers)

You can access the raw HTTP response data by using the `option.WithResponseInto()` request option. This is useful when
you need to examine response headers, status codes, or other details.

```go
// Create a variable to store the HTTP response
var response *http.Response
err := client.DeleteSubnet.Delete(
	context.TODO(),
	krutrimclient.DeleteSubnetDeleteParams{
		SubnetID: "REPLACE_ME",
		VpcID:    "vpc_id",
		XRegion:  "x-region",
	},
	option.WithResponseInto(&response),
)
if err != nil {
	// handle error
}
null

fmt.Printf("Status Code: %d\n", response.StatusCode)
fmt.Printf("Headers: %+#v\n", response.Header)
```

### Making custom/undocumented requests

This library is typed for convenient access to the documented API. If you need to access undocumented
endpoints, params, or response properties, the library can still be used.

#### Undocumented endpoints

To make requests to undocumented endpoints, you can use `client.Get`, `client.Post`, and other HTTP verbs.
`RequestOptions` on the client, such as retries, will be respected when making these requests.

```go
var (
    // params can be an io.Reader, a []byte, an encoding/json serializable object,
    // or a "…Params" struct defined in this library.
    params map[string]any

    // result can be an []byte, *http.Response, a encoding/json deserializable object,
    // or a model defined in this library.
    result *http.Response
)
err := client.Post(context.Background(), "/unspecified", params, &result)
if err != nil {
    …
}
```

#### Undocumented request params

To make requests using undocumented parameters, you may use either the `option.WithQuerySet()`
or the `option.WithJSONSet()` methods.

```go
params := FooNewParams{
    ID:   "id_xxxx",
    Data: FooNewParamsData{
        FirstName: krutrimclient.String("John"),
    },
}
client.Foo.New(context.Background(), params, option.WithJSONSet("data.last_name", "Doe"))
```

#### Undocumented response properties

To access undocumented response properties, you may either access the raw JSON of the response as a string
with `result.JSON.RawJSON()`, or get the raw JSON of a particular field on the result with
`result.JSON.Foo.Raw()`.

Any fields that are not present on the response struct will be saved and can be accessed by `result.JSON.ExtraFields()` which returns the extra fields as a `map[string]Field`.

### Middleware

We provide `option.WithMiddleware` which applies the given
middleware to requests.

```go
func Logger(req *http.Request, next option.MiddlewareNext) (res *http.Response, err error) {
	// Before the request
	start := time.Now()
	LogReq(req)

	// Forward the request to the next handler
	res, err = next(req)

	// Handle stuff after the request
	end := time.Now()
	LogRes(res, err, start - end)

    return res, err
}

client := krutrimclient.NewClient(
	option.WithMiddleware(Logger),
)
```

When multiple middlewares are provided as variadic arguments, the middlewares
are applied left to right. If `option.WithMiddleware` is given
multiple times, for example first in the client then the method, the
middleware in the client will run first and the middleware given in the method
will run next.

You may also replace the default `http.Client` with
`option.WithHTTPClient(client)`. Only one http client is
accepted (this overwrites any previous client) and receives requests after any
middleware has been applied.

## Semantic versioning

This package generally follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions, though certain backwards-incompatible changes may be released as minor versions:

1. Changes to library internals which are technically public but not intended or documented for external use. _(Please open a GitHub issue to let us know if you are relying on such internals.)_
2. Changes that we do not expect to impact the vast majority of users in practice.

We take backwards-compatibility seriously and work hard to ensure you can rely on a smooth upgrade experience.

We are keen for your feedback; please open an [issue](https://www.github.com/ola-krutrim/krutrim-go-sdk/issues) with questions, bugs, or suggestions.

## Contributing

See [the contributing documentation](./CONTRIBUTING.md).
