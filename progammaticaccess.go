package krutrim

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/ola-krutrim/krutrim-go-sdk/internal/requestconfig"
	"github.com/ola-krutrim/krutrim-go-sdk/option"
)

//
// =====================================
// FINAL IAM CLIENT
// =====================================
//

type IAMClient struct {
	Options []option.RequestOption
	Token   string
}

func NewIAMClient(opts ...option.RequestOption) IAMClient {
	opts = append(DefaultClientOptions(), opts...)
	return IAMClient{Options: opts}
}

//
// =====================================
// INTERNAL HELPERS
// =====================================
//

func (c *IAMClient) withAuth(opts []option.RequestOption) []option.RequestOption {
	if c.Token != "" {
		opts = append(opts, option.WithHeader("Authorization", "Bearer "+c.Token))
	}
	return opts
}

//
// =====================================
// USER APIs
// =====================================
//

// CREATE USER
type CreateUserParams struct {
	User struct {
		UserName      string `json:"userName"`
		Email         string `json:"email"`
		Password      string `json:"password"`
		ConsoleAccess bool   `json:"consoleAccess"`
	} `json:"user"`
}

func (c *IAMClient) CreateUser(ctx context.Context, body CreateUserParams, opts ...option.RequestOption) (res map[string]any, err error) {
	opts = slices.Concat(c.Options, opts)

	opts = append(opts,
		option.WithHeader("Accept", "*/*"),
		option.WithHeader("Content-Type", "application/json"),
	)

	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, "/iam/v1/user", body, &res, opts...)
	return
}

// GET USER
type GetUserParams struct {
	UserKRN string
}

