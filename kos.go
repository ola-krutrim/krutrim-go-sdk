package krutrim

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ola-krutrim/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
)


type KoService struct {
	Options    []option.RequestOption
	V1         *KoV1Service
}

func NewKoService(opts ...option.RequestOption) KoService {
	return KoService{
		Options: opts,
		V1:      NewKoV1Service(opts...),
	}
}



type KoV1Service struct {
	Options    []option.RequestOption
	Buckets    *KoV1BucketService
	Objects    *KoV1ObjectService
	AccessKeys *KoV1AccessKeyService
	Sessions   *KoV1SessionService
}

func NewKoV1Service(opts ...option.RequestOption) *KoV1Service {
	return &KoV1Service{
		Options:    opts,
		Buckets:    &KoV1BucketService{Options: opts},
		Objects:    &KoV1ObjectService{Options: opts},
		AccessKeys: &KoV1AccessKeyService{Options: opts},
		Sessions:   &KoV1SessionService{Options: opts},
	}
}



type KoV1BucketService struct {
	Options []option.RequestOption
}


type BucketCreateParams struct {
	XRegion string `json:"-"`
	XTier   string `json:"-"`
    Description     string            `json:"description,omitempty"`
	Name            string            `json:"name"`
	Tier            string            `json:"tier"`
	AnonymousAccess bool              `json:"anonymous_access,omitempty"`
	Versioning      bool              `json:"versioning,omitempty"`
	Tags            map[string]string `json:"tags,omitempty"`
}

type BucketListResponse struct {
	TotalItems int              `json:"totalItems"`
	Items      []map[string]any `json:"items"`
}

type BucketCreateResponse struct {
	KRN        string `json:"krn"`
	BucketName string `json:"bucketName"`
	CreatedAt  string `json:"createdAt"`
}
type BucketListParams struct {
	Tier string
}

type ObjectRenameResponse struct {
	Message string `json:"message"`
	OldKey  string `json:"oldKey"`
	NewKey  string `json:"newKey"`
}
type ObjectRenameParams struct {
	XRegion        string `json:"-"`
	XSessionToken string `json:"-"`

	BucketKRN string
	OldKey    string
	NewKey    string
}

type ObjectMoveParams struct {
	XRegion        string `json:"-"`
	XSessionToken string `json:"-"`

	BucketKRN      string
	SourceKey      string
	DestinationKey string
}

func (r *KoV1BucketService) Create(
	ctx context.Context,
	body BucketCreateParams,
	opts ...option.RequestOption,
) (*BucketCreateResponse, error) {

	if body.XRegion != "" {
		opts = append(opts, option.WithHeader("x-region-id", body.XRegion))
	}
	if body.XTier != "" {
		opts = append(opts, option.WithHeader("x-tier", body.XTier))
	}

	opts = append(r.Options, opts...)

	var res *BucketCreateResponse
	err := requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodPost,
		"kos/v1/buckets",
		body,
		&res,
		opts...,
	)
	return res, err
}


func (r *KoV1BucketService) Get(
	ctx context.Context,
	bucketKRN string,
	opts ...option.RequestOption,
) (map[string]any, error) {

	opts = append(r.Options, opts...)
	path := fmt.Sprintf("kos/v1/buckets/%s", url.PathEscape(bucketKRN))

	var res map[string]any
	err := requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodGet,
		path,
		nil,
		&res,
		opts...,
	)
	return res, err
}


func (r *KoV1BucketService) List(
	ctx context.Context,
	params *BucketListParams,
	opts ...option.RequestOption,
) ([]map[string]any, error) {

	opts = append(r.Options, opts...)

	path := "kos/v1/buckets"
	if params != nil && params.Tier != "" {
		path = fmt.Sprintf("%s?tier=%s", path, url.QueryEscape(params.Tier))
	}

	var res BucketListResponse
	err := requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodGet,
		path,
		nil,
		&res,
		opts...,
	)

	if err != nil {
		return nil, err
	}

	// Check if response empty
	if res.Items == nil || len(res.Items) == 0 {
		return []map[string]any{}, fmt.Errorf("no buckets exist in the specified region")
	}

	return res.Items, nil
}


func (r *KoV1BucketService) Delete(
	ctx context.Context,
	bucketKRN string,
	xRegion string,
	xTier string,
	opts ...option.RequestOption,
) error {

	if xRegion != "" {
		opts = append(opts, option.WithHeader("x-region-id", xRegion))
	}
	if xTier != "" {
		opts = append(opts, option.WithHeader("x-tier", xTier))
	}

	opts = append(r.Options, opts...)
	path := fmt.Sprintf("kos/v1/buckets/%s", url.PathEscape(bucketKRN))

	return requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodDelete,
		path,
		nil,
		nil,
		opts...,
	)
}


