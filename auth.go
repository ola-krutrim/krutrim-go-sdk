// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package krutrim

import (
	"context"
	"net/http"
	"slices"

	"github.com/ola-krutrim/krutrim-go-sdk/internal/apijson"
	"github.com/ola-krutrim/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
	"github.com/ola-krutrim/krutrim-go-sdk/packages/param"
)


type IAMService struct {
	Options []option.RequestOption
}


func NewIAMService(opts ...option.RequestOption) (r IAMService) {
	r = IAMService{}
	r.Options = opts
	return
}



type IAMSignInParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	paramObj
}

func (r IAMSignInParams) MarshalJSON() ([]byte, error) {
	type shadow IAMSignInParams
	return param.MarshalObject(r, (*shadow)(&r))
}

func (r *IAMSignInParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IAMSignInResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Message      string `json:"message"`
	MFAEnabled   int    `json:"mfaEnabled"`
}

func (r *IAMService) SignInAsRootUser(
	ctx context.Context,
	body IAMSignInParams,
	opts ...option.RequestOption,
) (res *IAMSignInResponse, err error) {

	opts = slices.Concat(r.Options, opts)

	opts = append(opts,
		option.WithHeader("Content-Type", "application/json"),
		option.WithHeader("Accept", "application/json"),
	)

	path := "iam/v1/signInAsRootUser"

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


type IAMUserSignInParams struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	AccountID string `json:"accountId"`
	paramObj
}

func (r IAMUserSignInParams) MarshalJSON() ([]byte, error) {
	type shadow IAMUserSignInParams
	return param.MarshalObject(r, (*shadow)(&r))
}

func (r *IAMUserSignInParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IAMUserSignInResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Message      string `json:"message"`
	MFAEnabled   int    `json:"mfaEnabled"`
}

func (r *IAMService) SignInAsIAMUser(
	ctx context.Context,
	body IAMUserSignInParams,
	opts ...option.RequestOption,
) (res *IAMUserSignInResponse, err error) {

	opts = slices.Concat(r.Options, opts)

	opts = append(opts,
		option.WithHeader("Content-Type", "application/json"),
		option.WithHeader("Accept", "application/json"),
	)

	path := "iam/v1/signInAsIAMUser"

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