func (c *IAMClient) GetUser(ctx context.Context, body GetUserParams, opts ...option.RequestOption) (res map[string]any, err error) {
	opts = slices.Concat(c.Options, opts)
	opts = c.withAuth(opts)

	path := fmt.Sprintf("/iam/v1/users/%s", body.UserKRN)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// DELETE USER
type DeleteUserParams struct {
	UserKRN string
}

func (c *IAMClient) DeleteUser(ctx context.Context, body DeleteUserParams, opts ...option.RequestOption) (res map[string]any, err error) {
	opts = slices.Concat(c.Options, opts)
	opts = c.withAuth(opts)

	path := fmt.Sprintf("/iam/v1/users/%s", body.UserKRN)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

//
// =====================================
// PROGRAMMATIC ACCESS
// =====================================
//

// SIGNIN PROGRAMMATIC USER
type ProgrammaticSigninParams struct {
	AccountID string `json:"accountId"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

func (c *IAMClient) SigninProgrammaticUser(ctx context.Context, body ProgrammaticSigninParams, opts ...option.RequestOption) (res map[string]any, err error) {
	opts = slices.Concat(c.Options, opts)

	opts = append(opts, option.WithHeader("Content-Type", "application/json"))

	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, "/iam/v1/signinProgrammaticUser", body, &res, opts...)

	// AUTO STORE TOKEN
	if token, ok := res["token"].(string); ok {
		c.Token = token
	} else if token, ok := res["access_token"].(string); ok {
		c.Token = token
	}

	return
}

// ENABLE PROGRAMMATIC ACCESS
type EnableProgrammaticAccessParams struct {
	UserKRN string
}

func (c *IAMClient) EnableProgrammaticAccess(ctx context.Context, body EnableProgrammaticAccessParams, opts ...option.RequestOption) (res map[string]any, err error) {
	opts = slices.Concat(c.Options, opts)
	opts = c.withAuth(opts)

	path := fmt.Sprintf("/iam/v1/users/programmaticAccess/%s", body.UserKRN)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// RESET
type ResetProgrammaticAccessParams struct {
	UserKRN string
}

func (c *IAMClient) ResetProgrammaticAccess(ctx context.Context, body ResetProgrammaticAccessParams, opts ...option.RequestOption) (res map[string]any, err error) {
	opts = slices.Concat(c.Options, opts)
	opts = c.withAuth(opts)

	path := fmt.Sprintf("/iam/v1/users/resetProgrammaticAccess/%s", body.UserKRN)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// DISABLE
type DisableProgrammaticAccessParams struct {
	UserKRN string
}

func (c *IAMClient) DisableProgrammaticAccess(ctx context.Context, body DisableProgrammaticAccessParams, opts ...option.RequestOption) (res map[string]any, err error) {
	opts = slices.Concat(c.Options, opts)
	opts = c.withAuth(opts)

	path := fmt.Sprintf("/iam/v1/users/disableProgrammaticAccess/%s", body.UserKRN)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

//
// =====================================
// ROLE APIs
// =====================================
//

// LIST ROLES
type ListRolesParams struct {
	Limit  int
	Offset int
}

func (c *IAMClient) ListRoles(ctx context.Context, body ListRolesParams, opts ...option.RequestOption) (res map[string]any, err error) {
	opts = slices.Concat(c.Options, opts)
	opts = c.withAuth(opts)

	opts = append(opts,
		option.WithQuery("limit", fmt.Sprintf("%d", body.Limit)),
		option.WithQuery("offset", fmt.Sprintf("%d", body.Offset)),
		option.WithQuery("krutrimManaged", "all"),
	)

	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, "/iam/v1/roles", nil, &res, opts...)
	return
}

// GET ROLE
type GetRoleParams struct {
	RoleKRN string
}

func (c *IAMClient) GetRole(ctx context.Context, body GetRoleParams, opts ...option.RequestOption) (res map[string]any, err error) {
	opts = slices.Concat(c.Options, opts)
	opts = c.withAuth(opts)

	path := fmt.Sprintf("/iam/v1/roles/%s", body.RoleKRN)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// DELETE ROLE
type DeleteRoleParams struct {
	RoleKRN string
}

func (c *IAMClient) DeleteRole(ctx context.Context, body DeleteRoleParams, opts ...option.RequestOption) (res map[string]any, err error) {
	opts = slices.Concat(c.Options, opts)
	opts = c.withAuth(opts)

	path := fmt.Sprintf("/iam/v1/roles/%s", body.RoleKRN)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// CREATE ROLE WITH POLICIES
type CreateRoleParams struct {
	Role struct {
		Name        string `json:"Name"`
		Description string `json:"Description"`
	} `json:"Role"`

	PolicyIDs []string `json:"PolicyIDs"`
}

func (c *IAMClient) CreateRoleWithPolicies(ctx context.Context, body CreateRoleParams, opts ...option.RequestOption) (res map[string]any, err error) {
	opts = slices.Concat(c.Options, opts)
	opts = c.withAuth(opts)

	opts = append(opts, option.WithHeader("Content-Type", "application/json"))

	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, "/iam/v1/role/rolepolicies", body, &res, opts...)
	return
}

//
// =====================================
// POLICY APIs
// =====================================
//

// LIST POLICIES
type ListPoliciesParams struct {
	Limit          int
	Offset         int
	KrutrimManaged string
}

func (c *IAMClient) ListPolicies(ctx context.Context, body ListPoliciesParams, opts ...option.RequestOption) (res map[string]any, err error) {
	opts = slices.Concat(c.Options, opts)
	opts = c.withAuth(opts)

	opts = append(opts,
		option.WithQuery("limit", fmt.Sprintf("%d", body.Limit)),
		option.WithQuery("offset", fmt.Sprintf("%d", body.Offset)),
		option.WithQuery("krutrimManaged", body.KrutrimManaged),
	)

	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, "/iam/v1/policies", nil, &res, opts...)
	return
}

// ATTACH POLICIES
type AttachPoliciesParams struct {
	RoleID    string   `json:"roleId"`
	PolicyIDs []string `json:"policyIds"`
}

func (c *IAMClient) AttachPoliciesToRole(ctx context.Context, body AttachPoliciesParams, opts ...option.RequestOption) (res map[string]any, err error) {
	opts = slices.Concat(c.Options, opts)
	opts = c.withAuth(opts)

	opts = append(opts, option.WithHeader("Content-Type", "application/json"))

	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, "/iam/v1/urgp/role/policy", body, &res, opts...)
	return
}

//
// =====================================
// USER-ROLE MAPPING
// =====================================
//

type AssignRolesParams struct {
	UserID   string   `json:"userId"`
	RoleIDs  []string `json:"roleIds"`
	GroupIDs []string `json:"groupIds,omitempty"`
}

func (c *IAMClient) AssignRolesToUser(ctx context.Context, body AssignRolesParams, opts ...option.RequestOption) (res map[string]any, err error) {
	opts = slices.Concat(c.Options, opts)
	opts = c.withAuth(opts)

	opts = append(opts, option.WithHeader("Content-Type", "application/json"))

	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, "/iam/v1/urgp/user/role/group", body, &res, opts...)
	return
}