type KoV1AccessKeyService struct {
	Options []option.RequestOption
}


func (r *KoV1AccessKeyService) Create(
	ctx context.Context,
	xRegion string,
	xTier string,
	opts ...option.RequestOption,
) (map[string]string, error) {

	if xRegion != "" {
		opts = append(opts, option.WithHeader("x-region-id", xRegion))
	}
	if xTier != "" {
		opts = append(opts, option.WithHeader("x-tier", xTier))
	}

	opts = append(r.Options, opts...)

	var res map[string]string
	err := requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodPost,
		"kos/v1/access_keys",
		map[string]any{},
		&res,
		opts...,
	)
	return res, err
}


func (r *KoV1AccessKeyService) List(
	ctx context.Context,
	xRegion string,
	xTier string,
	opts ...option.RequestOption,
) ([]map[string]any, error) {

	if xRegion != "" {
		opts = append(opts, option.WithHeader("x-region-id", xRegion))
	}
	if xTier != "" {
		opts = append(opts, option.WithHeader("x-tier", xTier))
	}

	opts = append(r.Options, opts...)

	var res []map[string]any
	err := requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodGet,
		"kos/v1/access_keys",
		nil,
		&res,
		opts...,
	)
	return res, err
}

func (r *KoV1AccessKeyService) Delete(
	ctx context.Context,
	accessKey string,
	xRegion string,
	xTier string,
	opts ...option.RequestOption,
) error {

	if xRegion != "" {
		opts = append(opts, option.WithHeader("x-region-id", xRegion))
	}
	if xTier != "" {
		opts = append(opts, option.WithHeader("x-tier", xTier))
	}

	opts = append(r.Options, opts...)
	path := fmt.Sprintf("kos/v1/access_keys/%s", accessKey)

	return requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodDelete,
		path,
		nil,
		nil,
		opts...,
	)
}



type KoV1SessionService struct {
	Options []option.RequestOption
}

type ActivateSessionParams struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

func (r *KoV1SessionService) Activate(
	ctx context.Context,
	xRegion string,
	xTier string,
	body ActivateSessionParams,
	opts ...option.RequestOption,
) (map[string]string, error) {

	if xRegion != "" {
		opts = append(opts, option.WithHeader("x-region-id", xRegion))
	}
	if xTier != "" {
		opts = append(opts, option.WithHeader("x-tier", xTier))
	}

	opts = append(r.Options, opts...)

	var res map[string]string
	err := requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodPost,
		"kos/v1/sessions/activate",
		body,
		&res,
		opts...,
	)
	return res, err
}


type KoV1ObjectService struct {
	Options []option.RequestOption
}


func (r *KoV1ObjectService) List(
	ctx context.Context,
	bucketKRN string,
	xRegion string,
	xSessionToken string,
	opts ...option.RequestOption,
) ([]map[string]any, error) {

	if xRegion != "" {
		opts = append(opts, option.WithHeader("x-region-id", xRegion))
	}
	if xSessionToken != "" {
		opts = append(opts, option.WithHeader("x-session-token", xSessionToken))
	}

	opts = append(r.Options, opts...)

	path := fmt.Sprintf(
		"kos/v1/buckets/%s/objects",
		url.PathEscape(bucketKRN),
	)

	var res []map[string]any
	err := requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodGet,
		path,
		nil,
		&res,
		opts...,
	)
	return res, err
}

func (r *KoV1ObjectService) GetUploadURL(
	ctx context.Context,
	bucketKRN string,
	objectKey string,
	opts ...option.RequestOption,
) (map[string]string, error) {

	opts = append(r.Options, opts...)
	path := fmt.Sprintf("kos/v1/buckets/%s/objects", url.PathEscape(bucketKRN))

	var res map[string]string
	err := requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodPut,
		path,
		map[string]string{"prefix": objectKey},
		&res,
		opts...,
	)
	return res, err
}

func (r *KoV1ObjectService) GetDownloadURL(
	ctx context.Context,
	bucketKRN string,
	objectKey string,
	opts ...option.RequestOption,
) (map[string]string, error) {

	opts = append(r.Options, opts...)
	path := fmt.Sprintf("kos/v1/buckets/%s/objects/download", url.PathEscape(bucketKRN))

	var res map[string]string
	err := requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodPost,
		path,
		map[string]string{"objectkey": objectKey},
		&res,
		opts...,
	)
	return res, err
}

