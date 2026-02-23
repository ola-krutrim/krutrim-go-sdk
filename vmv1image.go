// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package krutrim

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/ola-silicon/krutrim-go-sdk/internal/apiform"
	"github.com/ola-silicon/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-silicon/krutrim-go-sdk/option"
)

// VmV1ImageService contains methods and other services that help with interacting
// with the krutrim-vm-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVmV1ImageService] method instead.
type VmV1ImageService struct {
	Options []option.RequestOption
}

// NewVmV1ImageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVmV1ImageService(opts ...option.RequestOption) (r VmV1ImageService) {
	r = VmV1ImageService{}
	r.Options = opts
	return
}

// List VM images available in a region
func (r *VmV1ImageService) List(ctx context.Context, region string, opts ...option.RequestOption) (res string, err error) {
	opts = slices.Concat(r.Options, opts)
	if region == "" {
		err = errors.New("missing required region parameter")
		return
	}
	opts = append(opts,
		option.WithHeader("Accept", "application/json"),
	)
	path := fmt.Sprintf("vm/v1/image/%s", region)
	var raw json.RawMessage
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &raw, opts...)
	res = string(raw)
	return
}

// Delete a VM image by its KRN
func (r *VmV1ImageService) Delete(ctx context.Context, imageKrn string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if imageKrn == "" {
		err = errors.New("missing required imageKrn parameter")
		return
	}
	path := fmt.Sprintf("vm/v1/image/%s", imageKrn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get list of supported disk formats for images
func (r *VmV1ImageService) GetDiskFormats(ctx context.Context, opts ...option.RequestOption) (res *[]string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "vm/v1/image/disk-formats"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a new VM image file to a region
func (r *VmV1ImageService) Upload(ctx context.Context, region string, body VmV1ImageUploadParams, opts ...option.RequestOption) (res *VmV1ImageUploadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if region == "" {
		err = errors.New("missing required region parameter")
		return
	}
	path := fmt.Sprintf("vm/v1/image/%s", region)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type VmV1ImageListResponse map[string]any

type VmV1ImageUploadResponse map[string]any

type VmV1ImageUploadParams struct {
	// Disk format like qcow2
	DiskFormat string `json:"diskFormat,required"`
	// Image file binary data
	Image string `json:"image,required"`
	paramObj
}

func (r VmV1ImageUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