func (r *KoV1ObjectService) Rename(
	ctx context.Context,
	params ObjectRenameParams,
	opts ...option.RequestOption,
) (*ObjectRenameResponse, error) {

	if params.XRegion != "" {
		opts = append(opts, option.WithHeader("x-region-id", params.XRegion))
	}
	if params.XSessionToken != "" {
		opts = append(opts, option.WithHeader("x-session-token", params.XSessionToken))
	}

	opts = append(r.Options, opts...)

	path := fmt.Sprintf(
		"kos/v1/buckets/%s/objects/rename",
		url.PathEscape(params.BucketKRN),
	)

	var res ObjectRenameResponse
	err := requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodPost,
		path,
		map[string]string{
			"oldKey": params.OldKey,
			"newKey": params.NewKey,
		},
		&res,
		opts...,
	)

	return &res, err
}

func (r *KoV1ObjectService) Move(
	ctx context.Context,
	bucketKRN string,
	src string,
	dst string,
	opts ...option.RequestOption,
) error {

	opts = append(r.Options, opts...)
	path := fmt.Sprintf("kos/v1/buckets/%s/objects/move", url.PathEscape(bucketKRN))

	return requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodPost,
		path,
		map[string]string{
			"sourceKey":      src,
			"destinationKey": dst,
		},
		nil,
		opts...,
	)
}

func (r *KoV1ObjectService) Delete(
	ctx context.Context,
	bucketKRN string,
	objectKey string,
	opts ...option.RequestOption,
) error {

	opts = append(r.Options, opts...)
	path := fmt.Sprintf("kos/v1/buckets/%s/objects/delete", url.PathEscape(bucketKRN))

	return requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodPost,
		path,
		map[string]string{"objectkey": objectKey},
		nil,
		opts...,
	)
}
func (r *KoV1ObjectService) InitUpload(
	ctx context.Context,
	bucketKRN string,
	objectKey string,
	xRegion string,
	xSessionToken string,
	opts ...option.RequestOption,
) (map[string]string, error) {

	if xRegion != "" {
		opts = append(opts, option.WithHeader("x-region-id", xRegion))
	}
	if xSessionToken != "" {
		opts = append(opts, option.WithHeader("x-session-token", xSessionToken))
	}

	opts = append(r.Options, opts...)

	path := fmt.Sprintf(
		"kos/v1/buckets/%s/objects",
		url.PathEscape(bucketKRN),
	)

	var res map[string]string
	err := requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodPut,
		path,
		map[string]string{"prefix": objectKey},
		&res,
		opts...,
	)
	return res, err
}



func (r *KoV1ObjectService) GetPreSignedDownloadURL(
	ctx context.Context,
	bucketKRN string,
	objectKey string,
	xRegion string,
	xSessionToken string,
	opts ...option.RequestOption,
) (map[string]string, error) {

	if xRegion != "" {
		opts = append(opts, option.WithHeader("x-region-id", xRegion))
	}
	if xSessionToken != "" {
		opts = append(opts, option.WithHeader("x-session-token", xSessionToken))
	}

	opts = append(r.Options, opts...)

	path := fmt.Sprintf(
		"kos/v1/buckets/%s/objects/url",
		url.PathEscape(bucketKRN),
	)

	var res map[string]string
	err := requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodPost,
		path,
		map[string]string{"objectkey": objectKey},
		&res,
		opts...,
	)
	return res, err
}



type BucketUpdateParams struct {
	XRegion string `json:"-"`
	XTier   string `json:"-"`

	BucketKRN string `json:"-"`

	Versioning      bool              `json:"versioning,omitempty"`
	AnonymousAccess bool              `json:"anonymous_access,omitempty"`
	Tags             map[string]string `json:"tags,omitempty"`
}



func (r *KoV1BucketService) Update(
	ctx context.Context,
	params BucketUpdateParams,
	opts ...option.RequestOption,
) (map[string]any, error) {

	if params.XRegion != "" {
		opts = append(opts, option.WithHeader("x-region-id", params.XRegion))
	}
	if params.XTier != "" {
		opts = append(opts, option.WithHeader("x-tier", params.XTier))
	}

	opts = append(r.Options, opts...)

	path := fmt.Sprintf(
		"kos/v1/buckets/%s",
		url.PathEscape(params.BucketKRN),
	)

	var res map[string]any
	err := requestconfig.ExecuteNewRequest(
		ctx,
		http.MethodPut,
		path,
		map[string]any{
			"versioning":       params.Versioning,
			"anonymous_access": params.AnonymousAccess,
			"tags":             params.Tags,
		},
		&res,
		opts...,
	)

	return res, err
}